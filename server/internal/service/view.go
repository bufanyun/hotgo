// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model"
)

type (
	IView interface {
		// GetBreadCrumb 前台系统-获取面包屑列表
		GetBreadCrumb(ctx context.Context, in *model.ViewGetBreadCrumbInput) []model.ViewBreadCrumb
		// GetTitle 前台系统-获取标题
		GetTitle(ctx context.Context, in *model.ViewGetTitleInput) string
		// RenderTpl 渲染指定模板页面
		RenderTpl(ctx context.Context, tpl string, data ...model.View)
		// Render 渲染默认模板页面
		Render(ctx context.Context, data ...model.View)
		// Error 自定义错误页面
		Error(ctx context.Context, err error)
	}
)

var (
	localView IView
)

func View() IView {
	if localView == nil {
		panic("implement not found for interface IView, forgot register?")
	}
	return localView
}

func RegisterView(i IView) {
	localView = i
}
