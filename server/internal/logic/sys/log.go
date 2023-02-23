// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/library/location"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/excel"
	"hotgo/utility/validate"
)

type sSysLog struct{}

func NewSysLog() *sSysLog {
	return &sSysLog{}
}

func init() {
	service.RegisterSysLog(NewSysLog())
}

// Export 导出
func (s *sSysLog) Export(ctx context.Context, in sysin.LogListInp) (err error) {
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
		fileName   = "访问日志导出-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName  = "HotGo"
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

	if err = excel.ExportByStructs(ctx, titleList, exportList, fileName, sheetName); err != nil {
		return err
	}

	return
}

// RealWrite 真实写入
func (s *sSysLog) RealWrite(ctx context.Context, commonLog entity.SysLog) (err error) {
	_, err = dao.SysLog.Ctx(ctx).Data(commonLog).Insert()
	return
}

// AutoLog 根据配置自动记录请求日志
func (s *sSysLog) AutoLog(ctx context.Context) error {
	return g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				panic(err)
			}
		}()

		config, err := service.SysConfig().GetLoadLog(ctx)
		if err != nil {
			return
		}

		if !config.Switch {
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
		return
	})
}

// AnalysisLog 解析日志数据
func (s *sSysLog) AnalysisLog(ctx context.Context) entity.SysLog {
	var (
		modelContext = contexts.Get(ctx)
		response     = modelContext.Response
		user         = modelContext.User
		request      = ghttp.RequestFromCtx(ctx)
		module       = modelContext.Module
		clientIp     = location.GetClientIp(request)
		postData     = gjson.New(request.GetBodyString())
		getData      = gjson.New(request.URL.Query())
		headerData   = gjson.New(consts.NilJsonToString)
		errorData    = gjson.New(consts.NilJsonToString)
		data         entity.SysLog
		memberId     int64
		errorCode    int
		errorMsg     string
		traceID      string
		timestamp    int64
		appId        string
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

	// 请求头
	if reqHeadersBytes, _ := json.Marshal(request.Header); len(gconv.String(reqHeadersBytes)) > 0 {
		headerData = gjson.New(reqHeadersBytes)
	}

	// post参数
	postForm := gjson.New(gconv.String(request.PostForm)).Map()
	if len(postForm) > 0 {
		for k, v := range postForm {
			postData.MustAppend(k, v)
		}
	}
	if postData.IsNil() || len(postData.Map()) == 0 {
		postData = gjson.New(consts.NilJsonToString)
	}

	// 当前登录用户
	if user != nil {
		memberId = user.Id
		appId = user.App
	}

	ipData, err := location.GetLocation(ctx, clientIp)
	if err != nil {
		g.Log().Errorf(ctx, "location.GetLocation err:%+v", err)
	}
	if ipData == nil {
		ipData = new(location.IpLocationData)
	}

	data = entity.SysLog{
		AppId:      appId,
		MerchantId: 0,
		MemberId:   memberId,
		Method:     request.Method,
		Module:     module,
		Url:        request.RequestURI,
		GetData:    getData,
		PostData:   postData,
		HeaderData: headerData,
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
		TakeUpTime: modelContext.TakeUpTime,
	}
	return data
}

// View 获取指定字典类型信息
func (s *sSysLog) View(ctx context.Context, in sysin.LogViewInp) (res *sysin.LogViewModel, err error) {
	if err = dao.SysLog.Ctx(ctx).Hook(hook.CityLabel).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	isDemo := g.Cfg().MustGet(ctx, "hotgo.isDemo", false)
	if isDemo.Bool() {
		res.HeaderData = gjson.New(`{
		   "none": [
		       "` + consts.DemoTips + `"
		   ]
		}`)
	}
	return
}

// Delete 删除
func (s *sSysLog) Delete(ctx context.Context, in sysin.LogDeleteInp) (err error) {
	_, err = dao.SysLog.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// List 列表
func (s *sSysLog) List(ctx context.Context, in sysin.LogListInp) (list []*sysin.LogListModel, totalCount int, err error) {
	mod := dao.SysLog.Ctx(ctx)

	// 访问路径
	if in.Url != "" {
		mod = mod.WhereLike("url", "%"+in.Url+"%")
	}

	// 模块
	if in.Module != "" {
		mod = mod.Where("module", in.Module)
	}

	// 请求方式
	if in.Method != "" {
		mod = mod.Where("method", in.Method)
	}

	// 用户
	if in.MemberId > 0 {
		mod = mod.Where("member_id", in.MemberId)
	}

	// 访问IP
	if in.Ip != "" {
		mod = mod.Where("ip", in.Ip)
	}

	// 日期范围
	if in.StartTime != "" {
		mod = mod.WhereGTE("created_at", in.StartTime)
	}
	if in.EndTime != "" {
		mod = mod.WhereLTE("created_at", in.EndTime)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("created_at", gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	// 状态码
	if in.ErrorCode != "" {
		mod = mod.Where("error_code", in.ErrorCode)
	}

	// 请求耗时
	if in.TakeUpTime > 0 {
		mod = mod.WhereGTE("take_up_time", in.TakeUpTime)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, err
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}
	isDemo := g.Cfg().MustGet(ctx, "hotgo.isDemo", false)
	for i := 0; i < len(list); i++ {
		// 管理员
		if list[i].AppId == consts.AppAdmin {
			memberName, err := dao.AdminMember.Ctx(ctx).Fields("realname").Where("id", list[i].MemberId).Value()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return list, totalCount, err
			}
			list[i].MemberName = memberName.String()
		}
		// 接口
		if list[i].AppId == consts.AppApi {
			//memberName, err = dao.Member.Ctx(ctx).Fields("realname").Where("id", res.List[i].MemberId).Value()
			//if err != nil {
			//	err = gerror.Wrap(err, consts.ErrorORM)
			//	return nil, err
			//}
		}

		if list[i].MemberName == "" {
			list[i].MemberName = "游客"
		}

		//// 获取省市编码对应的地区名称
		//region, err := dao.SysProvinces.GetRegion(ctx, list[i].ProvinceId, list[i].CityId)
		//if err != nil {
		//	return list, totalCount, err
		//}
		//list[i].Region = region

		// 截取请求url路径
		if gstr.Contains(list[i].Url, "?") {
			list[i].Url = gstr.StrTillEx(list[i].Url, "?")
		}

		if isDemo.Bool() {
			list[i].HeaderData = gjson.New(`{
			   "none": [
			       "` + consts.DemoTips + `"
			   ]
			}`)
		}

	}

	return list, totalCount, err
}
