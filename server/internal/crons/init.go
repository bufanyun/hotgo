// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package crons

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"strings"
	"sync"
)

type cronStrategy interface {
	GetName() string
	Execute(ctx context.Context)
}

var (
	// 添加新的任务时，只需实现cronStrategy接口，并加入到cronList即可
	cronList []cronStrategy
	inst     = new(tasks)
)

type tasks struct {
	list []*TaskItem
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

func LoadCronList() {
	for _, cron := range cronList {
		inst.Add(&TaskItem{
			Name: cron.GetName(),
			Fun:  cron.Execute,
		})
	}
}

func StopALL() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
	}
}

// StartALL 启动任务
func StartALL(sysCron []*entity.SysCron) error {
	if len(inst.list) == 0 {
		LoadCronList()
	}

	var (
		err error
		ct  = gctx.New()
	)

	if len(sysCron) == 0 {
		g.Log().Info(ct, "没有可用的定时任务")
		return nil
	}

	for _, cron := range sysCron {
		f := inst.Get(cron.Name)
		if f == nil {
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
			_, err = dao.SysCron.Ctx(ct).Where("id", cron.Id).
				Data(g.Map{"status": consts.StatusDisable, "updated_at": gtime.Now()}).
				Update()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return err
			}
		}
	}

	g.Log().Debug(ct, "load cron success..")
	return nil
}

// RefreshStatus 刷新状态
func RefreshStatus(sysCron *entity.SysCron) (err error) {
	if sysCron == nil {
		return
	}
	g.DumpWithType(sysCron)

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
	for _, v := range cronList {
		if v.GetName() == sysCron.Name {
			go v.Execute(ctx)
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
	cr := gcron.Search(sysCron.Name)
	if cr != nil {
		cr.Start()
		return
	}

	return StartALL([]*entity.SysCron{sysCron})
}

// Add 添加任务
func (t *tasks) Add(task *TaskItem) *tasks {
	if task.Name == "" || task.Fun == nil {
		return t
	}
	t.Lock()
	defer t.Unlock()
	t.list = append(t.list, task)
	return t
}

// Get 找到任务
func (t *tasks) Get(name string) *TaskItem {
	if len(t.list) == 0 {
		return nil
	}

	for _, item := range t.list {
		if item.Name == name {
			return item
		}
	}
	return nil
}

// Del 删除任务
func (t *tasks) Del(name string) (newList []*TaskItem) {
	if len(t.list) == 0 {
		return nil
	}
	t.Lock()
	defer t.Unlock()

	for _, item := range t.list {
		if item.Name == name {
			continue
		}
		newList = append(newList, item)
	}
	return newList
}
