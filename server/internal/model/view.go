// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package model

// View 视图渲染内容对象
type View struct {
	Title       string                 // 页面标题
	Keywords    string                 // 页面Keywords
	Description string                 // 页面Description
	IpcCode     string                 // ICP备案号
	Error       string                 // 错误信息
	MainTpl     string                 // 自定义MainTpl展示模板文件
	Redirect    string                 // 引导页面跳转
	ContentType string                 // 内容模型
	BreadCrumb  []ViewBreadCrumb       // 面包屑
	GET         map[string]interface{} // GET参数
	Data        interface{}            // 页面参数
}

// ViewBreadCrumb 视图面包屑结构
type ViewBreadCrumb struct {
	Name string // 显示名称
	Url  string // 链接地址，当为空时表示被选中
}

// ViewGetBreadCrumbInput 获取面包屑请求
type ViewGetBreadCrumbInput struct {
	ContentId   uint   // (可选)内容ID
	ContentType string // (可选)内容类型
	CategoryId  uint   // (可选)栏目ID
}

// ViewGetTitleInput 获取title请求
type ViewGetTitleInput struct {
	ContentType string // (可选)内容类型
	CategoryId  uint   // (可选)栏目ID
	CurrentName string // (可选)当前名称
}
