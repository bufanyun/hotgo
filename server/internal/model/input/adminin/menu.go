// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// MenuEditInp 修改/新增菜单
type MenuEditInp struct {
	entity.AdminMenu
}

func (in *MenuEditInp) Filter(ctx context.Context) (err error) {
	if in.Title == "" {
		err = gerror.New("菜单名称不能为空")
		return
	}
	if in.Type != 3 && in.Path == "" {
		err = gerror.New("路由地址不能为空")
		return
	}
	if in.Name == "" {
		err = gerror.New("路由名称不能为空")
		return
	}
	return
}

type MenuEditModel struct{}

// MenuDeleteInp 删除菜单
type MenuDeleteInp struct {
	Id interface{} `json:"id" v:"required#菜单ID不能为空" dc:"菜单ID"`
}

func (in *MenuDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type MenuDeleteModel struct{}

// MenuListInp 获取菜单列表
type MenuListInp struct {
	form.PageReq
	Pid int64 `json:"pid" dc:"父ID"`
}

func (in *MenuListInp) Filter(ctx context.Context) (err error) {
	return
}

// MenuSearchListInp 查询菜单列表
type MenuSearchListInp struct {
	Name string `json:"name" dc:"菜单名称"`
	form.StatusReq
}

func (in *MenuSearchListInp) Filter(ctx context.Context) (err error) {
	return
}

type MenuSearchListModel []*model.TreeMenu

// MenuTree 树
type MenuTree struct {
	entity.AdminMenu
	Key      int64       `json:"key" `
	Label    string      `json:"label"     dc:"标签"`
	Children []*MenuTree `json:"children"`
}

type MenuListModel struct {
	List []*MenuTree `json:"list"`
}

// MenuRouteMeta 菜单路由
type MenuRouteMeta struct {
	// 解释参考：https://naive-ui-admin-docs.vercel.app/guide/router.html#%E5%A4%9A%E7%BA%A7%E8%B7%AF%E7%94%B1
	Title string `json:"title"` // 菜单名称 一般必填
	//Disabled   bool   `json:"disabled,omitempty"`   // 禁用菜单
	Icon       string `json:"icon,omitempty"`       // 菜单图标
	KeepAlive  bool   `json:"keepAlive,omitempty"`  // 缓存该路由
	Hidden     bool   `json:"hidden,omitempty"`     // 隐藏菜单
	Sort       int    `json:"sort,omitempty"`       // 排序越小越排前
	AlwaysShow bool   `json:"alwaysShow,omitempty"` // 取消自动计算根路由模式
	ActiveMenu string `json:"activeMenu,omitempty"` // 当路由设置了该属性，则会高亮相对应的侧边栏。
	// 这在某些场景非常有用，比如：一个列表页路由为：/list/basic-list
	// 点击进入详情页，这时候路由为/list/basic-info/1，但你想在侧边栏高亮列表的路由，就可以进行如下设置
	// 注意是配置高亮路由 `name`，不是path
	IsRoot      bool   `json:"isRoot,omitempty"`      // 是否跟路由 顶部混合菜单，必须传 true，否则左侧会显示异常（场景就是，分割菜单之后，当一级菜单没有子菜单）
	FrameSrc    string `json:"frameSrc,omitempty" `   // 内联外部地址
	Permissions string `json:"permissions,omitempty"` // 菜单包含权限集合，满足其中一个就会显示
	Affix       bool   `json:"affix,omitempty"`       // 是否固定 设置为 true 之后 多页签不可删除
	Type        int    `json:"type"`                  // 菜单类型
}

type MenuRoute struct {
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Redirect  string         `json:"redirect"`
	Component string         `json:"component"`
	Meta      *MenuRouteMeta `json:"meta"`
	Children  []*MenuRoute   `json:"children,omitempty" dc:"子路由"`
}

// MenuRouteSummary 菜单树结构
type MenuRouteSummary struct {
	entity.AdminMenu
	Children []*MenuRouteSummary
}

// DynamicMeta 动态路由元数据
type DynamicMeta struct {
	Title   string `json:"title"       description:"菜单标题"`
	Icon    string `json:"icon"        description:"菜单图标"`
	NoCache bool   `json:"noCache"     description:"是否缓存"`
	Remark  string `json:"remark"      description:"备注"`
}

// DynamicBase 动态路由
type DynamicBase struct {
	Id         int64        `json:"id"             description:"菜单ID"`
	Pid        int64        `json:"pid"            description:"父ID"`
	Name       string       `json:"name"           description:"菜单名称"`
	Code       string       `json:"code"           description:"菜单编码"`
	Path       string       `json:"path"           description:"路由地址"`
	Hidden     bool         `json:"hidden"         description:"是否隐藏"`
	Redirect   string       `json:"redirect"       description:"重定向"`
	Component  string       `json:"component"      description:"组件路径"`
	AlwaysShow bool         `json:"alwaysShow"     description:"暂时不知道干啥"`
	IsFrame    string       `json:"isFrame"        description:"是否为外链（0是 1否）"`
	Meta       *DynamicMeta `json:"meta"           description:"配置数据集"`
}

// DynamicMenu 动态路由菜单
type DynamicMenu struct {
	DynamicBase
	Children []*DynamicBase `json:"children"   description:"子菜单"`
}
