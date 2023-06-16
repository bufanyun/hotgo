// Package view
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package view

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model"
	"hotgo/internal/service"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
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

// Error 自定义错误页面
func (s *sView) Error(ctx context.Context, err error) {
	var (
		request = g.RequestFromCtx(ctx)
		code    = gerror.Code(err)
		stack   string
	)

	// 是否输出错误堆栈到页面
	if g.Cfg().MustGet(ctx, "hotgo.debug", true).Bool() {
		stack = charset.SerializeStack(err)
	}

	request.Response.ClearBuffer()
	_ = request.Response.WriteTplContent(simple.DefaultErrorTplContent(ctx), g.Map{
		"code":    code.Code(),
		"message": code.Message(),
		"stack":   stack,
	})
}
