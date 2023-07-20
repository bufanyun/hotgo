// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/library/hggen/internal/cmd/gendao"
	"hotgo/internal/model/input/sysin"
	"strings"
)

// DoTableColumns 获取指定表生成字段列表
func DoTableColumns(ctx context.Context, in *sysin.GenCodesColumnListInp, config gendao.CGenDaoInput) (fields []*sysin.GenCodesColumnListModel, err error) {
	var (
		sql  = "select ORDINAL_POSITION as `id`, COLUMN_NAME as `name`, COLUMN_COMMENT as `dc`, DATA_TYPE as `dataType`, COLUMN_TYPE as `sqlType`, CHARACTER_MAXIMUM_LENGTH as `length`, IS_NULLABLE as `isAllowNull`, COLUMN_DEFAULT as `defaultValue`, COLUMN_KEY as `index`, EXTRA as `extra` from information_schema.COLUMNS where TABLE_SCHEMA = '%s' and TABLE_NAME = '%s' ORDER BY `id` ASC"
		conf = g.DB(in.Name).GetConfig()
	)

	err = g.DB(in.Name).Ctx(ctx).Raw(fmt.Sprintf(sql, conf.Name, in.Table)).Scan(&fields)
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		if in.IsLink == 1 {
			CustomLinkAttributes(ctx, in.Alias, field, config)
		} else {
			CustomAttributes(ctx, field, config)
		}
	}
	return
}

// CustomLinkAttributes 可自定义关联表的字段属性
func CustomLinkAttributes(ctx context.Context, alias string, field *sysin.GenCodesColumnListModel, in gendao.CGenDaoInput) {
	field.GoName, field.GoType, field.TsName, field.TsType = GenGotype(ctx, field, in)

	field.GoName = gstr.UcFirst(alias + field.GoName)
	field.TsName = gstr.LcFirst(field.GoName)

	setDefaultQueryWhere(field)
	setDefaultValue(field)
}

// CustomAttributes 可自定义的字段属性
func CustomAttributes(ctx context.Context, field *sysin.GenCodesColumnListModel, in gendao.CGenDaoInput) {
	field.GoName, field.GoType, field.TsName, field.TsType = GenGotype(ctx, field, in)
	setDefault(field)
}

// GenGotype 生成字段的go类型
func GenGotype(ctx context.Context, field *sysin.GenCodesColumnListModel, in gendao.CGenDaoInput) (goName, typeName, tsName, tsType string) {
	var err error
	tsName = getJsonTagFromCase(field.Name, in.JsonCase)
	goName = gstr.CaseCamel(field.Name)

	typeName, err = CheckLocalTypeForField(ctx, field.DataType, nil)
	if err != nil {
		panic(err)
	}
	switch typeName {
	case gdb.LocalTypeDate, gdb.LocalTypeDatetime:
		if in.StdTime {
			typeName = "time.Time"
		} else {
			typeName = "*gtime.Time"
		}

	case gdb.LocalTypeInt64Bytes:
		typeName = "int64"

	case gdb.LocalTypeUint64Bytes:
		typeName = "uint64"

	// Special type handle.
	case gdb.LocalTypeJson, gdb.LocalTypeJsonb:
		if in.GJsonSupport {
			typeName = "*gjson.Json"
		} else {
			typeName = "string"
		}
	}

	tsType = ShiftMap[typeName]
	return
}

// CheckLocalTypeForField checks and returns corresponding type for given db type.
func CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (string, error) {
	var (
		typeName    string
		typePattern string
	)
	match, _ := gregex.MatchString(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = gstr.Trim(match[1])
		typePattern = gstr.Trim(match[2])
	} else {
		typeName = gstr.Split(fieldType, " ")[0]
	}
	typeName = strings.ToLower(typeName)
	switch typeName {
	case
		"binary",
		"varbinary",
		"blob",
		"tinyblob",
		"mediumblob",
		"longblob":
		return gdb.LocalTypeBytes, nil

	case
		"int",
		"tinyint",
		"small_int",
		"smallint",
		"medium_int",
		"mediumint",
		"serial":
		if gstr.ContainsI(fieldType, "unsigned") {
			return gdb.LocalTypeUint, nil
		}
		return gdb.LocalTypeInt, nil

	case
		"big_int",
		"bigint",
		"bigserial":
		if gstr.ContainsI(fieldType, "unsigned") {
			return gdb.LocalTypeUint64, nil
		}
		return gdb.LocalTypeInt64, nil

	case
		"real":
		return gdb.LocalTypeFloat32, nil

	case
		"float",
		"double",
		"decimal",
		"money",
		"numeric",
		"smallmoney":
		return gdb.LocalTypeFloat64, nil

	case
		"bit":
		// It is suggested using bit(1) as boolean.
		if typePattern == "1" {
			return gdb.LocalTypeBool, nil
		}
		s := gconv.String(fieldValue)
		// mssql is true|false string.
		if strings.EqualFold(s, "true") || strings.EqualFold(s, "false") {
			return gdb.LocalTypeBool, nil
		}
		if gstr.ContainsI(fieldType, "unsigned") {
			return gdb.LocalTypeUint64Bytes, nil
		}
		return gdb.LocalTypeInt64Bytes, nil

	case
		"bool":
		return gdb.LocalTypeBool, nil

	case
		"date":
		return gdb.LocalTypeDate, nil

	case
		"datetime",
		"timestamp",
		"timestamptz":
		return gdb.LocalTypeDatetime, nil

	case
		"json":
		return gdb.LocalTypeJson, nil

	case
		"jsonb":
		return gdb.LocalTypeJsonb, nil

	default:
		// Auto-detect field type, using key match.
		switch {
		case strings.Contains(typeName, "text") || strings.Contains(typeName, "char") || strings.Contains(typeName, "character"):
			return gdb.LocalTypeString, nil

		case strings.Contains(typeName, "float") || strings.Contains(typeName, "double") || strings.Contains(typeName, "numeric"):
			return gdb.LocalTypeFloat64, nil

		case strings.Contains(typeName, "bool"):
			return gdb.LocalTypeBool, nil

		case strings.Contains(typeName, "binary") || strings.Contains(typeName, "blob"):
			return gdb.LocalTypeBytes, nil

		case strings.Contains(typeName, "int"):
			if gstr.ContainsI(fieldType, "unsigned") {
				return gdb.LocalTypeUint, nil
			}
			return gdb.LocalTypeInt, nil

		case strings.Contains(typeName, "time"):
			return gdb.LocalTypeDatetime, nil

		case strings.Contains(typeName, "date"):
			return gdb.LocalTypeDatetime, nil

		default:
			return gdb.LocalTypeString, nil
		}
	}
}

// getJsonTagFromCase call gstr.Case* function to convert the s to specified case.
func getJsonTagFromCase(str, caseStr string) string {
	switch gstr.ToLower(caseStr) {
	case gstr.ToLower("Camel"):
		return gstr.CaseCamel(str)

	case gstr.ToLower("CamelLower"):
		return gstr.CaseCamelLower(str)

	case gstr.ToLower("Kebab"):
		return gstr.CaseKebab(str)

	case gstr.ToLower("KebabScreaming"):
		return gstr.CaseKebabScreaming(str)

	case gstr.ToLower("Snake"):
		return gstr.CaseSnake(str)

	case gstr.ToLower("SnakeFirstUpper"):
		return gstr.CaseSnakeFirstUpper(str)

	case gstr.ToLower("SnakeScreaming"):
		return gstr.CaseSnakeScreaming(str)
	}
	return str
}
