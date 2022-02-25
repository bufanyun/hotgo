//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysService

import (
	"context"
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/adminForm"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/bufanyun/hotgo/app/utils"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

var Dict = new(dict)

type dict struct{}

//
//  @Title  数据键值是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *dict) DataUnique(ctx context.Context, req *adminForm.DictDataUniqueReq) (*adminForm.DictDataUniqueRes, error) {
	var (
		res adminForm.DictDataUniqueRes
		err error
	)

	res.IsUnique, err = dao.SysDictData.IsUnique(ctx, req.Id, req.Type, req.Value)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

//
//  @Title  查询字典数据最大排序
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictDataMaxSortRes
//  @Return  error
//
func (service *dict) DataMaxSort(ctx context.Context, req *adminForm.DictDataMaxSortReq) (*adminForm.DictDataMaxSortRes, error) {
	var (
		m   = dao.SysDictData.Ctx(ctx).Where("type", req.Type).Order("sort desc")
		res adminForm.DictDataMaxSortRes
		err error
	)

	if err = m.Scan(&res); err != nil && err.Error() != "sql: no rows in result set" {

		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.Sort = res.Sort + 10

	return &res, nil
}

//
//  @Title  删除字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *dict) DataDelete(ctx context.Context, req *adminForm.DictDataDeleteReq) error {
	var (
		m   = dao.SysDictData.Ctx(ctx).Where("id", req.Id)
		err error
	)

	_, err = m.Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

//
//  @Title  修改/新增字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *dict) DataEdit(ctx context.Context, req *adminForm.DictDataEditReq) error {
	var (
		m        = dao.SysDictData.Ctx(ctx)
		isUnique bool
		err      error
	)

	if req.Label == "" {
		err = gerror.New("字典标签不能为空")
		return err
	}
	if req.Type == "" {
		err = gerror.New("字典类型不能为空")
		return err
	}
	if req.Value == "" {
		err = gerror.New("字典键值不能为空")
		return err
	}

	isUnique, err = dao.SysDictData.IsUnique(ctx, req.Id, req.Type, req.Value)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if !isUnique {
		err = gerror.New("字典键值已存在")
		return err
	}

	req.UpdatedAt = gtime.Now()

	// 修改
	if req.Id > 0 {
		_, err = m.Where("id", req.Id).Data(req).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	req.CreatedAt = gtime.Now()

	// 新增
	_, err = m.Data(req).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

//
//  @Title  获取指定字典数据信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeViewRes
//  @Return  error
//
func (service *dict) DataView(ctx context.Context, req *adminForm.DictDataViewReq) (*adminForm.DictDataViewRes, error) {
	var (
		m   = dao.SysDictData.Ctx(ctx).Where("id", req.Id)
		res adminForm.DictDataViewRes
		err error
	)

	if err = m.Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

//
//  @Title  获取指定字典类型的属性数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictAttributeRes
//  @Return  error
//
func (service *dict) Attribute(ctx context.Context, req *adminForm.DictAttributeReq) (*adminForm.DictAttributeRes, error) {
	var (
		m   = dao.SysDictData.Ctx(ctx).Where("type", req.Type).Order("sort asc,id desc")
		res adminForm.DictAttributeRes
		err error
	)

	if err = m.Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

//
//  @Title  获取字典数据列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *dict) DataList(ctx context.Context, req *adminForm.DictDataListReq) (*adminForm.DictDataListRes, error) {
	var (
		m          = dao.SysDictData.Ctx(ctx).Where("type", req.Type)
		list       []*entity.SysDictData
		res        adminForm.DictDataListRes
		totalCount int
		err        error
	)

	totalCount, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	if err = m.Page(req.Page, req.Limit).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.List = list
	res.Page = req.Page
	res.Limit = req.Limit
	res.TotalCount = totalCount

	return &res, nil
}

//
//  @Title  导出字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictDataListRes
//  @Return  error
//
func (service *dict) TypeExport(ctx context.Context, req *adminForm.DictTypeExportReq) error {

	//  导出格式
	type exportImage struct {
		Id        int64  `json:"id" `
		Name      string `json:"name" `
		Type      string `json:"type" `
		Remark    string `json:"remark" `
		Status    string `json:"status" `
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	var (
		list      []exportImage
		titleList = []string{"ID", "字典名称", "字典类型", "备注", "状态", "创建时间", "更新时间"}
		fileName  = "字典类型导出-" + com.Context.Get(ctx).ReqId + ".xlsx"
		sheetName = "HotGo"
		err       error
	)

	if err = dao.SysDictType.Ctx(ctx).Page(req.Page, req.Limit).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	// TODO  格式化格式
	for i := 0; i < len(list); i++ {
		if list[i].Status == consts.StatusEnabled {
			list[i].Status = "启用"
		} else if list[i].Status == consts.StatusDisable {
			list[i].Status = "禁用"
		} else if list[i].Status == consts.StatusDelete {
			list[i].Status = "已删除"
		}
	}

	// TODO  强转类型
	writer := com.Context.Get(ctx).Request.Response.Writer
	w, _ := interface{}(writer).(*ghttp.ResponseWriter)

	g.Log().Print(ctx, "gconv.Interfaces(list):", gconv.Interfaces(list))
	if err = utils.Excel.ExportByStruct(w, titleList, gconv.Interfaces(list), fileName, sheetName); err != nil {
		err = gerror.Wrap(err, "ExportByStruct:")
		return err
	}

	// TODO  加入到上下文
	com.Context.SetResponse(ctx, &model.Response{
		Code:      consts.CodeOK,
		Message:   "导出成功",
		Timestamp: time.Now().Unix(),
		ReqId:     com.Context.Get(ctx).ReqId,
	})

	return nil
}

//
//  @Title  删除字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *dict) TypeDelete(ctx context.Context, req *adminForm.DictTypeDeleteReq) error {
	var (
		m   = dao.SysDictType.Ctx(ctx).Where("id", req.Id)
		err error
	)

	_, err = m.Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

//
//  @Title  修改/新增字典类型
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  error
//
func (service *dict) TypeEdit(ctx context.Context, req *adminForm.DictTypeEditReq) error {
	var (
		m        = dao.SysDictType.Ctx(ctx)
		isUnique bool
		err      error
	)

	if req.Name == "" {
		err = gerror.New("字典名称不能为空")
		return err
	}
	if req.Type == "" {
		err = gerror.New("字典类型不能为空")
		return err
	}

	isUnique, err = dao.SysDictType.IsUnique(ctx, req.Id, req.Type)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if !isUnique {
		err = gerror.New("字典类型已存在")
		return err
	}

	req.UpdatedAt = gtime.Now()

	// 修改
	if req.Id > 0 {
		_, err = m.Where("id", req.Id).Data(req).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	req.CreatedAt = gtime.Now()

	// 新增
	_, err = m.Where("id", req.Id).Data(req).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	return nil
}

//
//  @Title  类型是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  *adminForm.DictTypeUniqueRes
//  @Return  error
//
func (service *dict) TypeUnique(ctx context.Context, req *adminForm.DictTypeUniqueReq) (*adminForm.DictTypeUniqueRes, error) {
	var (
		res adminForm.DictTypeUniqueRes
		err error
	)

	res.IsUnique, err = dao.SysDictType.IsUnique(ctx, req.Id, req.Type)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

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
func (service *dict) TypeView(ctx context.Context, req *adminForm.DictTypeViewReq) (*adminForm.DictTypeViewRes, error) {
	var (
		m   = dao.SysDictType.Ctx(ctx).Where("id", req.Id)
		res adminForm.DictTypeViewRes
		err error
	)

	if err = m.Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return &res, nil
}

//
//  @Title  获取字典类型列表
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (service *dict) TypeList(ctx context.Context, req *adminForm.DictTypeListReq) (*adminForm.DictTypeListRes, error) {
	var (
		m          = dao.SysDictType.Ctx(ctx)
		list       []*entity.SysDictType
		res        adminForm.DictTypeListRes
		totalCount int
		err        error
	)

	if req.Name != "" {
		m = m.WhereLike("name", "%"+req.Name+"%")
	}

	if req.Type != "" {
		m = m.Where("type", req.Type)
	}

	// 日期范围
	if req.StartTime != "" {
		m = m.WhereGTE("created_at", req.StartTime)
	}
	if req.EndTime != "" {
		m = m.WhereLTE("created_at", req.EndTime)
	}

	// 状态
	if req.Status > 0 {
		m = m.Where("status", req.Status)
	}

	totalCount, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	if err = m.Page(req.Page, req.Limit).Order("sort asc,id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	res.List = list
	res.Page = req.Page
	res.Limit = req.Limit
	res.TotalCount = totalCount

	return &res, nil
}
