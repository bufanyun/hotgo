// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISysServeLicense interface {
		// Model 服务许可证ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取服务许可证列表
		List(ctx context.Context, in *sysin.ServeLicenseListInp) (list []*sysin.ServeLicenseListModel, totalCount int, err error)
		// Export 导出服务许可证
		Export(ctx context.Context, in *sysin.ServeLicenseListInp) (err error)
		// Edit 修改/新增服务许可证
		Edit(ctx context.Context, in *sysin.ServeLicenseEditInp) (err error)
		// Delete 删除服务许可证
		Delete(ctx context.Context, in *sysin.ServeLicenseDeleteInp) (err error)
		// View 获取服务许可证指定信息
		View(ctx context.Context, in *sysin.ServeLicenseViewInp) (res *sysin.ServeLicenseViewModel, err error)
		// Status 更新服务许可证状态
		Status(ctx context.Context, in *sysin.ServeLicenseStatusInp) (err error)
		// AssignRouter 分配服务许可证路由
		AssignRouter(ctx context.Context, in *sysin.ServeLicenseAssignRouterInp) (err error)
	}
	ISysSmsLog interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.SmsLogDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.SmsLogEditInp) (err error)
		// Status 更新短信状态
		Status(ctx context.Context, in *sysin.SmsLogStatusInp) (err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.SmsLogViewInp) (res *sysin.SmsLogViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.SmsLogListInp) (list []*sysin.SmsLogListModel, totalCount int, err error)
		// SendCode 发送验证码
		SendCode(ctx context.Context, in *sysin.SendCodeInp) (err error)
		// GetTemplate 获取指定短信模板
		GetTemplate(ctx context.Context, template string, config *model.SmsConfig) (val string, err error)
		// AllowSend 是否允许发送
		AllowSend(ctx context.Context, models *entity.SysSmsLog, config *model.SmsConfig) (err error)
		// VerifyCode 效验验证码
		VerifyCode(ctx context.Context, in *sysin.VerifyCodeInp) (err error)
	}
	ISysConfig interface {
		// InitConfig 初始化系统配置
		InitConfig(ctx context.Context)
		// LoadConfig 加载系统配置
		LoadConfig(ctx context.Context) (err error)
		// GetLogin 获取登录配置
		GetLogin(ctx context.Context) (conf *model.LoginConfig, err error)
		// GetWechat 获取微信配置
		GetWechat(ctx context.Context) (conf *model.WechatConfig, err error)
		// GetPay 获取支付配置
		GetPay(ctx context.Context) (conf *model.PayConfig, err error)
		// GetSms 获取短信配置
		GetSms(ctx context.Context) (conf *model.SmsConfig, err error)
		// GetGeo 获取地理配置
		GetGeo(ctx context.Context) (conf *model.GeoConfig, err error)
		// GetUpload 获取上传配置
		GetUpload(ctx context.Context) (conf *model.UploadConfig, err error)
		// GetSmtp 获取邮件配置
		GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error)
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetLoadTCP 获取本地tcp配置
		GetLoadTCP(ctx context.Context) (conf *model.TCPConfig, err error)
		// GetLoadCache 获取本地缓存配置
		GetLoadCache(ctx context.Context) (conf *model.CacheConfig, err error)
		// GetLoadGenerate 获取本地生成配置
		GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error)
		// GetLoadToken 获取本地token配置
		GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error)
		// GetLoadLog 获取本地日志配置
		GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error)
		// GetLoadServeLog 获取本地服务日志配置
		GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error)
		// GetConfigByGroup 获取指定分组的配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// ConversionType 转换类型
		ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) (err error)
		// ClusterSync 集群同步
		ClusterSync(ctx context.Context, message *gredis.Message)
	}
	ISysEmsLog interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.EmsLogDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.EmsLogEditInp) (err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *sysin.EmsLogStatusInp) (err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.EmsLogViewInp) (res *sysin.EmsLogViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.EmsLogListInp) (list []*sysin.EmsLogListModel, totalCount int, err error)
		// Send 发送邮件
		Send(ctx context.Context, in *sysin.SendEmsInp) (err error)
		// GetTemplate 获取指定邮件模板
		GetTemplate(ctx context.Context, template string, config *model.EmailConfig) (val string, err error)
		// AllowSend 是否允许发送
		AllowSend(ctx context.Context, models *entity.SysEmsLog, config *model.EmailConfig) (err error)
		// VerifyCode 效验验证码
		VerifyCode(ctx context.Context, in *sysin.VerifyEmsCodeInp) (err error)
	}
	ISysLoginLog interface {
		// Model 登录日志Orm模型
		Model(ctx context.Context) *gdb.Model
		// List 获取登录日志列表
		List(ctx context.Context, in *sysin.LoginLogListInp) (list []*sysin.LoginLogListModel, totalCount int, err error)
		// Export 导出登录日志
		Export(ctx context.Context, in *sysin.LoginLogListInp) (err error)
		// Delete 删除登录日志
		Delete(ctx context.Context, in *sysin.LoginLogDeleteInp) (err error)
		// View 获取登录日志指定信息
		View(ctx context.Context, in *sysin.LoginLogViewInp) (res *sysin.LoginLogViewModel, err error)
		// Push 推送登录日志
		Push(ctx context.Context, in *sysin.LoginLogPushInp)
		// RealWrite 真实写入
		RealWrite(ctx context.Context, models entity.SysLoginLog) (err error)
	}
	ISysAddons interface {
		// List 获取列表
		List(ctx context.Context, in *sysin.AddonsListInp) (list []*sysin.AddonsListModel, totalCount int, err error)
		// Selects 选项
		Selects(ctx context.Context, in *sysin.AddonsSelectsInp) (res *sysin.AddonsSelectsModel, err error)
		// Build 提交生成
		Build(ctx context.Context, in *sysin.AddonsBuildInp) (err error)
		// Install 安装模块
		Install(ctx context.Context, in *sysin.AddonsInstallInp) (err error)
		// Upgrade 更新模块
		Upgrade(ctx context.Context, in *sysin.AddonsUpgradeInp) (err error)
		// UnInstall 卸载模块
		UnInstall(ctx context.Context, in *sysin.AddonsUnInstallInp) (err error)
	}
	ISysAttachment interface {
		// Model ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Delete 删除附件
		Delete(ctx context.Context, in *sysin.AttachmentDeleteInp) (err error)
		// View 获取附件信息
		View(ctx context.Context, in *sysin.AttachmentViewInp) (res *sysin.AttachmentViewModel, err error)
		// List 获取附件列表
		List(ctx context.Context, in *sysin.AttachmentListInp) (list []*sysin.AttachmentListModel, totalCount int, err error)
		// ClearKind 清空上传类型
		ClearKind(ctx context.Context, in *sysin.AttachmentClearKindInp) (err error)
	}
	ISysCurdDemo interface {
		// Model 生成演示ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取生成演示列表
		List(ctx context.Context, in *sysin.CurdDemoListInp) (list []*sysin.CurdDemoListModel, totalCount int, err error)
		// Export 导出生成演示
		Export(ctx context.Context, in *sysin.CurdDemoListInp) (err error)
		// Edit 修改/新增生成演示
		Edit(ctx context.Context, in *sysin.CurdDemoEditInp) (err error)
		// Delete 删除生成演示
		Delete(ctx context.Context, in *sysin.CurdDemoDeleteInp) (err error)
		// MaxSort 获取生成演示最大排序
		MaxSort(ctx context.Context, in *sysin.CurdDemoMaxSortInp) (res *sysin.CurdDemoMaxSortModel, err error)
		// View 获取生成演示指定信息
		View(ctx context.Context, in *sysin.CurdDemoViewInp) (res *sysin.CurdDemoViewModel, err error)
		// Status 更新生成演示状态
		Status(ctx context.Context, in *sysin.CurdDemoStatusInp) (err error)
		// Switch 更新生成演示开关
		Switch(ctx context.Context, in *sysin.CurdDemoSwitchInp) (err error)
	}
	ISysGenCodes interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.GenCodesDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.GenCodesEditInp) (res *sysin.GenCodesEditModel, err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *sysin.GenCodesStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.GenCodesMaxSortInp) (res *sysin.GenCodesMaxSortModel, err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.GenCodesViewInp) (res *sysin.GenCodesViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.GenCodesListInp) (list []*sysin.GenCodesListModel, totalCount int, err error)
		// Selects 选项
		Selects(ctx context.Context, in *sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error)
		// TableSelect 表选项
		TableSelect(ctx context.Context, in *sysin.GenCodesTableSelectInp) (res []*sysin.GenCodesTableSelectModel, err error)
		// ColumnSelect 表字段选项
		ColumnSelect(ctx context.Context, in *sysin.GenCodesColumnSelectInp) (res []*sysin.GenCodesColumnSelectModel, err error)
		// ColumnList 表字段列表
		ColumnList(ctx context.Context, in *sysin.GenCodesColumnListInp) (res []*sysin.GenCodesColumnListModel, err error)
		// Preview 生成预览
		Preview(ctx context.Context, in *sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error)
		// Build 提交生成
		Build(ctx context.Context, in *sysin.GenCodesBuildInp) (err error)
	}
	ISysCron interface {
		StartCron(ctx context.Context)
		// Delete 删除
		Delete(ctx context.Context, in *sysin.CronDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.CronEditInp) (err error)
		// Status 更新状态
		Status(ctx context.Context, in *sysin.CronStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.CronMaxSortInp) (res *sysin.CronMaxSortModel, err error)
		// View 获取指定信息
		View(ctx context.Context, in *sysin.CronViewInp) (res *sysin.CronViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.CronListInp) (list []*sysin.CronListModel, totalCount int, err error)
		// OnlineExec 在线执行
		OnlineExec(ctx context.Context, in *sysin.OnlineExecInp) (err error)
	}
	ISysCronGroup interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.CronGroupDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.CronGroupEditInp) (err error)
		// Status 更新状态
		Status(ctx context.Context, in *sysin.CronGroupStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.CronGroupMaxSortInp) (res *sysin.CronGroupMaxSortModel, err error)
		// View 获取指定信息
		View(ctx context.Context, in *sysin.CronGroupViewInp) (res *sysin.CronGroupViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.CronGroupListInp) (list []*sysin.CronGroupListModel, totalCount int, err error)
		// Select 选项
		Select(ctx context.Context, in *sysin.CronGroupSelectInp) (res *sysin.CronGroupSelectModel, err error)
	}
	ISysDictData interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.DictDataDeleteInp) error
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.DictDataEditInp) (err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.DictDataListInp) (list []*sysin.DictDataListModel, totalCount int, err error)
		// Select 获取列表
		Select(ctx context.Context, in *sysin.DataSelectInp) (list sysin.DataSelectModel, err error)
	}
	ISysDictType interface {
		// Tree 树
		Tree(ctx context.Context) (list []*sysin.DictTypeTree, err error)
		// Delete 删除
		Delete(ctx context.Context, in *sysin.DictTypeDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.DictTypeEditInp) (err error)
		// TreeSelect 获取类型关系树选项
		TreeSelect(ctx context.Context, in *sysin.DictTreeSelectInp) (list []*sysin.DictTypeTree, err error)
	}
	ISysLog interface {
		// Export 导出
		Export(ctx context.Context, in *sysin.LogListInp) (err error)
		// RealWrite 真实写入
		RealWrite(ctx context.Context, log entity.SysLog) (err error)
		// AutoLog 根据配置自动记录请求日志
		AutoLog(ctx context.Context) error
		// AnalysisLog 解析日志数据
		AnalysisLog(ctx context.Context) entity.SysLog
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.LogViewInp) (res *sysin.LogViewModel, err error)
		// Delete 删除
		Delete(ctx context.Context, in *sysin.LogDeleteInp) (err error)
		// List 列表
		List(ctx context.Context, in *sysin.LogListInp) (list []*sysin.LogListModel, totalCount int, err error)
	}
	ISysProvinces interface {
		// Tree 关系树选项列表
		Tree(ctx context.Context) (list []*sysin.ProvincesTree, err error)
		// Delete 删除省市区数据
		Delete(ctx context.Context, in *sysin.ProvincesDeleteInp) (err error)
		// Edit 修改/新增省市区数据
		Edit(ctx context.Context, in *sysin.ProvincesEditInp) (err error)
		// Status 更新省市区状态
		Status(ctx context.Context, in *sysin.ProvincesStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *sysin.ProvincesMaxSortInp) (res *sysin.ProvincesMaxSortModel, err error)
		// View 获取省市区信息
		View(ctx context.Context, in *sysin.ProvincesViewInp) (res *sysin.ProvincesViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.ProvincesListInp) (list []*sysin.ProvincesListModel, totalCount int, err error)
		// ChildrenList 获取省市区下级列表
		ChildrenList(ctx context.Context, in *sysin.ProvincesChildrenListInp) (list []*sysin.ProvincesChildrenListModel, totalCount int, err error)
		// UniqueId 获取省市区下级列表
		UniqueId(ctx context.Context, in *sysin.ProvincesUniqueIdInp) (res *sysin.ProvincesUniqueIdModel, err error)
		// Select 省市区选项
		Select(ctx context.Context, in *sysin.ProvincesSelectInp) (res *sysin.ProvincesSelectModel, err error)
	}
	ISysServeLog interface {
		// Model 服务日志Orm模型
		Model(ctx context.Context) *gdb.Model
		// List 获取服务日志列表
		List(ctx context.Context, in *sysin.ServeLogListInp) (list []*sysin.ServeLogListModel, totalCount int, err error)
		// Export 导出服务日志
		Export(ctx context.Context, in *sysin.ServeLogListInp) (err error)
		// Delete 删除服务日志
		Delete(ctx context.Context, in *sysin.ServeLogDeleteInp) (err error)
		// View 获取服务日志指定信息
		View(ctx context.Context, in *sysin.ServeLogViewInp) (res *sysin.ServeLogViewModel, err error)
		// RealWrite 真实写入
		RealWrite(ctx context.Context, models entity.SysServeLog) (err error)
	}
	ISysAddonsConfig interface {
		// GetConfigByGroup 获取指定分组的配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetAddonsConfigInp) (res *sysin.GetAddonsConfigModel, err error)
		// ConversionType 转换类型
		ConversionType(ctx context.Context, models *entity.SysAddonsConfig) (value interface{}, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateAddonsConfigInp) (err error)
	}
	ISysBlacklist interface {
		// Delete 删除
		Delete(ctx context.Context, in *sysin.BlacklistDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *sysin.BlacklistEditInp) (err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *sysin.BlacklistStatusInp) (err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *sysin.BlacklistViewInp) (res *sysin.BlacklistViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *sysin.BlacklistListInp) (list []*sysin.BlacklistListModel, totalCount int, err error)
		// VariableLoad 变化加载
		VariableLoad(ctx context.Context, err error)
		// Load 加载黑名单
		Load(ctx context.Context)
		// VerifyRequest 验证请求的访问IP是否在黑名单，如果存在则返回错误
		VerifyRequest(r *ghttp.Request) (err error)
		// ClusterSync 集群同步
		ClusterSync(ctx context.Context, message *gredis.Message)
	}
)

var (
	localSysConfig       ISysConfig
	localSysEmsLog       ISysEmsLog
	localSysServeLicense ISysServeLicense
	localSysSmsLog       ISysSmsLog
	localSysAddons       ISysAddons
	localSysAttachment   ISysAttachment
	localSysLoginLog     ISysLoginLog
	localSysCron         ISysCron
	localSysCronGroup    ISysCronGroup
	localSysCurdDemo     ISysCurdDemo
	localSysGenCodes     ISysGenCodes
	localSysProvinces    ISysProvinces
	localSysServeLog     ISysServeLog
	localSysAddonsConfig ISysAddonsConfig
	localSysBlacklist    ISysBlacklist
	localSysDictData     ISysDictData
	localSysDictType     ISysDictType
	localSysLog          ISysLog
)

func SysAddons() ISysAddons {
	if localSysAddons == nil {
		panic("implement not found for interface ISysAddons, forgot register?")
	}
	return localSysAddons
}

func RegisterSysAddons(i ISysAddons) {
	localSysAddons = i
}

func SysAttachment() ISysAttachment {
	if localSysAttachment == nil {
		panic("implement not found for interface ISysAttachment, forgot register?")
	}
	return localSysAttachment
}

func RegisterSysAttachment(i ISysAttachment) {
	localSysAttachment = i
}

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}

func SysCron() ISysCron {
	if localSysCron == nil {
		panic("implement not found for interface ISysCron, forgot register?")
	}
	return localSysCron
}

func RegisterSysCron(i ISysCron) {
	localSysCron = i
}

func SysCronGroup() ISysCronGroup {
	if localSysCronGroup == nil {
		panic("implement not found for interface ISysCronGroup, forgot register?")
	}
	return localSysCronGroup
}

func RegisterSysCronGroup(i ISysCronGroup) {
	localSysCronGroup = i
}

func SysCurdDemo() ISysCurdDemo {
	if localSysCurdDemo == nil {
		panic("implement not found for interface ISysCurdDemo, forgot register?")
	}
	return localSysCurdDemo
}

func RegisterSysCurdDemo(i ISysCurdDemo) {
	localSysCurdDemo = i
}

func SysGenCodes() ISysGenCodes {
	if localSysGenCodes == nil {
		panic("implement not found for interface ISysGenCodes, forgot register?")
	}
	return localSysGenCodes
}

func RegisterSysGenCodes(i ISysGenCodes) {
	localSysGenCodes = i
}

func SysProvinces() ISysProvinces {
	if localSysProvinces == nil {
		panic("implement not found for interface ISysProvinces, forgot register?")
	}
	return localSysProvinces
}

func RegisterSysProvinces(i ISysProvinces) {
	localSysProvinces = i
}

func SysServeLog() ISysServeLog {
	if localSysServeLog == nil {
		panic("implement not found for interface ISysServeLog, forgot register?")
	}
	return localSysServeLog
}

func RegisterSysServeLog(i ISysServeLog) {
	localSysServeLog = i
}

func SysAddonsConfig() ISysAddonsConfig {
	if localSysAddonsConfig == nil {
		panic("implement not found for interface ISysAddonsConfig, forgot register?")
	}
	return localSysAddonsConfig
}

func RegisterSysAddonsConfig(i ISysAddonsConfig) {
	localSysAddonsConfig = i
}

func SysBlacklist() ISysBlacklist {
	if localSysBlacklist == nil {
		panic("implement not found for interface ISysBlacklist, forgot register?")
	}
	return localSysBlacklist
}

func RegisterSysBlacklist(i ISysBlacklist) {
	localSysBlacklist = i
}

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}

func SysLog() ISysLog {
	if localSysLog == nil {
		panic("implement not found for interface ISysLog, forgot register?")
	}
	return localSysLog
}

func RegisterSysLog(i ISysLog) {
	localSysLog = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysEmsLog() ISysEmsLog {
	if localSysEmsLog == nil {
		panic("implement not found for interface ISysEmsLog, forgot register?")
	}
	return localSysEmsLog
}

func RegisterSysEmsLog(i ISysEmsLog) {
	localSysEmsLog = i
}

func SysServeLicense() ISysServeLicense {
	if localSysServeLicense == nil {
		panic("implement not found for interface ISysServeLicense, forgot register?")
	}
	return localSysServeLicense
}

func RegisterSysServeLicense(i ISysServeLicense) {
	localSysServeLicense = i
}

func SysSmsLog() ISysSmsLog {
	if localSysSmsLog == nil {
		panic("implement not found for interface ISysSmsLog, forgot register?")
	}
	return localSysSmsLog
}

func RegisterSysSmsLog(i ISysSmsLog) {
	localSysSmsLog = i
}
