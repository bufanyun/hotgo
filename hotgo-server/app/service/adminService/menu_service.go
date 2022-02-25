//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminService

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/bufanyun/hotgo/app/service/internal/dto"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var Menu = new(menu)

type menu struct{}

//
//  @Title  查询角色菜单列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.MenuSearchListRes
//  @Return  error
//
func (service *menu) RoleList(ctx context.Context, in input.MenuRoleListInp) (*input.MenuRoleListModel, error) {

	var (
		mod         = dao.AdminRoleMenu.Ctx(ctx)
		roleMenu    []*entity.AdminRoleMenu
		lst         []*model.LabelTreeMenu
		res         input.MenuRoleListModel
		err         error
		checkedKeys []int64
	)

	// TODO  获取选中菜单ID
	if in.RoleId > 0 {
		mod = mod.Where("role_id", in.RoleId)
	}
	err = mod.Fields().Scan(&roleMenu)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}
	for i := 0; i < len(roleMenu); i++ {
		checkedKeys = append(checkedKeys, roleMenu[i].MenuId)
	}
	res.CheckedKeys = checkedKeys

	// TODO  获取菜单树
	lst, err = dao.AdminMenu.GenLabelTreeList(ctx, 0)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	_ = gconv.Structs(lst, &res.Menus)

	return &res, nil
}

//
//  @Title  查询菜单列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.MenuSearchListRes
//  @Return  error
//
func (service *menu) SearchList(ctx context.Context, req *adminForm.MenuSearchListReq) (*adminForm.MenuSearchListRes, error) {

	var (
		mod          = dao.AdminMenu.Ctx(ctx)
		lst          []*model.TreeMenu
		res          adminForm.MenuSearchListRes
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

//
//  @Title  最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictDataMaxSortRes
//  @Return  error
//
func (service *menu) MaxSort(ctx context.Context, req *adminForm.MenuMaxSortReq) (*adminForm.MenuMaxSortRes, error) {
	var (
		res adminForm.MenuMaxSortRes
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

//
//  @Title  菜单名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *menu) NameUnique(ctx context.Context, req *adminForm.MenuNameUniqueReq) (*adminForm.MenuNameUniqueRes, error) {
	var (
		res adminForm.MenuNameUniqueRes
		err error
	)

	res.IsUnique, err = dao.AdminMenu.IsUniqueName(ctx, req.Id, req.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

//
//  @Title  菜单编码是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *menu) CodeUnique(ctx context.Context, req *adminForm.MenuCodeUniqueReq) (*adminForm.MenuCodeUniqueRes, error) {
	var (
		res adminForm.MenuCodeUniqueRes
		err error
	)

	res.IsUnique, err = dao.AdminMenu.IsUniqueCode(ctx, req.Id, req.Code)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

//
//  @Title  删除
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *menu) Delete(ctx context.Context, req *adminForm.MenuDeleteReq) error {

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

//
//  @Title  修改/新增
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *menu) Edit(ctx context.Context, req *adminForm.MenuEditReq) (err error) {
	var (
		pidData    *dto.AdminMenu
		uniqueName bool
		uniqueCode bool
	)

	if req.Name == "" {
		err = gerror.New("菜单名称不能为空")
		return err
	}
	if req.Path == "" {
		err = gerror.New("菜单路径不能为空")
		return err
	}
	if req.Code == "" {
		err = gerror.New("菜单编码不能为空")
		return err
	}

	uniqueName, err = dao.AdminMenu.IsUniqueName(ctx, req.Id, req.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("菜单名称已存在")
		return err
	}

	uniqueCode, err = dao.AdminMenu.IsUniqueCode(ctx, req.Id, req.Code)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueCode {
		err = gerror.New("菜单编码已存在")
		return err
	}

	// TODO  维护菜单等级
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

		return nil
	}

	// 新增
	req.CreatedAt = gtime.Now()
	_, err = dao.AdminMenu.Ctx(ctx).Data(req).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

//
//  @Title  获取信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeViewRes
//  @Return  error
//
func (service *menu) View(ctx context.Context, req *adminForm.MenuViewReq) (res *adminForm.MenuViewRes, err error) {
	//var (
	//	res adminForm.MenuViewRes
	//)

	if err = dao.AdminMenu.Ctx(ctx).Where("id", req.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

//
//  @Title  获取菜单列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *menu) List(ctx context.Context, req *adminForm.MenuListReq) (*adminForm.MenuListRes, error) {
	var (
		m          = dao.AdminMenu.Ctx(ctx)
		list       []*entity.AdminMenu
		res        adminForm.MenuListRes
		totalCount int
		err        error
	)

	if req.Pid == 0 {
		m = m.Where("level", 1)
	} else {
		m = m.Where("pid", req.Pid)
	}

	totalCount, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	err = m.Page(req.Page, req.Limit).Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.List = list
	res.Page = req.Page
	res.Limit = req.Limit
	res.TotalCount = totalCount

	return &res, nil
}

type RelationTree struct {
	adminForm.RoleDynamicBase
	Children []*RelationTree `json:"children"`
}

//
//  @Title  获取菜单列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   member_id
//
func (service *menu) GetMenuList(ctx context.Context, member_id int64) (lists *adminForm.RoleDynamicRes, err error) {

	var (
		results       []*entity.AdminMenu
		models        []*RelationTree
		recursion     []*adminForm.RoleDynamicBase
		finalResponse adminForm.RoleDynamicRes
	)

	err = dao.AdminMenu.Ctx(ctx).Order("sort asc,id desc").Scan(&results)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	for i := 0; i < len(results); i++ {

		// 元数据
		var (
			meta adminForm.RoleDynamicMeta
			rec  adminForm.RoleDynamicBase
		)

		meta.Title = results[i].Name
		meta.Icon = results[i].Icon
		meta.NoCache = gconv.Bool(results[i].IsCache)
		meta.Remark = results[i].Remark

		rec.Id = results[i].Id
		rec.Pid = results[i].Pid
		rec.IsFrame = results[i].IsFrame
		rec.Name = results[i].Name
		rec.Code = results[i].Code
		rec.Path = results[i].Path
		rec.Hidden = results[i].IsVisible == "1"
		rec.Redirect = service.getRedirect(results[i])
		rec.Component = service.getComponent(results[i])
		rec.AlwaysShow = true
		rec.Meta = &meta

		recursion = append(recursion, &rec)
	}

	_ = gconv.Structs(recursion, &models)

	childIds := service.getChildIds(ctx, models, 0)

	_ = gconv.Structs(childIds, &finalResponse)

	return &finalResponse, nil
}

//
//  @Title  获取菜单的组件配置
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   model
//  @Return  string
//
func (service *menu) getComponent(mod *entity.AdminMenu) string {

	if mod.Type == "M" {
		return "Layout"
	}

	if mod.Type == "C" {
		return mod.Component
	}

	return mod.Component
}

//
//  @Title  获取菜单是否重定向
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   model
//  @Return  string
//
func (service *menu) getRedirect(model *entity.AdminMenu) string {
	if model.Type == "M" {
		return "noRedirect"
	}

	return ""
}

//
//  @Title  将菜单转为父子关系菜单
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   lists
//  @Param   pid
//  @Return  []*RelationTree
//
func (service *menu) getChildIds(ctx context.Context, lists []*RelationTree, pid int64) []*RelationTree {

	var (
		count    = len(lists)
		newLists []*RelationTree
	)

	if count == 0 {
		return nil
	}

	for i := 0; i < len(lists); i++ {
		if lists[i].Id > 0 && lists[i].Pid == pid {
			var row *RelationTree
			if err := gconv.Structs(lists[i], &row); err != nil {
				panic(err)
			}
			row.Children = service.getChildIds(ctx, lists, row.Id)
			newLists = append(newLists, row)
		}
	}

	return newLists
}

//
//  @Title  根据条件查询一行的数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   where
//  @Return  *entity.AdminMenu
//
func (service *menu) WhereScan(ctx context.Context, where dto.AdminMenu) *entity.AdminMenu {
	var (
		mod *entity.AdminMenu
		err error
	)

	if err = dao.AdminMenu.Ctx(ctx).Where(where).Scan(&mod); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil
	}

	return mod
}
