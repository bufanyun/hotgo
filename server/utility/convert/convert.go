// Package convert
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package convert

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/utility/validate"
	"reflect"
	"strings"
	"unicode"
)

var (
	descTags  = []string{"description", "dc", "json"} // 实体描述标签
	fieldTags = []string{"json"}                      // 实体字段名称映射
)

// UniqueSliceInt64 切片去重
func UniqueSliceInt64(languages []int64) []int64 {
	result := make([]int64, 0, len(languages))
	temp := map[int64]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// UniqueSliceString 切片去重
func UniqueSliceString(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// UnderlineToUpperCamelCase 下划线单词转为大写驼峰单词
func UnderlineToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderlineToLowerCamelCase 下划线单词转为小写驼峰单词
func UnderlineToLowerCamelCase(s string) string {
	s = UnderlineToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//CamelCaseToUnderline 驼峰单词转下划线单词
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
	if validate.InSliceString(filterTags, "description") {
		if description, ok := field.Tag.Lookup("description"); ok && description != "" {
			return description
		}
	}

	if validate.InSliceString(filterTags, "dc") {
		if dc, ok := field.Tag.Lookup("dc"); ok && dc != "" {
			return dc
		}
	}

	if validate.InSliceString(filterTags, "json") {
		if jsonName, ok := field.Tag.Lookup("json"); ok && jsonName != "" {
			return jsonName
		}
	}

	if !isDef {
		return ""
	}
	return field.Name
}
