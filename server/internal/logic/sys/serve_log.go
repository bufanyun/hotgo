// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sSysServeLog struct{}

func NewSysServeLog() *sSysServeLog {
	return &sSysServeLog{}
}

func init() {
	service.RegisterSysServeLog(NewSysServeLog())
}

// Model 服务日志Orm模型
func (s *sSysServeLog) Model(ctx context.Context) *gdb.Model {
	return dao.SysServeLog.Ctx(ctx)
}

// List 获取服务日志列表
func (s *sSysServeLog) List(ctx context.Context, in sysin.ServeLogListInp) (list []*sysin.ServeLogListModel, totalCount int, err error) {
	mod := dao.SysServeLog.Ctx(ctx)

	// 查询链路ID
	if in.TraceId != "" {
		mod = mod.Where(dao.SysServeLog.Columns().TraceId, in.TraceId)
	}

	// 查询日志级别
	if in.LevelFormat != "" {
		mod = mod.WhereLike(dao.SysServeLog.Columns().LevelFormat, in.LevelFormat)
	}

	// 查询触发时间(ns)
	if len(in.TriggerNs) == 2 {
		mod = mod.WhereBetween(dao.SysServeLog.Columns().TriggerNs, in.TriggerNs[0], in.TriggerNs[1])
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysServeLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 关联表sysLog
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.SysServeLog.Table(), dao.SysServeLog.Columns().TraceId, // 主表表名,关联条件
		dao.SysLog.Table(), "sysLog", dao.SysLog.Columns().ReqId, // 关联表表名,别名,关联条件
	)...)

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	//关联表select
	fields, err := hgorm.GenJoinSelect(ctx, sysin.ServeLogListModel{}, dao.SysServeLog, []*hgorm.Join{
		{Dao: dao.SysLog, Alias: "sysLog"},
	})
	if err = mod.Fields(fields).Handler(handler.FilterAuth).Page(in.Page, in.PerPage).OrderDesc(dao.SysServeLog.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	return list, totalCount, err
}

// Export 导出服务日志
func (s *sSysServeLog) Export(ctx context.Context, in sysin.ServeLogListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return err
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.ServeLogExportModel{})
	if err != nil {
		return err
	}

	var (
		fileName  = "导出服务日志-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.ServeLogExportModel
	)

	err = gconv.Scan(list, &exports)
	if err != nil {
		return err
	}
	if err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName); err != nil {
		return
	}
	return
}

// Delete 删除服务日志
func (s *sSysServeLog) Delete(ctx context.Context, in sysin.ServeLogDeleteInp) (err error) {
	_, err = dao.SysServeLog.Ctx(ctx).Where(dao.SysServeLog.Columns().Id, in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// View 获取服务日志指定信息
func (s *sSysServeLog) View(ctx context.Context, in sysin.ServeLogViewInp) (res *sysin.ServeLogViewModel, err error) {
	if err = dao.SysServeLog.Ctx(ctx).Where(dao.SysServeLog.Columns().Id, in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// RealWrite 真实写入
func (s *sSysServeLog) RealWrite(ctx context.Context, models entity.SysServeLog) (err error) {
	_, err = dao.SysServeLog.Ctx(ctx).Data(models).Insert()
	return
}
