// Package casbin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package casbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"net/http"
	"strings"
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
		link   = g.Cfg().MustGet(ctx, "database.default.link")
		a, err = NewAdapter(link.String())
	)

	if err != nil {
		g.Log().Panicf(ctx, "casbin.NewAdapter err . %v", err)
		return
	}

	Enforcer, err = casbin.NewEnforcer("./manifest/config/casbin.conf", a)
	if err != nil {
		g.Log().Panicf(ctx, "casbin.NewEnforcer err . %v", err)
		return
	}

	loadPermissions(ctx)
}

func loadPermissions(ctx context.Context) {
	type Policy struct {
		Key         string `json:"key"`
		Permissions string `json:"permissions"`
	}
	var (
		rules        [][]string
		polices      []*Policy
		err          error
		superRoleKey = g.Cfg().MustGet(ctx, "hotgo.admin.superRoleKey")
	)

	err = g.Model("hg_admin_role r").
		LeftJoin("hg_admin_role_menu rm", "r.id=rm.role_id").
		LeftJoin("hg_admin_menu m", "rm.menu_id=m.id").
		Fields("r.key,m.permissions").
		Where("r.status", consts.StatusEnabled).
		Where("m.status", consts.StatusEnabled).
		Where("m.permissions !=?", "").
		Where("r.key !=?", superRoleKey.String()).
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
	_, err = Enforcer.RemovePolicies(Enforcer.GetPolicy())
	if err != nil {
		g.Log().Warningf(ctx, "Enforcer RemovePolicies err:%+v", err)
		return
	}

	// 检查是否清理干净
	if len(Enforcer.GetPolicy()) > 0 {
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
