// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/backend/menu"
	"hotgo/api/backend/role"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
	"hotgo/utility/tree"
)

type sAdminMenu struct{}

func NewAdminMenu() *sAdminMenu {
	return &sAdminMenu{}
}

func init() {
	service.RegisterAdminMenu(NewAdminMenu())
}

// RoleList 查询角色菜单列表
func (s *sAdminMenu) RoleList(ctx context.Context, in adminin.MenuRoleListInp) (*adminin.MenuRoleListModel, error) {
	var (
		mod         = dao.AdminRoleMenu.Ctx(ctx)
		roleMenu    []*entity.AdminRoleMenu
		lst         []*model.LabelTreeMenu
		res         adminin.MenuRoleListModel
		err         error
		checkedKeys []int64
	)

	// 获取选中菜单ID
	if in.RoleId > 0 {
		mod = mod.Where("role_id", in.RoleId)
	}

	if err = mod.Fields().Scan(&roleMenu); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	for i := 0; i < len(roleMenu); i++ {
		checkedKeys = append(checkedKeys, roleMenu[i].MenuId)
	}
	res.CheckedKeys = checkedKeys

	// 获取菜单树
	lst, err = dao.AdminMenu.GenLabelTreeList(ctx, 0)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	_ = gconv.Structs(lst, &res.Menus)

	return &res, nil
}

// SearchList 查询菜单列表
func (s *sAdminMenu) SearchList(ctx context.Context, req *menu.SearchListReq) (*menu.SearchListRes, error) {
	var (
		mod          = dao.AdminMenu.Ctx(ctx)
		lst          []*model.TreeMenu
		res          menu.SearchListRes
		searchResult []*entity.AdminMenu
		id           int64
		ids          []int64
		err          error
	)

	if req.Name != "" {
		mod = mod.WhereLike("name", "%"+req.Name+"%")
	}

	if req.Status > 0 {
		mod = mod.Where("status", req.Status)
	}

	if req.Name != "" || req.Status > 0 {
		err = mod.Scan(&searchResult)
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
		for i := 0; i < len(searchResult); i++ {
			id, err = dao.AdminMenu.TopPid(ctx, searchResult[i])
			ids = append(ids, id)
		}
	}

	lst, err = dao.AdminMenu.GenTreeList(ctx, 0, ids)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	_ = gconv.Structs(lst, &res)

	return &res, nil
}

// MaxSort 最大排序
func (s *sAdminMenu) MaxSort(ctx context.Context, req *menu.MaxSortReq) (*menu.MaxSortRes, error) {
	var (
		res menu.MaxSortRes
		err error
	)

	if req.Id > 0 {
		if err = dao.AdminMenu.Ctx(ctx).Where("id", req.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

// NameUnique 菜单名称是否唯一
func (s *sAdminMenu) NameUnique(ctx context.Context, req *menu.NameUniqueReq) (*menu.NameUniqueRes, error) {
	var (
		res menu.NameUniqueRes
		err error
	)

	res.IsUnique, err = dao.AdminMenu.IsUniqueName(ctx, req.Id, req.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

// CodeUnique 菜单编码是否唯一
func (s *sAdminMenu) CodeUnique(ctx context.Context, req *menu.CodeUniqueReq) (*menu.CodeUniqueRes, error) {
	var (
		res menu.CodeUniqueRes
		err error
	)

	res.IsUnique, err = dao.AdminMenu.IsUniqueName(ctx, req.Id, req.Code)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

// Delete 删除
func (s *sAdminMenu) Delete(ctx context.Context, req *menu.DeleteReq) error {
	exist, err := dao.AdminMenu.Ctx(ctx).Where("pid", req.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先删除该菜单下的所有菜单！")
	}
	_, err = dao.AdminMenu.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sAdminMenu) Edit(ctx context.Context, req *menu.EditReq) (err error) {
	var (
		pidData    *do.AdminMenu
		uniqueName bool
		uniqueCode bool
	)

	if req.Title == "" {
		err = gerror.New("菜单名称不能为空")
		return err
	}
	if req.Type != 3 && req.Path == "" {
		err = gerror.New("路由地址不能为空")
		return err
	}
	if req.Name == "" {
		err = gerror.New("路由名称不能为空")
		return err
	}

	uniqueName, err = dao.AdminMenu.IsUniqueTitle(ctx, req.Id, req.Title)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("菜单名称已存在")
		return err
	}

	uniqueCode, err = dao.AdminMenu.IsUniqueName(ctx, req.Id, req.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueCode {
		err = gerror.New("菜单编码已存在")
		return err
	}

	// 维护菜单等级
	if req.Pid == 0 {
		req.Level = 1
	} else {
		if err = dao.AdminMenu.Ctx(ctx).Where("id", req.Pid).Scan(&pidData); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}
		if pidData == nil {
			return gerror.New("上级菜单信息错误")
		}
		req.Level = gconv.Int(pidData.Level) + 1
	}

	// 修改
	req.UpdatedAt = gtime.Now()
	if req.Id > 0 {
		if req.Pid == req.Id {
			return gerror.New("上级菜单不能是当前菜单")
		}
		_, err = dao.AdminMenu.Ctx(ctx).Where("id", req.Id).Data(req).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return casbin.Refresh(ctx)
	}

	// 新增
	req.CreatedAt = gtime.Now()
	_, err = dao.AdminMenu.Ctx(ctx).Data(req).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return casbin.Refresh(ctx)
}

// View 获取信息
func (s *sAdminMenu) View(ctx context.Context, req *menu.ViewReq) (res *menu.ViewRes, err error) {
	if err = dao.AdminMenu.Ctx(ctx).Where("id", req.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取菜单列表
func (s *sAdminMenu) List(ctx context.Context, req *menu.ListReq) (lists []map[string]interface{}, err error) {
	var models []*adminin.MenuTree
	err = dao.AdminMenu.Ctx(ctx).Order("sort asc,id desc").Scan(&models)
	if err != nil {
		return lists, err
	}
	return tree.GenTree(gconv.SliceMap(models)), nil
}

// genNaiveMenus 生成NaiveUI菜单格式
func (s *sAdminMenu) genNaiveMenus(menus []adminin.MenuRouteSummary) (sources []adminin.MenuRoute) {
	for _, men := range menus {
		var source adminin.MenuRoute
		source.Name = men.Name
		source.Path = men.Path
		source.Redirect = men.Redirect
		source.Component = men.Component
		source.Meta = adminin.MenuRouteMeta{
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
func (s *sAdminMenu) getChildrenList(menu *adminin.MenuRouteSummary, treeMap map[string][]adminin.MenuRouteSummary) (err error) {
	menu.Children = treeMap[gconv.String(menu.Id)]
	for i := 0; i < len(menu.Children); i++ {
		err = s.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// GetMenuList 获取菜单列表
func (s *sAdminMenu) GetMenuList(ctx context.Context, memberId int64) (lists role.DynamicRes, err error) {
	var (
		allMenus []adminin.MenuRouteSummary
		menus    []adminin.MenuRouteSummary
		treeMap  = make(map[string][]adminin.MenuRouteSummary)
		mod      = dao.AdminMenu.Ctx(ctx).Where("status", consts.StatusEnabled).WhereIn("type", []int{1, 2})
	)

	// 非超管验证允许的菜单列表
	if !service.AdminMember().VerifySuperId(ctx, memberId) {
		array, err := dao.AdminRoleMenu.Ctx(ctx).
			Fields("menu_id").
			Where("role_id", contexts.GetRoleId(ctx)).
			Array()
		if err != nil {
			return role.DynamicRes{}, err
		}
		mod = mod.Where("id", array)
	}

	if err = mod.Order("sort asc,id desc").Scan(&allMenus); err != nil {
		return lists, err
	}

	if len(allMenus) == 0 {
		return lists, nil
	}

	for _, v := range allMenus {
		treeMap[gconv.String(v.Pid)] = append(treeMap[gconv.String(v.Pid)], v)
	}

	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = s.getChildrenList(&menus[i], treeMap)
	}

	lists.List = append(lists.List, s.genNaiveMenus(menus)...)
	return
}

// LoginPermissions 获取登录成功后的细分权限
func (s *sAdminMenu) LoginPermissions(ctx context.Context, memberId int64) (lists []*adminin.MemberLoginPermissions, err error) {
	// 空跑
	lists = append(lists, &adminin.MemberLoginPermissions{
		Value: "value",
	})

	type Permissions struct {
		Permissions string `json:"permissions"`
	}

	var (
		allPermissions []*Permissions
		mod            = dao.AdminMenu.Ctx(ctx).Fields("permissions").Where("status", consts.StatusEnabled).Where("permissions != ?", "")
	)

	// 非超管验证允许的菜单列表
	if !service.AdminMember().VerifySuperId(ctx, memberId) {
		array, err := dao.AdminRoleMenu.Ctx(ctx).
			Fields("menu_id").
			Where("role_id", contexts.GetRoleId(ctx)).
			Array()
		if err != nil {
			return nil, err
		}
		mod = mod.Where("id", array)
	}

	if err = mod.Scan(&allPermissions); err != nil {
		return lists, err
	}

	if len(allPermissions) == 0 {
		return
	}

	for _, v := range allPermissions {
		for _, p := range gstr.Explode(`,`, v.Permissions) {
			lists = append(lists, &adminin.MemberLoginPermissions{
				Value: p,
			})
		}
	}

	return
}
