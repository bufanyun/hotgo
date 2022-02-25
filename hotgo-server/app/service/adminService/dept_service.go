package adminService

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var Dept = dept{}

type dept struct{}

//
//  @Title  菜单名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *dept) NameUnique(ctx context.Context, in input.AdminDeptNameUniqueInp) (*input.AdminDeptNameUniqueModel, error) {

	var res input.AdminDeptNameUniqueModel
	isUnique, err := dao.AdminDept.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.IsUnique = isUnique
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
func (service *dept) Delete(ctx context.Context, in input.AdminDeptDeleteInp) error {

	exist, err := dao.AdminRoleDept.Ctx(ctx).Where("dept_id", in.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("请先解除该部门下所有已关联用户关联关系！")
	}
	_, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Delete()
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
func (service *dept) Edit(ctx context.Context, in input.AdminDeptEditInp) (err error) {

	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return err
	}

	uniqueName, err := dao.AdminDept.IsUniqueName(ctx, in.Id, in.Name)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if !uniqueName {
		err = gerror.New("名称已存在")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.AdminDept.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
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
func (service *dept) MaxSort(ctx context.Context, in input.AdminDeptMaxSortInp) (*input.AdminDeptMaxSortModel, error) {
	var res input.AdminDeptMaxSortModel

	if in.Id > 0 {
		if err := dao.AdminDept.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

//
//  @Title  获取指定字典类型信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeViewRes
//  @Return  error
//
func (service *dept) View(ctx context.Context, in input.AdminDeptViewInp) (res *input.AdminDeptViewModel, err error) {

	if err = dao.AdminDept.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

//
//  @Title  获取列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *dept) List(ctx context.Context, in input.AdminDeptListInp) (list []*input.AdminDeptListModel, err error) {

	mod := dao.AdminDept.Ctx(ctx)

	var (
		dataList []*entity.AdminDept
		models   []*DeptTree
		//searchResult []*entity.AdminDept
		//id           int64
		//ids          []int64
	)

	// 部门名称
	if in.Name != "" {
		//err = dao.AdminDept.Ctx(ctx).WhereLike("name", "%"+in.Name+"%").Scan(&searchResult)
		//if err != nil {
		//	err = gerror.Wrap(err, consts.ErrorORM)
		//	return nil, err
		//}
		//for i := 0; i < len(searchResult); i++ {
		//	id, err = dao.AdminDept.TopPid(ctx, searchResult[i])
		//	ids = append(ids, id)
		//}
		//
		//if len(ids) == 0 {
		//	return nil, nil
		//}
		//mod = mod.Where("id", ids)
	}

	err = mod.Order("id desc").Scan(&dataList)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	_ = gconv.Structs(dataList, &models)

	childIds := service.getDeptChildIds(ctx, models, 0)

	_ = gconv.Structs(childIds, &list)

	return list, nil
}

type DeptTree struct {
	entity.AdminDept
	Children []*DeptTree `json:"children"`
}

//
//  @Title  将列表转为父子关系列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   lists
//  @Param   pid
//  @Return  []*RelationTree
//
func (service *dept) getDeptChildIds(ctx context.Context, lists []*DeptTree, pid int64) []*DeptTree {

	var (
		count    = len(lists)
		newLists []*DeptTree
	)

	if count == 0 {
		return nil
	}

	for i := 0; i < len(lists); i++ {
		if lists[i].Id > 0 && lists[i].Pid == pid {
			var row *DeptTree
			if err := gconv.Structs(lists[i], &row); err != nil {
				panic(err)
			}
			row.Children = service.getDeptChildIds(ctx, lists, row.Id)
			newLists = append(newLists, row)
		}
	}

	return newLists
}

type DeptListTree struct {
	Id       int64           `json:"id" `
	Key      int64           `json:"key" `
	Pid      int64           `json:"pid"  `
	Label    string          `json:"label"`
	Title    string          `json:"title"`
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Children []*DeptListTree `json:"children"`
}

//
//  @Title  获取列表树
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *dept) ListTree(ctx context.Context, in input.AdminDeptListTreeInp) (list []*input.AdminDeptListTreeModel, err error) {

	mod := dao.AdminDept.Ctx(ctx)

	var (
		dataList []*entity.AdminDept
		models   []*DeptListTree
	)

	err = mod.Order("id desc").Scan(&dataList)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, err
	}

	_ = gconv.Structs(dataList, &models)

	// TODO  重写树入参
	for i := 0; i < len(models); i++ {
		models[i].Key = models[i].Id
		models[i].Title = models[i].Name
		models[i].Label = models[i].Name
	}

	childIds := service.getDeptTreeChildIds(ctx, models, 0)

	_ = gconv.Structs(childIds, &list)

	return list, nil
}

//
//  @Title  将列表转为父子关系列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   lists
//  @Param   pid
//  @Return  []*RelationTree
//
func (service *dept) getDeptTreeChildIds(ctx context.Context, lists []*DeptListTree, pid int64) []*DeptListTree {

	var (
		count    = len(lists)
		newLists []*DeptListTree
	)

	if count == 0 {
		return nil
	}

	for i := 0; i < len(lists); i++ {
		if lists[i].Id > 0 && lists[i].Pid == pid {
			var row *DeptListTree
			if err := gconv.Structs(lists[i], &row); err != nil {
				panic(err)
			}
			row.Children = service.getDeptTreeChildIds(ctx, lists, row.Id)
			newLists = append(newLists, row)
		}
	}

	return newLists
}

//
//  @Title  获取部门名称
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   id
//  @Return  name
//  @Return  err
//
func (service *dept) GetName(ctx context.Context, id int64) (name string, err error) {

	var data entity.AdminDept

	err = dao.AdminDept.Ctx(ctx).
		Where("id", id).
		Fields("name").
		Scan(&data)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return name, err
	}

	return data.Name, nil
}
