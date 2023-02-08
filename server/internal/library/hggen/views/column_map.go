// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package views

import (
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/model/input/sysin"
)

// 字段映射关系

// go类型
const (
	GoTypeString      = "string"
	GoTypeDate        = "date"
	GoTypeDatetime    = "datetime"
	GoTypeInt         = "int"
	GoTypeUint        = "uint"
	GoTypeInt64       = "int64"
	GoTypeUint64      = "uint64"
	GoTypeIntSlice    = "[]int"
	GoTypeInt64Slice  = "[]int64"
	GoTypeUint64Slice = "[]uint64"
	GoTypeFloat32     = "float32"
	GoTypeFloat64     = "float64"
	GoTypeBytes       = "[]byte"
	GoTypeBool        = "bool"
	GoTypeTime        = "time.Time"
	GoTypeGTime       = "*gtime.Time"
	GoTypeJson        = "*gjson.Json"
)

var GoTypeNameMap = map[string]string{
	GoTypeString:      GoTypeString,
	GoTypeDate:        GoTypeDate,
	GoTypeDatetime:    GoTypeDatetime,
	GoTypeInt:         GoTypeInt,
	GoTypeUint:        GoTypeUint,
	GoTypeInt64:       GoTypeInt64,
	GoTypeUint64:      GoTypeUint64,
	GoTypeIntSlice:    GoTypeIntSlice,
	GoTypeInt64Slice:  GoTypeInt64Slice,
	GoTypeUint64Slice: GoTypeUint64Slice,
	GoTypeFloat32:     GoTypeFloat32,
	GoTypeFloat64:     GoTypeFloat64,
	GoTypeBytes:       GoTypeBytes,
	GoTypeBool:        GoTypeBool,
	GoTypeTime:        GoTypeTime,
	GoTypeGTime:       GoTypeGTime,
	GoTypeJson:        GoTypeJson,
}

// ts类型
const (
	TsTypeString  = "string"
	TsTypeNumber  = "number"
	TsTypeBoolean = "boolean"
	TsTypeArray   = "array"
	TsTypeTuple   = "tuple"
	TsTypeAny     = "any"
)

var TsTypeNameMap = map[string]string{
	TsTypeString:  TsTypeString,
	TsTypeNumber:  TsTypeNumber,
	TsTypeBoolean: TsTypeBoolean,
	TsTypeArray:   TsTypeArray,
	TsTypeTuple:   TsTypeTuple,
	TsTypeAny:     TsTypeAny,
}

// ShiftMap Go -> Ts 类型转换
var ShiftMap = map[string]string{
	GoTypeString:      TsTypeString,
	GoTypeDate:        TsTypeString,
	GoTypeDatetime:    TsTypeString,
	GoTypeInt:         TsTypeNumber,
	GoTypeUint:        TsTypeNumber,
	GoTypeInt64:       TsTypeNumber,
	GoTypeUint64:      TsTypeNumber,
	GoTypeIntSlice:    TsTypeArray,
	GoTypeInt64Slice:  TsTypeArray,
	GoTypeUint64Slice: TsTypeArray,
	GoTypeFloat32:     TsTypeNumber,
	GoTypeFloat64:     TsTypeNumber,
	GoTypeBytes:       TsTypeString,
	GoTypeBool:        TsTypeBoolean,
	GoTypeTime:        TsTypeString,
	GoTypeGTime:       TsTypeString,
	GoTypeJson:        TsTypeAny,
}

// 表单组件
const (
	FormModeInput          = "Input"          // 文本输入
	FormModeInputNumber    = "InputNumber"    // 数字输入
	FormModeInputTextarea  = "InputTextarea"  // 文本域
	FormModeInputEditor    = "InputEditor"    // 富文本
	FormModeInputDynamic   = "InputDynamic"   // 动态键值对
	FormModeDate           = "Date"           // 日期选择(Y-M-D)
	FormModeDateRange      = "DateRange"      // 日期范围选择
	FormModeTime           = "Time"           // 时间选择(Y-M-D H:i:s)
	FormModeTimeRange      = "TimeRange"      // 时间范围选择
	FormModeRadio          = "Radio"          // 单选按钮
	FormModeCheckbox       = "Checkbox"       // 复选按钮
	FormModeSelect         = "Select"         // 单选下拉框
	FormModeSelectMultiple = "SelectMultiple" // 多选下拉框
	FormModeUploadImage    = "UploadImage"    // 单图上传
	FormModeUploadImages   = "UploadImages"   // 多图上传
	FormModeUploadFile     = "UploadFile"     // 单文件上传
	FormModeUploadFiles    = "UploadFiles"    // 多文件上传
	FormModeSwitch         = "Switch"         // 开关
	FormModeRate           = "Rate"           // 评分
	FormModeCitySelector   = "CitySelector"   // 省市区选择
)

var FormModes = []string{
	FormModeInput, FormModeInputNumber, FormModeInputTextarea, FormModeInputEditor, FormModeInputDynamic,
	FormModeDate, FormModeDateRange, FormModeTime, FormModeTimeRange,
	FormModeRadio, FormModeCheckbox, FormModeSelect, FormModeSelectMultiple,
	FormModeUploadImage, FormModeUploadImages, FormModeUploadFile, FormModeUploadFiles,
	FormModeSwitch,
	FormModeRate,
	FormModeCitySelector,
}

var FormModeMap = map[string]string{
	FormModeInput:          "文本输入",
	FormModeInputNumber:    "数字输入",
	FormModeInputTextarea:  "文本域",
	FormModeInputEditor:    "富文本",
	FormModeInputDynamic:   "动态键值对",
	FormModeDate:           "日期选择(Y-M-D)",
	FormModeDateRange:      "日期范围选择",
	FormModeTime:           "时间选择(Y-M-D H:i:s)",
	FormModeTimeRange:      "时间范围选择",
	FormModeRadio:          "单选按钮",
	FormModeCheckbox:       "复选按钮",
	FormModeSelect:         "单选下拉框",
	FormModeSelectMultiple: "多选下拉框",
	FormModeUploadImage:    "单图上传",
	FormModeUploadImages:   "多图上传",
	FormModeUploadFile:     "单文件上传",
	FormModeUploadFiles:    "多文件上传",
	FormModeSwitch:         "开关",
	FormModeRate:           "评分",
	FormModeCitySelector:   "省市区选择",
}

// 表单验证
const (
	FormRoleNone       = "none"
	FormRoleIp         = "ip"
	FormRolePercentage = "percentage"
	FormRoleTel        = "tel"
	FormRolePhone      = "phone"
	FormRoleQq         = "qq"
	FormRoleEmail      = "email"
	FormRoleIdCard     = "idCard"
	FormRoleNum        = "num"
	FormRoleBankCard   = "bankCard"
	FormRoleWeibo      = "weibo"
	FormRoleUserName   = "userName"
	FormRoleAccount    = "account"
	FormRolePassword   = "password"
	FormRoleAmount     = "amount"
)

var FormRoleMap = map[string]string{
	FormRoleNone:       "不验证",
	FormRoleIp:         "Ipv4或Ipv6",
	FormRolePercentage: "0-100百分比",
	FormRoleTel:        "固话格式",
	FormRolePhone:      "手机号",
	FormRoleQq:         "QQ号码",
	FormRoleEmail:      "邮箱",
	FormRoleIdCard:     "身份证",
	FormRoleNum:        "非零正整数",
	FormRoleBankCard:   "银行卡",
	FormRoleWeibo:      "微博号",
	FormRoleUserName:   "用户名",
	FormRoleAccount:    "账号",
	FormRolePassword:   "密码",
	FormRoleAmount:     "金额",
}

// 查询条件
const (
	WhereModeEq           = "="                            // =
	WhereModeNeq          = "!="                           // !=
	WhereModeGt           = ">"                            // >
	WhereModeGte          = ">="                           // >=
	WhereModeLt           = "<"                            // <
	WhereModeLte          = "<="                           // <=
	WhereModeIn           = "IN"                           // IN (...)
	WhereModeNotIn        = "NOT IN"                       // NOT IN (...)
	WhereModeBetween      = "BETWEEN"                      // BETWEEN
	WhereModeNotBetween   = "NOT BETWEEN"                  // NOT BETWEEN
	WhereModeLike         = "LIKE"                         // LIKE
	WhereModeLikeAll      = "LIKE %...%"                   // LIKE %...%
	WhereModeNotLike      = "NOT LIKE"                     // NOT LIKE
	WhereModeJsonContains = "JSON_CONTAINS(json_doc, val)" // JSON_CONTAINS(json_doc, val[, path]) // 判断是否包含某个json值
)

var WhereModes = []string{WhereModeEq,
	WhereModeNeq, WhereModeGt, WhereModeGte, WhereModeLt, WhereModeLte,
	WhereModeIn, WhereModeNotIn,
	WhereModeBetween, WhereModeNotBetween,
	WhereModeLike, WhereModeLikeAll, WhereModeNotLike,
	WhereModeJsonContains,
}

// IsNumberType 是否是数字类型
func IsNumberType(goType string) bool {
	switch goType {
	case GoTypeInt, GoTypeUint, GoTypeInt64, GoTypeUint64:
		return true
	}
	return false
}

func HasColumn(masterFields []*sysin.GenCodesColumnListModel, column string) bool {
	for _, field := range masterFields {
		if field.GoName == column {
			return true
		}
	}
	return false
}

func HasMaxSort(masterFields []*sysin.GenCodesColumnListModel) bool {
	return HasColumn(masterFields, "Sort")
}

func HasStatus(headOps []string, masterFields []*sysin.GenCodesColumnListModel) bool {
	if !gstr.InArray(headOps, "status") {
		return false
	}
	return HasColumn(masterFields, "Status")
}

func HasSwitch(headOps []string, masterFields []*sysin.GenCodesColumnListModel) bool {
	if !gstr.InArray(headOps, "switch") {
		return false
	}
	return HasColumn(masterFields, "Switch")
}
