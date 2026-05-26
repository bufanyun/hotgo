// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/sysin"
)

// 默认表单组件映射 Ts -> 表单组件
var defaultFormModeMap = map[string]string{
	TsTypeString:  FormModeInput,
	TsTypeNumber:  FormModeInputNumber,
	TsTypeBoolean: FormModeInputNumber,
	TsTypeArray:   FormModeInputDynamic,
	TsTypeTuple:   FormModeInputDynamic,
	TsTypeAny:     FormModeInput,
}

var defaultEditFields = map[string]bool{
	"id":         false,
	"created_at": false,
	"updated_at": false,
	"deleted_at": false,
}

var defaultEditSwitch = map[string]bool{
	"id":         false,
	"level":      false,
	"tree":       false,
	"created_by": false,
	"updated_by": false,
	"deleted_by": false,
	"created_at": false,
	"updated_at": false,
	"deleted_at": false,
}

var defaultListSwitch = map[string]bool{
	"level":      false,
	"tree":       false,
	"deleted_at": false,
}

var defaultExportSwitch = map[string]bool{
	"level":      false,
	"tree":       false,
	"deleted_at": false,
}

var defaultQuerySwitch = map[string]bool{
	"level":      false,
	"tree":       false,
	"deleted_at": false,
}

var defaultSort = map[string]bool{
	"id":   true,
	"sort": true,
}

var defaultTreeFields = []string{"pid", "level", "tree"}

// 默认表单验证映射 物理类型命名识别
var defaultFormRoleMap = map[string]string{
	"mobile":    FormRolePhone,
	"qq":        FormRoleQq,
	"email":     FormRoleEmail,
	"id_card":   FormRoleIdCard,
	"bank_card": FormRoleBankCard,
	"password":  FormRolePassword,
	"pass":      FormRolePassword,
	"price":     FormRoleAmount,
}

// 默认查询条件映射 go类型识别
var defaultWhereModeMap = map[string]string{
	GoTypeString:      WhereModeLike,
	GoTypeDate:        WhereModeEq,
	GoTypeDatetime:    WhereModeEq,
	GoTypeInt:         WhereModeEq,
	GoTypeUint:        WhereModeEq,
	GoTypeInt64:       WhereModeEq,
	GoTypeUint64:      WhereModeEq,
	GoTypeIntSlice:    WhereModeIn,
	GoTypeInt64Slice:  WhereModeIn,
	GoTypeUint64Slice: WhereModeIn,
	GoTypeFloat32:     WhereModeEq,
	GoTypeFloat64:     WhereModeEq,
	GoTypeBytes:       WhereModeEq,
	GoTypeTime:        WhereModeEq,
	GoTypeGTime:       WhereModeEq,
	GoTypeJson:        WhereModeJsonContains,
}

// setDefault 设置默认属性
func setDefault(field *sysin.GenCodesColumnListModel) {
	setDefaultEdit(field)

	setDefaultFormMode(field)

	setDefaultFormRole(field)

	setDefaultDictType(field)

	setDefaultList(field)

	setDefaultExport(field)

	setDefaultQuery(field)

	setDefaultQueryWhere(field)

	setDefaultValue(field)

	if field.IsAllowNull == "NO" {
		field.Required = true
	}

	if IsIndexUNI(field.Index) {
		field.Unique = true
	}

	if df, ok := defaultSort[field.Name]; ok {
		field.IsSort = df
	}

	if field.Dc == "" {
		field.Dc = field.Name
	}
}

// setDefaultEdit 设置默认编辑
func setDefaultEdit(field *sysin.GenCodesColumnListModel) {
	field.IsEdit = true

	if IsIndexPK(field.Index) {
		field.IsEdit = false
		return
	}

	if df, ok := defaultEditFields[field.Name]; ok {
		field.IsEdit = df
	}

	if df, ok := defaultEditSwitch[field.Name]; ok {
		field.IsEdit = df
	}
}

// setDefaultFormMode 设置默认表单组件
func setDefaultFormMode(field *sysin.GenCodesColumnListModel) {
	field.FormMode = FormModeInput
	if df, ok := defaultFormModeMap[field.TsType]; ok {
		field.FormMode = df
	}

	if gstr.HasSuffix(field.GoName, "Status") && IsNumberType(field.GoType) {
		field.FormMode = FormModeSelect
		return
	}

	if field.GoName == "CreatedAt" {
		field.FormMode = FormModeTimeRange
		return
	}

	if (field.GoName == "ProvinceId" || field.GoName == "CityId") && IsNumberType(field.GoType) {
		field.FormMode = FormModeCitySelector
		return
	}

	if field.DataType == "datetime" || field.DataType == "timestamp" || field.DataType == "timestamptz" {
		field.FormMode = FormModeTime
		return
	}

	if field.DataType == "date" {
		field.FormMode = FormModeDate
		return
	}

	if field.GoType == GoTypeString && field.Length >= 256 && field.Length <= 512 {
		field.FormMode = FormModeInputTextarea
		return
	}

	if field.GoType == GoTypeString && field.Length > 512 {
		field.FormMode = FormModeInputEditor
		return
	}
}

// setDefaultFormRole 设置默认表单验证
func setDefaultFormRole(field *sysin.GenCodesColumnListModel) {
	field.FormRole = FormRoleNone

	switch field.GoType {
	case GoTypeUint, GoTypeUint64:
		field.FormRole = FormRoleNum
		return
	}

	if df, ok := defaultFormRoleMap[field.Name]; ok {
		field.FormRole = df
	}
}

// setDefaultDictType 设置默认字典类型
func setDefaultDictType(field *sysin.GenCodesColumnListModel) {
	if gstr.HasSuffix(field.GoName, "Status") && IsNumberType(field.GoType) {
		field.DictType = 3 // 默认系统状态ID
		return
	}
}

// setDefaultList 设置默认列表
func setDefaultList(field *sysin.GenCodesColumnListModel) {
	field.IsList = true
	switch field.GoType {
	case GoTypeIntSlice, GoTypeInt64Slice, GoTypeUint64Slice, GoTypeBytes, GoTypeJson:
		field.IsList = false
		return
	}

	if field.Length >= 500 {
		field.IsList = false
		return
	}

	if df, ok := defaultListSwitch[field.Name]; ok {
		field.IsList = df
	}
}

// setDefaultExport 设置默认导出
func setDefaultExport(field *sysin.GenCodesColumnListModel) {
	field.IsExport = true
	switch field.GoType {
	case GoTypeIntSlice, GoTypeInt64Slice, GoTypeUint64Slice, GoTypeBytes, GoTypeJson:
		field.IsExport = false
		return
	}

	if field.Length >= 500 {
		field.IsExport = false
		return
	}

	if df, ok := defaultExportSwitch[field.Name]; ok {
		field.IsExport = df
	}
}

// setDefaultQuery 设置默认查询
func setDefaultQuery(field *sysin.GenCodesColumnListModel) {
	field.IsQuery = false

	if IsIndexPK(field.Index) {
		field.IsQuery = true
		return
	}

	if gstr.HasSuffix(field.GoName, "Status") && IsNumberType(field.GoType) {
		field.IsQuery = true
		return
	}

	if field.GoName == "CreatedAt" {
		field.IsQuery = true
		return
	}

	if df, ok := defaultQuerySwitch[field.Name]; ok {
		field.IsQuery = df
	}
}

// setDefaultQueryWhere 设置默认查询条件
func setDefaultQueryWhere(field *sysin.GenCodesColumnListModel) {
	field.QueryWhere = WhereModeEq

	if field.GoName == "CreatedAt" {
		field.QueryWhere = WhereModeBetween
		return
	}

	if field.Length >= 500 {
		field.QueryWhere = WhereModeLikeAll
		return
	}

	if df, ok := defaultWhereModeMap[field.GoType]; ok {
		field.QueryWhere = df
	}
}

// setDefaultValue 设置默认value
func setDefaultValue(field *sysin.GenCodesColumnListModel) {
	var value interface{}
	if field.DefaultValue == nil {
		switch field.GoType {
		case GoTypeString, GoTypeBytes, GoTypeDate, GoTypeDatetime, GoTypeTime, GoTypeGTime:
			value = ""
		case GoTypeIntSlice, GoTypeInt64Slice, GoTypeUint64Slice, GoTypeJson:
			value = nil
		case GoTypeInt, GoTypeUint, GoTypeInt64, GoTypeUint64:
			value = 0
		case GoTypeBool:
			value = false
		}
	} else {
		value = consts.ConvType(field.DefaultValue, field.GoType)
	}

	// 时间类型不做默认值处理
	if field.GoType == GoTypeGTime {
		value = ""
	}

	field.DefaultValue = value
}
