// Package cron
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cron

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/utility/simple"
	"strings"
	"sync"
)

var crons = &cronManager{
	tasks: make(map[string]*TaskItem),
}

// cronStrategy 任务接口
type cronStrategy interface {
	GetName() string
	Execute(ctx context.Context)
}

// consumerManager 任务管理者
type cronManager struct {
	tasks map[string]*TaskItem
	sync.RWMutex
}

type TaskItem struct {
	Pattern string        // 表达式，参考：https://goframe.org/pages/viewpage.action?pageId=30736411
	Name    string        // 唯一的任务名称
	Params  string        // 函数参数，多个用,隔开
	Fun     gcron.JobFunc // 执行的函数接口
	Policy  int64         // 策略 1：并行 2：单例 3：单次 4：多次
	Count   int           // 执行次数，仅Policy=4时有效
}

func Logger() *glog.Logger {
	return g.Log("cron")
}

// Register 注册任务
func Register(c cronStrategy) {
	crons.Lock()
	defer crons.Unlock()

	name := c.GetName()
	if _, ok := crons.tasks[name]; ok {
		Logger().Debugf(gctx.GetInitCtx(), "cron.Register name:%v duplicate registration.", name)
		return
	}
	crons.tasks[name] = &TaskItem{Name: c.GetName(), Fun: c.Execute}
}

// StopALL 停止所有任务
func StopALL() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
	}
}

// StartALL 启动所有任务
func StartALL(sysCron []*entity.SysCron) (err error) {
	if len(crons.tasks) == 0 {
		g.Log().Debug(gctx.GetInitCtx(), "no scheduled task is available.")
		return
	}

	for _, cron := range sysCron {
		f, ok := crons.tasks[cron.Name]
		if !ok {
			return gerror.Newf("该任务没有加入任务列表:%v", cron.Name)
		}

		// 没有则添加
		if gcron.Search(cron.Name) == nil {
			var (
				t   *gcron.Entry
				ctx = context.WithValue(gctx.New(), consts.ContextKeyCronArgs, strings.Split(cron.Params, consts.CronSplitStr))
			)
			switch cron.Policy {
			case consts.CronPolicySame:
				t, err = gcron.Add(ctx, cron.Pattern, f.Fun, cron.Name)

			case consts.CronPolicySingle:
				t, err = gcron.AddSingleton(ctx, cron.Pattern, f.Fun, cron.Name)

			case consts.CronPolicyOnce:
				t, err = gcron.AddOnce(ctx, cron.Pattern, f.Fun, cron.Name)

			case consts.CronPolicyTimes:
				if f.Count <= 0 {
					f.Count = 1
				}
				t, err = gcron.AddTimes(ctx, cron.Pattern, int(cron.Count), f.Fun, cron.Name)

			default:
				return gerror.Newf("使用无效的策略, cron.Policy=%v", cron.Policy)
			}

			if err != nil {
				return err
			}
			if t == nil {
				return gerror.New("启动任务失败")
			}
		}

		gcron.Start(cron.Name)

		// 执行完毕，单次和多次执行的任务更新状态
		if cron.Policy == consts.CronPolicyOnce || cron.Policy == consts.CronPolicyTimes {
			if _, err = dao.SysCron.Ctx(gctx.GetInitCtx()).Where("id", cron.Id).Data(g.Map{"status": consts.StatusDisable, "updated_at": gtime.Now()}).Update(); err != nil {
				err = gerror.Wrap(err, "定时任务执行失败！")
				return err
			}
		}
	}

	Logger().Debug(gctx.GetInitCtx(), "load cron success..")
	return nil
}

// RefreshStatus 刷新状态
func RefreshStatus(sysCron *entity.SysCron) (err error) {
	if sysCron == nil {
		return
	}

	if sysCron.Status == consts.StatusEnabled {
		return Start(sysCron)
	}
	return Stop(sysCron)
}

// Stop 停止单个任务
func Stop(sysCron *entity.SysCron) (err error) {
	cr := gcron.Search(sysCron.Name)
	if cr == nil {
		return
	}
	cr.Stop()
	return
}

// Once 立即执行一次某个任务
func Once(ctx context.Context, sysCron *entity.SysCron) error {
	for _, v := range crons.tasks {
		if v.Name == sysCron.Name {
			simple.SafeGo(ctx, func(ctx context.Context) {
				v.Fun(ctx)
			})
			return nil
		}
	}
	return gerror.Newf("定时任务不存在：%+v", sysCron.Name)
}

// Delete 删除任务
func Delete(sysCron *entity.SysCron) (err error) {
	if sysCron == nil {
		return
	}

	for _, v := range gcron.Entries() {
		if v.Name == sysCron.Name {
			gcron.Remove(v.Name)
		}
	}
	return
}

// Start 启动单个任务
func Start(sysCron *entity.SysCron) (err error) {
	if sysCron == nil {
		return
	}

	c := gcron.Search(sysCron.Name)
	if c != nil {
		c.Start()
		return
	}
	return StartALL([]*entity.SysCron{sysCron})
}
