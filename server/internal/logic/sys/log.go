// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/dict"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/library/location"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/excel"
	"hotgo/utility/simple"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
)

type sSysLog struct{}

func NewSysLog() *sSysLog {
	return &sSysLog{}
}

func init() {
	service.RegisterSysLog(NewSysLog())
}

// Model 请求日志Orm模型
func (s *sSysLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysLog.Ctx(ctx), option...)
}

// Export 导出
func (s *sSysLog) Export(ctx context.Context, in *sysin.LogListInp) (err error) {
	//  导出格式
	type exportImage struct {
		Id         int64       `json:"id"           description:""`
		AppId      string      `json:"app_id"       description:"应用id"`
		Method     string      `json:"method"       description:"提交类型"`
		Module     string      `json:"module"       description:"模块"`
		Url        string      `json:"url"          description:"提交url"`
		Ip         string      `json:"ip"           description:"ip地址"`
		ErrorCode  int         `json:"error_code"   description:"报错code"`
		ErrorMsg   string      `json:"error_msg"    description:"报错信息"`
		ReqId      string      `json:"req_id"       description:"对外id"`
		TakeUpTime int64       `json:"take_up_time" description:"请求耗时"`
		CreatedAt  *gtime.Time `json:"created_at"   description:"创建时间"`
		MemberName string      `json:"memberName"`
		Region     string      `json:"region"`
	}

	var (
		titleList  = []string{"ID", "应用", "提交类型", "模块", "提交url", "ip地址", "报错code", "报错信息", "对外id", "请求耗时", "创建时间", "用户", "访问地"}
		fileName   = "访问日志导出-" + gctx.CtxId(ctx)
		sheetName  = simple.AppName(ctx)
		exportList []exportImage
		row        exportImage
	)

	list, _, err := s.List(ctx, in)
	if err != nil {
		return err
	}

	// 格式化格式
	for i := 0; i < len(list); i++ {
		row.Id = list[i].Id
		row.AppId = list[i].AppId
		row.Module = list[i].Module
		row.Method = list[i].Method
		row.Url = list[i].Url
		row.Ip = list[i].Ip
		row.ReqId = list[i].ReqId
		row.ErrorCode = list[i].ErrorCode
		row.ErrorMsg = list[i].ErrorMsg
		row.TakeUpTime = list[i].TakeUpTime
		row.CreatedAt = list[i].CreatedAt
		row.MemberName = list[i].MemberName
		row.Region = list[i].Region
		exportList = append(exportList, row)
	}

	err = excel.ExportByStructs(ctx, titleList, exportList, fileName, sheetName)
	return
}

// RealWrite 真实写入
func (s *sSysLog) RealWrite(ctx context.Context, log entity.SysLog) (err error) {
	_, err = dao.SysLog.Ctx(ctx).FieldsEx(dao.SysLog.Columns().Id).Data(log).Unscoped().OmitEmptyData().Insert()
	return
}

// AutoLog 根据配置自动记录请求日志
func (s *sSysLog) AutoLog(ctx context.Context) error {
	return g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				g.Log().Error(ctx, "autoLog err:%+v", err)
			}
		}()

		config, err := service.SysConfig().GetLoadLog(ctx)
		if err != nil {
			return
		}

		if config == nil || !config.Switch {
			return
		}

		data := s.AnalysisLog(ctx)
		if ok := validate.InSliceExistStr(config.Module, data.Module); !ok {
			return
		}

		if ok := validate.InSliceExistStr(config.SkipCode, gconv.String(data.ErrorCode)); ok {
			return
		}

		if config.Queue {
			err = queue.Push(consts.QueueLogTopic, data)
			return
		}

		err = s.RealWrite(ctx, data)
	})
}

// AnalysisLog 解析日志数据
func (s *sSysLog) AnalysisLog(ctx context.Context) entity.SysLog {
	var (
		mctx       = contexts.Get(ctx)
		response   = mctx.Response
		user       = mctx.User
		request    = ghttp.RequestFromCtx(ctx)
		module     = mctx.Module
		clientIp   = location.GetClientIp(request)
		postData   = gjson.New(consts.NilJsonToString)
		getData    = gjson.New(request.URL.Query())
		headerData = gjson.New(consts.NilJsonToString)
		errorData  = gjson.New(consts.NilJsonToString)
		data       entity.SysLog
		memberId   int64
		errorCode  int
		errorMsg   string
		traceID    string
		timestamp  int64
		appId      string
		takeUpTime int64
	)

	// 响应数据
	if response != nil {
		errorCode = response.Code
		errorMsg = response.Message
		traceID = response.TraceID
		timestamp = response.Timestamp
		if len(gconv.String(response.Error)) > 0 {
			errorData = gjson.New(response.Error)
		}
	}

	if timestamp == 0 {
		timestamp = gtime.Timestamp()
	}

	// 请求头
	if reqHeadersBytes, _ := gjson.New(request.Header).MarshalJSON(); len(reqHeadersBytes) > 0 {
		headerData = gjson.New(reqHeadersBytes)
	}

	// post参数
	if body, ok := mctx.Data["request.body"].(*gjson.Json); ok {
		postData = body
	}

	// post表单
	postForm := gjson.New(gconv.String(request.PostForm)).Map()
	if len(postForm) > 0 {
		for k, v := range postForm {
			postData.MustSet(k, v)
		}
	}

	if postData.IsNil() || len(postData.Map()) == 0 {
		postData = gjson.New(consts.NilJsonToString)
	} else {
		// 隐藏明文显示的密码
		for k := range postData.Map() {
			if gstr.ContainsI(k, "password") {
				postData.MustSet(k, "******")
			}
		}
	}

	// 当前登录用户
	if user != nil {
		memberId = user.Id
		appId = user.App
	}

	ipData, err := location.GetLocation(ctx, clientIp)
	if err != nil {
		g.Log().Debugf(ctx, "location.GetLocation clientIp:%v, err:%+v", clientIp, err)
	}

	if ipData == nil {
		ipData = new(location.IpLocationData)
	}

	// 请求耗时
	if tt, ok := mctx.Data["request.takeUpTime"].(int64); ok {
		takeUpTime = tt
	}

	headerData.MustAppend("Enter-Time", request.EnterTime.String())

	data = entity.SysLog{
		AppId:      appId,
		MerchantId: 0,
		MemberId:   memberId,
		Method:     request.Method,
		Module:     module,
		Url:        request.URL.Path,
		GetData:    getData,
		PostData:   postData,
		HeaderData: s.SimplifyHeaderParams(headerData),
		Ip:         clientIp,
		ProvinceId: ipData.ProvinceCode,
		CityId:     ipData.CityCode,
		ErrorCode:  errorCode,
		ErrorMsg:   errorMsg,
		ErrorData:  errorData,
		ReqId:      traceID,
		Timestamp:  timestamp,
		UserAgent:  request.Header.Get("User-Agent"),
		Status:     consts.StatusEnabled,
		TakeUpTime: takeUpTime,
		UpdatedAt:  gtime.Now(),
		CreatedAt:  request.EnterTime,
	}
	return data
}

// SimplifyHeaderParams 过滤掉请求头中的大参数
func (s *sSysLog) SimplifyHeaderParams(data *gjson.Json) *gjson.Json {
	if data == nil || data.IsNil() {
		return data
	}
	var filters = []string{"Accept", "Authorization", "Cookie"}
	for _, filter := range filters {
		v := data.Get(filter)
		if len(v.String()) > 128 {
			data.MustRemove(filter)
		}
	}
	return data
}

// View 获取指定请求日志信息
func (s *sSysLog) View(ctx context.Context, in *sysin.LogViewInp) (res *sysin.LogViewModel, err error) {
	mod := s.Model(ctx)

	count, err := service.SysLoginLog().Model(ctx).
		LeftJoinOnFields(dao.SysLog.Table(), dao.SysLoginLog.Columns().ReqId, "=", dao.SysLog.Columns().ReqId).
		WherePrefix(dao.SysLog.Table(), dao.SysLog.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, err
	}

	if count > 0 {
		mod = dao.SysLog.Ctx(ctx)
	}

	if err = mod.Hook(hook.CityLabel).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if res == nil {
		return
	}

	routes := global.LoadHTTPRoutes(ghttp.RequestFromCtx(ctx))
	key := global.GenRouteKey(res.Method, res.Url)
	route, ok := routes[key]
	if ok {
		res.Tags = route.Tags
		res.Summary = route.Summary
		res.Description = route.Description
	}

	if simple.IsDemo(ctx) {
		res.HeaderData = gjson.New(`{
		   "none": [
		       "` + consts.DemoTips + `"
		   ]
		}`)
	}
	return
}

// Delete 删除请求日志
func (s *sSysLog) Delete(ctx context.Context, in *sysin.LogDeleteInp) (err error) {
	_, err = s.Model(ctx).WherePri(in.Id).Delete()
	return
}

// List 请求日志列表
func (s *sSysLog) List(ctx context.Context, in *sysin.LogListInp) (list []*sysin.LogListModel, totalCount int, err error) {
	mod := s.Model(ctx).FieldsEx("get_data", "header_data", "post_data")

	// 访问路径
	if in.Url != "" {
		mod = mod.WhereLike("url", "%"+in.Url+"%")
	}

	// 模块
	if in.Module != "" {
		mod = mod.Where("module", in.Module)
	}

	// 链路ID
	if in.ReqId != "" {
		mod = mod.Where("req_id", in.ReqId)
	}

	// 请求方式
	if in.Method != "" {
		mod = mod.Where("method", in.Method)
	}

	// 用户
	if in.MemberId > 0 {
		mod = mod.Where("member_id", in.MemberId)
	}

	// 操作人筛选
	if len(in.ComplexMemberId) == 2 && len(in.ComplexMemberId[0]) > 0 {
		memberIds, err := service.AdminMember().GetComplexMemberIds(ctx, in.ComplexMemberId[0], in.ComplexMemberId[1])
		if err != nil {
			return nil, 0, err
		}
		if len(memberIds) == 0 {
			return nil, 0, nil
		}
		mod = mod.WhereIn("member_id", memberIds)
	}

	// 访问IP
	if in.Ip != "" {
		mod = mod.Where("ip", in.Ip)
	}

	// 日期范围
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("created_at", gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	// 状态码
	if in.ErrorCode != "" {
		mod = mod.Where("error_code", in.ErrorCode)
	}

	// 请求耗时
	if dict.HasOptionKey(consts.HTTPHandlerTimeOptions, in.TakeUpTime) {
		mod = mod.Where(gdb.Raw(fmt.Sprintf("(`take_up_time` %v)", in.TakeUpTime)))
	}

	// 非生产环境，允许关键词查询日志
	// 生成环境使用需谨慎，日志量大易产生慢日志
	if !gmode.IsProduct() && in.Keyword != "" {
		filterColumns := map[string]string{
			"get_data":    "LIKE",
			"post_data":   "LIKE",
			"header_data": "LIKE",
			"error_data":  "LIKE",
			"error_msg":   "LIKE",
		}
		mod = hgorm.FilterKeywordsWithOr(mod, filterColumns, in.Keyword)
	}

	totalCount, err = mod.Count()
	if err != nil || totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Hook(hook.CityLabel).Order("id desc").Scan(&list); err != nil {
		return
	}

	routes := global.LoadHTTPRoutes(ghttp.RequestFromCtx(ctx))
	for _, v := range list {
		if v.AppId == consts.AppAdmin {
			memberName, err := dao.AdminMember.Ctx(ctx).Fields(dao.AdminMember.Columns().RealName).WherePri(v.MemberId).Value()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return list, totalCount, err
			}
			v.MemberName = memberName.String()
		}

		if v.MemberName == "" {
			v.MemberName = "游客"
		}

		// 截取请求url路径
		if gstr.Contains(v.Url, "?") {
			v.Url = gstr.StrTillEx(v.Url, "?")
		}

		key := global.GenRouteKey(v.Method, v.Url)
		route, ok := routes[key]
		if ok {
			v.Tags = route.Tags
			v.Summary = route.Summary
			v.Description = route.Description
		}

		if simple.IsDemo(ctx) {
			v.HeaderData = gjson.New(`{
			   "none": [
			       "` + consts.DemoTips + `"
			   ]
			}`)
		}
	}
	return
}
