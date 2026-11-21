// Package casbin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package casbin

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"net/http"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	ActionGet    = http.MethodGet
	ActionPost   = http.MethodPost
	ActionPut    = http.MethodPut
	ActionDelete = http.MethodDelete
	ActionAll    = "GET|POST|PUT|DELETE|PATCH|OPTIONS|HEAD"
)

var (
	Enforcer *casbin.SyncedEnforcer
	// 刷新串行化，避免旧快照在并发刷新后覆盖新快照
	refreshMu sync.Mutex
)

// InitEnforcer 初始化
func InitEnforcer(ctx context.Context) {
	var (
		link   = getDbLink(ctx)
		a, err = NewAdapter(link.String())
	)

	if err != nil {
		g.Log().Panicf(ctx, "casbin.NewAdapter err . %v", err)
		return
	}

	path := "manifest/config/casbin.conf"

	// 优先从本地加载casbin.conf，如果不存在就从资源文件中找
	modelContent := gfile.GetContents(path)
	if len(modelContent) == 0 {
		if !gres.IsEmpty() && gres.Contains(path) {
			modelContent = string(gres.GetContent(path))
		}
	}

	if len(modelContent) == 0 {
		g.Log().Panicf(ctx, "casbin model file does not exist：%v", path)
	}

	m, err := model.NewModelFromString(modelContent)
	if err != nil {
		g.Log().Panicf(ctx, "casbin NewModelFromString err：%v", err)
	}

	Enforcer, err = casbin.NewSyncedEnforcer(m, a)
	if err != nil {
		g.Log().Panicf(ctx, "casbin NewSyncedEnforcer err：%v", err)
	}
	if err = Refresh(ctx); err != nil {
		g.Log().Fatalf(ctx, "casbin Refresh err:%v", err)
	}
}

// GetDbLink 获取数据库链接
func getDbLink(ctx context.Context) *gvar.Var {

	link := g.Cfg().MustGet(ctx, "database.default")
	//读写分离
	if !link.IsSlice() {
		return g.Cfg().MustGet(ctx, "database.default.link")
	}

	for _, v := range link.Array() {
		// 只获取主库
		val := v.(map[string]interface{})
		if val["role"] == "master" {
			return gvar.New(val["link"])
		}
	}
	return gvar.New("database.default.0.link")

}

func buildRules(ctx context.Context) (rules [][]string, err error) {
	type Policy struct {
		Key         string `json:"key"`
		Permissions string `json:"permissions"`
	}
	var polices []*Policy
	//别名拼接 r.key m.permissions
	q := func(alias string, column string) string {
		return fmt.Sprintf("%s.%s", alias, column)
	}
	err = g.Model(gstr.Join([]string{dao.AdminRole.Table(), "r"}, " ")).
		LeftJoin(gstr.Join([]string{dao.AdminRoleMenu.Table(), "rm"}, " "), "r.id=rm.role_id").
		LeftJoin(gstr.Join([]string{dao.AdminMenu.Table(), "m"}, " "), "rm.menu_id=m.id").
		Fields(q("r", dao.AdminRole.Columns().Key), q("m", dao.AdminMenu.Columns().Permissions)).
		Where(q("r", dao.AdminRole.Columns().Status), consts.StatusEnabled).
		Where(q("m", dao.AdminMenu.Columns().Status), consts.StatusEnabled).
		WhereNot(q("m", dao.AdminMenu.Columns().Permissions), "").
		WhereNot(q("r", dao.AdminRole.Columns().Key), consts.SuperRoleKey).
		Scan(&polices)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色权限失败")
	}

	// 同一角色经不同菜单拿到相同权限时只留一条，重复策略会让 AddPolicies 返回 false
	seen := make(map[string]struct{}, len(polices))
	for _, policy := range polices {
		for _, permissions := range strings.Split(policy.Permissions, ",") {
			permissions = strings.TrimSpace(permissions)
			if permissions == "" {
				continue
			}
			key := policy.Key + "\x00" + permissions
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			rules = append(rules, []string{policy.Key, permissions, ActionAll})
		}
	}
	return rules, nil
}

func Clear(ctx context.Context) (err error) {
	refreshMu.Lock()
	defer refreshMu.Unlock()

	lock := Enforcer.GetLock()
	lock.Lock()
	defer lock.Unlock()

	policy, err := Enforcer.Enforcer.GetPolicy()
	if err != nil {
		return gerror.Wrap(err, "获取 casbin 策略失败")
	}
	if len(policy) == 0 {
		return
	}
	// RemovePolicies 会原地压缩策略切片，不复制会漏删
	if _, err = Enforcer.Enforcer.RemovePolicies(append([][]string(nil), policy...)); err != nil {
		return gerror.Wrap(err, "清理 casbin 策略失败")
	}
	return
}

// 查库不持 Enforcer 锁，避免阻塞鉴权
func Refresh(ctx context.Context) (err error) {
	refreshMu.Lock()
	defer refreshMu.Unlock()

	rules, err := buildRules(ctx)
	if err != nil {
		return err
	}

	lock := Enforcer.GetLock()
	lock.Lock()
	defer lock.Unlock()

	policy, err := Enforcer.Enforcer.GetPolicy()
	if err != nil {
		return gerror.Wrap(err, "获取 casbin 策略失败")
	}
	// RemovePolicies 会原地压缩策略切片，不复制会漏删
	old := append([][]string(nil), policy...)
	if len(old) > 0 {
		if _, err = Enforcer.Enforcer.RemovePolicies(old); err != nil {
			g.Log().Errorf(ctx, "casbin 清理旧策略失败，内存策略可能已部分清除 err:%+v", err)
			return gerror.Wrap(err, "清理 casbin 策略失败")
		}
	}
	if len(rules) > 0 {
		ok, addErr := Enforcer.Enforcer.AddPolicies(rules)
		if addErr != nil || !ok {
			if _, restoreErr := Enforcer.Enforcer.AddPolicies(old); restoreErr != nil {
				g.Log().Errorf(ctx, "casbin 恢复旧策略失败 err:%+v", restoreErr)
			}
			if addErr != nil {
				return gerror.Wrap(addErr, "加载 casbin 策略失败")
			}
			return gerror.New("加载 casbin 策略失败：规则未写入")
		}
	}
	return
}
