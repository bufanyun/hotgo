// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package router

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/controller/backend/admin"
	"hotgo/internal/controller/backend/common"
	"hotgo/internal/controller/backend/sys"
	"hotgo/internal/service"
)

func Admin(ctx context.Context, group *ghttp.RouterGroup) {

	routerPrefix, _ := g.Cfg().Get(ctx, "router.admin.prefix", "/admin")
	group.Group(routerPrefix.String(), func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().AdminAuth)
		group.Bind(
			common.Site,    // 基础
			common.Console, // 控制台
			common.Ems,     // 邮件
			common.Upload,  // 上传
			sys.Log,        // 日志
			sys.Config,     // 配置
			sys.DictType,   // 字典类型
			sys.DictData,   // 字典数据
			sys.Attachment, // 附件
			sys.Provinces,  // 省市区
			sys.Cron,       // 定时任务
			sys.CronGroup,  // 定时任务分组
			sys.Blacklist,  // 黑名单
			admin.Member,   // 用户
			admin.Monitor,  // 监控
			admin.Role,     // 路由
			admin.Dept,     // 部门
			admin.Menu,     // 菜单
			admin.Notice,   // 公告
			admin.Post,     // 岗位
		)
	})
}
