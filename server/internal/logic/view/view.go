// Package view
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package view

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model"
	"hotgo/internal/service"
)

type sView struct{}

func init() {
	service.RegisterView(New())
}

func New() *sView {
	return &sView{}
}

// GetBreadCrumb 前台系统-获取面包屑列表
func (s *sView) GetBreadCrumb(ctx context.Context, in *model.ViewGetBreadCrumbInput) []model.ViewBreadCrumb {
	breadcrumb := []model.ViewBreadCrumb{
		{Name: "首页", Url: "/"},
	}
	return breadcrumb
}

// GetTitle 前台系统-获取标题
func (s *sView) GetTitle(ctx context.Context, in *model.ViewGetTitleInput) string {
	return "title"
}

// RenderTpl 渲染指定模板页面
func (s *sView) RenderTpl(ctx context.Context, tpl string, data ...model.View) {
	var (
		viewObj = model.View{}
		request = g.RequestFromCtx(ctx)
	)
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.Title == "" {
		viewObj.Title = g.Cfg().MustGet(ctx, `setting.title`).String()
	} else {
		viewObj.Title = viewObj.Title + ` - ` + g.Cfg().MustGet(ctx, `setting.title`).String()
	}
	if viewObj.Keywords == "" {
		viewObj.Keywords = g.Cfg().MustGet(ctx, `setting.keywords`).String()
	}
	if viewObj.Description == "" {
		viewObj.Description = g.Cfg().MustGet(ctx, `setting.description`).String()
	}
	if viewObj.IpcCode == "" {
		viewObj.IpcCode = g.Cfg().MustGet(ctx, `setting.icpCode`).String()
	}

	if viewObj.GET == nil {
		viewObj.GET = request.GetQueryMap()
	}

	// 去掉空数据
	viewData := gconv.Map(viewObj)
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}
	// 内置对象
	viewData["BuildIn"] = &viewBuildIn{httpRequest: request}

	// 渲染模板
	_ = request.Response.WriteTpl(tpl, viewData)
}

// Render 渲染默认模板页面
func (s *sView) Render(ctx context.Context, data ...model.View) {
	s.RenderTpl(ctx, g.Cfg().MustGet(ctx, "viewer.homeLayout").String(), data...)
}

// Render302 跳转中间页面
func (s *sView) Render302(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "页面跳转中"
	}
	s.RenderTpl(ctx, "default/pages/302.html", view)
}

// Render401 401页面
func (s *sView) Render401(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	s.RenderTpl(ctx, "default/pages/401.html", view)
}

// Render403 403页面
func (s *sView) Render403(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	s.RenderTpl(ctx, "default/pages/403.html", view)
}

// Render404 404页面
func (s *sView) Render404(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "资源不存在"
	}
	s.RenderTpl(ctx, "default/pages/404.html", view)
}

// Render500 500页面
func (s *sView) Render500(ctx context.Context, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "请求执行错误"
	}
	s.RenderTpl(ctx, "default/pages/500.html", view)
}

func (s *sView) Error(ctx context.Context, err error) {
	view := model.View{
		Title: "错误提示",
		Error: err.Error(),
	}
	s.RenderTpl(ctx, "default/pages/500.html", view)
}
