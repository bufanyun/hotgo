// Package convert
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package convert

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/utility/validate"
	"reflect"
	"unicode"
	"unsafe"
)

var (
	descTags  = []string{"description", "dc", "json"} // 实体描述标签
	fieldTags = []string{"json"}                      // 实体字段名称映射
)

// GetModelTable 获取模型中的表定义
func GetModelTable(m *gdb.Model) (tablesInit, tables string) {
	if m == nil {
		return "", ""
	}

	v := reflect.ValueOf(m).Elem()
	t := reflect.TypeOf(m).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)

		if field.Name == "tablesInit" {
			tablesInit = reflect.NewAt(val.Type(), unsafe.Pointer(val.UnsafeAddr())).Elem().String()
			tablesInit = gstr.Replace(tablesInit, "`", "")
			continue
		}

		if field.Name == "tables" {
			tables = reflect.NewAt(val.Type(), unsafe.Pointer(val.UnsafeAddr())).Elem().String()
			continue
		}
	}
	return
}

// GetMapKeys 获取map的所有key
func GetMapKeys[K comparable](m map[K]any) []K {
	j := 0
	keys := make([]K, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

// CamelCaseToUnderline 驼峰单词转下划线单词
func CamelCaseToUnderline(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

// GetEntityFieldTags 获取实体中的字段名称
func GetEntityFieldTags(entity interface{}) (tags []string, err error) {
	var formRef = reflect.TypeOf(entity)
	for i := 0; i < formRef.NumField(); i++ {
		field := formRef.Field(i)
		if field.Type.Kind() == reflect.Struct {
			addTags, err := reflectTag(field.Type, fieldTags, []string{})
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, reflectTagName(field, fieldTags, true))
	}
	return
}

// GetEntityDescTags 获取实体中的描述标签
func GetEntityDescTags(entity interface{}) (tags []string, err error) {
	var formRef = reflect.TypeOf(entity)
	for i := 0; i < formRef.NumField(); i++ {
		field := formRef.Field(i)
		if field.Type.Kind() == reflect.Struct {
			addTags, err := reflectTag(field.Type, descTags, []string{})
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, reflectTagName(field, descTags, true))
	}
	return
}

// reflectTag 层级递增解析tag
func reflectTag(reflectType reflect.Type, filterTags []string, tags []string) ([]string, error) {
	if reflectType.Kind() == reflect.Ptr {
		return nil, gerror.Newf("reflect type do not support reflect.Ptr yet, reflectType:%+v", reflectType)
	}
	if reflectType.Kind() != reflect.Struct {
		return nil, nil
	}
	for i := 0; i < reflectType.NumField(); i++ {
		tag := reflectTagName(reflectType.Field(i), filterTags, false)
		if tag == "" {
			addTags, err := reflectTag(reflectType.Field(i).Type, filterTags, tags)
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// reflectTagName 解析实体中的描述标签，优先级：description > dc > json > Name
func reflectTagName(field reflect.StructField, filterTags []string, isDef bool) string {
	if validate.InSlice(filterTags, "description") {
		if description, ok := field.Tag.Lookup("description"); ok && description != "" {
			return description
		}
	}

	if validate.InSlice(filterTags, "dc") {
		if dc, ok := field.Tag.Lookup("dc"); ok && dc != "" {
			return dc
		}
	}

	if validate.InSlice(filterTags, "json") {
		if jsonName, ok := field.Tag.Lookup("json"); ok && jsonName != "" {
			return jsonName
		}
	}

	if !isDef {
		return ""
	}
	return field.Name
}
