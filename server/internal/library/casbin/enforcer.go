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

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/container/gvar"
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

var Enforcer *casbin.Enforcer

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

	Enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		g.Log().Panicf(ctx, "casbin NewEnforcer err：%v", err)
	}
	loadPermissions(ctx)
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

func loadPermissions(ctx context.Context) {
	type Policy struct {
		Key         string `json:"key"`
		Permissions string `json:"permissions"`
	}
	var (
		rules   [][]string
		polices []*Policy
		err     error
	)
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
		g.Log().Fatalf(ctx, "loadPermissions Scan err:%v", err)
		return
	}

	for _, policy := range polices {
		if strings.Contains(policy.Permissions, ",") {
			lst := strings.Split(policy.Permissions, ",")
			for _, permissions := range lst {
				rules = append(rules, []string{policy.Key, permissions, ActionAll})
			}
		} else {
			rules = append(rules, []string{policy.Key, policy.Permissions, ActionAll})
		}
	}

	if _, err = Enforcer.AddPolicies(rules); err != nil {
		g.Log().Fatalf(ctx, "loadPermissions AddPolicies err:%v", err)
		return
	}
}

func Clear(ctx context.Context) (err error) {
	policy, err := Enforcer.GetPolicy()
	if err != nil {
		g.Log().Warningf(ctx, "Enforcer RemovePolicies err:%+v", err)
		return
	}
	_, err = Enforcer.RemovePolicies(policy)
	if err != nil {
		g.Log().Warningf(ctx, "Enforcer RemovePolicies err:%+v", err)
		return
	}

	// 检查是否清理干净
	if len(policy) > 0 {
		return Clear(ctx)
	}
	return
}

func Refresh(ctx context.Context) (err error) {
	if err = Clear(ctx); err != nil {
		return err
	}
	loadPermissions(ctx)
	return
}
