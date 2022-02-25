//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysService

import (
	"context"
	"encoding/json"
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/factory/queue"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/bufanyun/hotgo/app/utils"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

var Log = new(log)

type log struct{}

//
//  @Title  导出
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   in
//  @Return  err
//
func (service *log) Export(ctx context.Context, in input.LogListInp) (err error) {

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
		MemberName string      `json:"member_name"`
		Region     string      `json:"region"`
	}

	var (
		titleList  = []string{"ID", "应用", "提交类型", "模块", "提交url", "ip地址", "报错code", "报错信息", "对外id", "请求耗时", "创建时间", "用户", "访问地"}
		fileName   = "全局日志导出-" + com.Context.Get(ctx).ReqId + ".xlsx"
		sheetName  = "HotGo"
		exportList []exportImage
		row        exportImage
	)

	list, _, err := service.List(ctx, in)
	if err != nil {
		return err
	}

	// TODO  格式化格式
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

	// TODO  强转类型
	writer := com.Context.Get(ctx).Request.Response.Writer
	w, _ := interface{}(writer).(*ghttp.ResponseWriter)

	if err = utils.Excel.ExportByStruct(w, titleList, gconv.Interfaces(exportList), fileName, sheetName); err != nil {
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

	return
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
func (service *log) List(ctx context.Context, in input.LogListInp) (list []*input.LogListModel, totalCount int, err error) {

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

	err = mod.Page(in.Page, in.Limit).Order("id desc").Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for i := 0; i < len(list); i++ {
		// TODO  管理员
		if list[i].AppId == consts.AppAdmin {
			memberName, err := dao.AdminMember.Ctx(ctx).Fields("realname").Where("id", list[i].MemberId).Value()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return list, totalCount, err
			}
			list[i].MemberName = memberName.String()
		}
		// TODO  接口
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

		// TODO  获取省市编码对应的地区名称
		region, err := dao.SysProvinces.GetRegion(ctx, list[i].ProvinceId, list[i].CityId)
		if err != nil {
			return list, totalCount, err
		}
		list[i].Region = region

		// TODO  截取请求url路径
		if gstr.Contains(list[i].Url, "?") {
			list[i].Url = gstr.StrTillEx(list[i].Url, "?")
		}
	}

	return list, totalCount, err
}

//
//  @Title  真实写入
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   commonLog
//  @Return  err
//
func (service *log) RealWrite(ctx context.Context, commonLog entity.SysLog) error {

	result, err := dao.SysLog.Ctx(ctx).Data(commonLog).Insert()
	if err != nil {
		return err
	}

	if _, err := result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

//
//  @Title  根据配置自动记录请求日志
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Return  err
//
func (service *log) AutoLog(ctx context.Context) (err error) {
	// TODO  日志开关
	logSwitch, _ := g.Cfg().Get(ctx, "hotgo.log.switch", true)
	if !logSwitch.Bool() {
		return nil
	}

	data := service.AnalysisLog(ctx)

	// TODO  判断模块是否需要记录
	module, _ := g.Cfg().Get(ctx, "hotgo.log.module", nil)
	if module == nil {
		return nil
	}
	if exist := utils.Charset.IsExists(module.Strings(), data.Module); !exist {
		return nil
	}

	// TODO  判断状态码是否需要记录
	code, _ := g.Cfg().Get(ctx, "hotgo.log.skipCode", nil)
	if code != nil {
		if exist := utils.Charset.IsExists(code.Strings(), gconv.String(data.ErrorCode)); exist {
			return nil
		}
	}

	// TODO  是否开启队列
	queueSwitch, _ := g.Cfg().Get(ctx, "hotgo.log.queue", true)
	if queueSwitch.Bool() {
		// TODO  获取生产者实例
		queueInstance, err := queue.InstanceProducer()
		if err != nil {
			queue.FatalLog(ctx, "InstanceProducer异常", err)
			return err
		}

		// TODO  生产消息
		mqMsg, err := queueInstance.SendMsg(consts.QueueLogTopic, gconv.String(data))

		// TODO  记录生产日志
		queue.ProducerLog(ctx, consts.QueueLogTopic, mqMsg.MsgId, err)

		return err
	}

	return service.RealWrite(ctx, data)
}

//
//  @Title  队列消费
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   mqMsg
//  @Return  err
//
func (service *log) QueueJob(ctx context.Context, mqMsg queue.MqMsg) (err error) {

	var data entity.SysLog
	if err = json.Unmarshal(mqMsg.Body, &data); err != nil {
		return err
	}

	return service.RealWrite(ctx, data)
}

//
//  @Title  解析日志数据
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Return  entity.SysLog
//
func (service *log) AnalysisLog(ctx context.Context) entity.SysLog {

	var (
		modelContext = com.Context.Get(ctx)
		response     = modelContext.ComResponse
		user         = modelContext.User
		request      = modelContext.Request
		module       = modelContext.Module
		ip           = request.GetClientIp()
		locationData = com.Ip.GetLocation(ctx, ip)
		postData     = "null"
		getData      = "null"
		headerData   = "null"
		data         = entity.SysLog{}
		memberId     = 0
		errorCode    = 0
		errorMsg     = ""
		errorData    = "null"
		reqId        = ""
		timestamp    = 0
		appId        = ""
	)

	// TODO  响应数据
	if response != nil {
		errorCode = response.Code
		errorMsg = response.Message
		reqId = response.ReqId
		timestamp = gconv.Int(response.Timestamp)

		if len(gconv.String(response.Error)) > 0 {
			errorData = gconv.String(response.Error)
		}
	}

	// TODO  请求头
	if reqHeadersBytes, _ := json.Marshal(request.Header); len(gconv.String(reqHeadersBytes)) > 0 {
		headerData = gconv.String(reqHeadersBytes)
	}

	// TODO  post参数
	if gconv.String(request.PostForm) != "" {
		postData = gconv.String(request.PostForm)
	}

	// TODO  get参数
	if len(request.URL.Query()) > 0 {
		getData = gconv.String(request.URL.Query())
	}

	// TODO  当前登录用户
	if user != nil {
		memberId = int(user.Id)
		appId = user.App
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
		Ip:         ip,
		ProvinceId: locationData.ProvinceCode,
		CityId:     locationData.CityCode,
		ErrorCode:  errorCode,
		ErrorMsg:   errorMsg,
		ErrorData:  errorData,
		ReqId:      reqId,
		Timestamp:  timestamp,
		UserAgent:  request.Header.Get("User-Agent"),
		Status:     consts.StatusEnabled,
		TakeUpTime: modelContext.TakeUpTime,
	}

	return data
}
