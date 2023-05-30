// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/library/location"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/useragent"
)

type sSysLoginLog struct{}

func NewSysLoginLog() *sSysLoginLog {
	return &sSysLoginLog{}
}

func init() {
	service.RegisterSysLoginLog(NewSysLoginLog())
}

// Model 登录日志Orm模型
func (s *sSysLoginLog) Model(ctx context.Context) *gdb.Model {
	return dao.SysLoginLog.Ctx(ctx)
}

// List 获取登录日志列表
func (s *sSysLoginLog) List(ctx context.Context, in sysin.LoginLogListInp) (list []*sysin.LoginLogListModel, totalCount int, err error) {
	mod := dao.SysLoginLog.Ctx(ctx)

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysLoginLog.Columns().Status, in.Status)
	}

	// 查询登录时间
	if len(in.LoginAt) == 2 {
		mod = mod.WhereBetween(dao.SysLoginLog.Columns().LoginAt, in.LoginAt[0], in.LoginAt[1])
	}

	// 查询IP地址
	if in.SysLogIp != "" {
		mod = mod.Where("sysLog."+dao.SysLog.Columns().Ip, in.SysLogIp)
	}

	// 用户名
	if in.Username != "" {
		mod = mod.Where(dao.SysLoginLog.Columns().Username, in.Username)
	}

	// 关联表sysLog
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.SysLoginLog.Table(), dao.SysLoginLog.Columns().ReqId, // 主表表名,关联条件
		dao.SysLog.Table(), "sysLog", dao.SysLog.Columns().ReqId, // 关联表表名,别名,关联条件
	)...)

	totalCount, err = mod.Clone().Count(1)
	if err != nil || totalCount == 0 {
		return
	}

	//关联表select
	fields, err := hgorm.GenJoinSelect(ctx, sysin.LoginLogListModel{}, dao.SysLoginLog, []*hgorm.Join{
		{Dao: dao.SysLog, Alias: "sysLog"},
	})

	if err != nil {
		return
	}

	if err = mod.Fields(fields).Hook(hook.CityLabel).Handler(handler.FilterAuth).Page(in.Page, in.PerPage).OrderDesc(dao.SysLoginLog.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for _, v := range list {
		//// 获取省市编码对应的地区名称
		//region, err := location.ParseRegion(ctx, v.SysLogProvinceId, v.SysLogCityId, 0)
		//if err != nil {
		//	return list, totalCount, err
		//}
		//v.Region = region
		v.Os = useragent.GetOs(v.SysLogUserAgent)
		v.Browser = useragent.GetBrowser(v.SysLogUserAgent)
	}

	return
}

// Export 导出登录日志
func (s *sSysLoginLog) Export(ctx context.Context, in sysin.LoginLogListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.LoginLogExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出登录日志-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.LoginLogExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Delete 删除登录日志
func (s *sSysLoginLog) Delete(ctx context.Context, in sysin.LoginLogDeleteInp) (err error) {
	_, err = dao.SysLoginLog.Ctx(ctx).Where(dao.SysLoginLog.Columns().Id, in.Id).Delete()
	return
}

// View 获取登录日志指定信息
func (s *sSysLoginLog) View(ctx context.Context, in sysin.LoginLogViewInp) (res *sysin.LoginLogViewModel, err error) {
	err = dao.SysLoginLog.Ctx(ctx).Where(dao.SysLoginLog.Columns().Id, in.Id).Scan(&res)
	return
}

// Push 推送登录日志
func (s *sSysLoginLog) Push(ctx context.Context, in sysin.LoginLogPushInp) {
	if in.Response == nil {
		in.Response = new(adminin.LoginModel)
	}
	var models entity.SysLoginLog
	models.ReqId = gctx.CtxId(ctx)
	models.MemberId = in.Response.Id
	models.Username = in.Response.Username
	models.LoginAt = gtime.Now()
	models.LoginIp = location.GetClientIp(ghttp.RequestFromCtx(ctx))
	models.Status = consts.StatusEnabled

	if in.Err != nil {
		models.Status = consts.StatusDisable
		models.ErrMsg = in.Err.Error()
	}

	models.Response = gjson.New(consts.NilJsonToString)
	if in.Response != nil {
		models.Response = gjson.New(in.Response)
	}

	if err := queue.Push(consts.QueueLoginLogTopic, models); err != nil {
		g.Log().Warningf(ctx, "push err:%+v, models:%+v", err, models)
	}
}

// RealWrite 真实写入
func (s *sSysLoginLog) RealWrite(ctx context.Context, models entity.SysLoginLog) (err error) {
	_, err = dao.SysLoginLog.Ctx(ctx).FieldsEx(dao.SysLog.Columns().Id).Data(models).Insert()
	return
}
