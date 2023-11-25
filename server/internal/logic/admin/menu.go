// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"hotgo/api/admin/role"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/validate"
)

type sAdminMenu struct{}

func NewAdminMenu() *sAdminMenu {
	return &sAdminMenu{}
}

func init() {
	service.RegisterAdminMenu(NewAdminMenu())
}

// Delete 删除
func (s *sAdminMenu) Delete(ctx context.Context, in *adminin.MenuDeleteInp) (err error) {
	exist, err := dao.AdminMenu.Ctx(ctx).Where("pid", in.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先删除该菜单下的所有菜单！")
	}
	_, err = dao.AdminMenu.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// VerifyUnique 验证菜单唯一属性
func (s *sAdminMenu) VerifyUnique(ctx context.Context, in *adminin.VerifyUniqueInp) (err error) {
	if in.Where == nil {
		return
	}

	cols := dao.AdminMenu.Columns()
	msgMap := g.MapStrStr{
		cols.Name:  "菜单编码已存在，请换一个",
		cols.Title: "菜单名称已存在，请换一个",
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.AdminMenu, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}

// Edit 修改/新增
func (s *sAdminMenu) Edit(ctx context.Context, in *adminin.MenuEditInp) (err error) {
	// 验证唯一性
	err = s.VerifyUnique(ctx, &adminin.VerifyUniqueInp{
		Id: in.Id,
		Where: g.Map{
			dao.AdminMenu.Columns().Title: in.Title,
			dao.AdminMenu.Columns().Name:  in.Name,
		},
	})
	if err != nil {
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		in.Pid, in.Level, in.Tree, err = hgorm.AutoUpdateTree(ctx, &dao.AdminMenu, in.Id, in.Pid)
		if err != nil {
			return err
		}

		if in.Id > 0 {
			if _, err = dao.AdminMenu.Ctx(ctx).Where("id", in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改菜单失败！")
				return err
			}
		} else {
			if _, err = dao.AdminMenu.Ctx(ctx).Data(in).Insert(); err != nil {
				err = gerror.Wrap(err, "新增菜单失败！")
				return err
			}
		}
		return casbin.Refresh(ctx)
	})
}

// List 获取菜单列表
func (s *sAdminMenu) List(ctx context.Context, in *adminin.MenuListInp) (res *adminin.MenuListModel, err error) {
	var models []*entity.AdminMenu
	if err = dao.AdminMenu.Ctx(ctx).Order("sort asc,id desc").Scan(&models); err != nil {
		return
	}

	res = new(adminin.MenuListModel)
	res.List = s.treeList(0, models)
	return
}

// genNaiveMenus 生成NaiveUI菜单格式
func (s *sAdminMenu) genNaiveMenus(menus []*adminin.MenuRouteSummary) (sources []*adminin.MenuRoute) {
	for _, men := range menus {
		var source = new(adminin.MenuRoute)
		source.Name = men.Name
		source.Path = men.Path
		source.Redirect = men.Redirect
		source.Component = men.Component
		source.Meta = &adminin.MenuRouteMeta{
			Title:      men.Title,
			Icon:       men.Icon,
			KeepAlive:  men.KeepAlive == 1,
			Hidden:     men.Hidden == 1,
			Sort:       men.Sort,
			AlwaysShow: men.AlwaysShow == 1,
			ActiveMenu: men.ActiveMenu,
			IsRoot:     men.IsRoot == 1,
			FrameSrc:   men.FrameSrc,
			//Permissions: men.Permissions,
			Affix: men.Affix == 1,
			Type:  men.Type,
		}
		if len(men.Children) > 0 {
			source.Children = append(source.Children, s.genNaiveMenus(men.Children)...)
		}
		sources = append(sources, source)
	}
	return
}

// getChildrenList 生成菜单树
func (s *sAdminMenu) getChildrenList(menu *adminin.MenuRouteSummary, treeMap map[string][]*adminin.MenuRouteSummary) (err error) {
	menu.Children = treeMap[gconv.String(menu.Id)]
	for i := 0; i < len(menu.Children); i++ {
		if err = s.getChildrenList(menu.Children[i], treeMap); err != nil {
			return
		}
	}
	return
}

// GetMenuList 获取菜单列表
func (s *sAdminMenu) GetMenuList(ctx context.Context, memberId int64) (res *role.DynamicRes, err error) {
	var (
		allMenus []*adminin.MenuRouteSummary
		menus    []*adminin.MenuRouteSummary
		treeMap  = make(map[string][]*adminin.MenuRouteSummary)
		mod      = dao.AdminMenu.Ctx(ctx).Where("status", consts.StatusEnabled).WhereIn("type", []int{1, 2})
	)

	// 非超管验证允许的菜单列表
	if !service.AdminMember().VerifySuperId(ctx, memberId) {
		menuIds, err := dao.AdminRoleMenu.Ctx(ctx).Fields("menu_id").Where("role_id", contexts.GetRoleId(ctx)).Array()
		if err != nil {
			return nil, err
		}
		if len(menuIds) > 0 {
			pidList, err := dao.AdminMenu.Ctx(ctx).Fields("pid").WhereIn("id", menuIds).Group("pid").Array()
			if err != nil {
				return nil, err
			}
			if len(pidList) > 0 {
				menuIds = append(pidList, menuIds...)
			}
		}
		mod = mod.Where("id", menuIds)
	}

	if err = mod.Order("sort asc,id desc").Scan(&allMenus); err != nil || len(allMenus) == 0 {
		return
	}

	// 生产环境下隐藏一些菜单
	if gmode.IsProduct() {
		newMenus := make([]*adminin.MenuRouteSummary, 0)
		devMenus := []string{"Develops", "doc"} // 如果你还有其他需要在生产环境隐藏的菜单，将菜单别名加入即可
		for _, menu := range allMenus {
			if !validate.InSlice(devMenus, menu.Name) {
				newMenus = append(newMenus, menu)
			}
		}
		allMenus = newMenus
	}

	for _, v := range allMenus {
		treeMap[gconv.String(v.Pid)] = append(treeMap[gconv.String(v.Pid)], v)
	}

	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = s.getChildrenList(menus[i], treeMap)
	}

	res = new(role.DynamicRes)
	res.List = append(res.List, s.genNaiveMenus(menus)...)
	return
}

// LoginPermissions 获取登录成功后的细粒度权限
func (s *sAdminMenu) LoginPermissions(ctx context.Context, memberId int64) (lists adminin.MemberLoginPermissions, err error) {
	type Permissions struct {
		Permissions string `json:"permissions"`
	}

	var (
		allPermissions []*Permissions
		mod            = dao.AdminMenu.Ctx(ctx).Fields("permissions").Where("status", consts.StatusEnabled).Where("permissions != ?", "")
	)

	// 非超管验证允许的菜单列表
	if !service.AdminMember().VerifySuperId(ctx, memberId) {
		menuIds, err := dao.AdminRoleMenu.Ctx(ctx).Fields("menu_id").Where("role_id", contexts.GetRoleId(ctx)).Array()
		if err != nil {
			return nil, err
		}
		mod = mod.Where("id", menuIds)
	}

	if err = mod.Scan(&allPermissions); err != nil {
		return
	}

	// 无权限
	if len(allPermissions) == 0 {
		lists = append(lists, "value")
		return
	}

	for _, v := range allPermissions {
		for _, p := range gstr.Explode(`,`, v.Permissions) {
			lists = append(lists, p)
		}
	}

	lists = convert.UniqueSlice(lists)
	return
}

// treeList 树状列表
func (s *sAdminMenu) treeList(pid int64, nodes []*entity.AdminMenu) (list []*adminin.MenuTree) {
	list = make([]*adminin.MenuTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(adminin.MenuTree)
			item.AdminMenu = *v
			item.Label = v.Title
			item.Key = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}
