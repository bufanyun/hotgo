// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

// Option 字典数据选项
type Option struct {
	Key       interface{} `json:"key"`
	Label     string      `json:"label"     description:"字典标签"`
	Value     interface{} `json:"value"     description:"字典键值"`
	ValueType string      `json:"valueType" description:"键值数据类型"`
	Type      string      `json:"type"      description:"字典类型"`
	ListClass string      `json:"listClass" description:"表格回显样式"`
}
