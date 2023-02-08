// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sysin

import (
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// GenCodesMaxSortInp 最大排序
type GenCodesMaxSortInp struct {
	Id int64
}

type GenCodesMaxSortModel struct {
	Sort int
}

// GenCodesEditInp 修改/新增数据
type GenCodesEditInp struct {
	entity.SysGenCodes
}
type GenCodesEditModel struct {
	entity.SysGenCodes
}

// GenCodesDeleteInp 删除
type GenCodesDeleteInp struct {
	Id interface{}
}
type GenCodesDeleteModel struct{}

// GenCodesViewInp 获取信息
type GenCodesViewInp struct {
	Id int64
}

type GenCodesViewModel struct {
	entity.SysGenCodes
}

// GenCodesListInp 获取列表
type GenCodesListInp struct {
	form.PageReq
	form.StatusReq
	GenType int    `json:"genType"`
	VarName string `json:"varName"`
}

type GenCodesListModel struct {
	entity.SysGenCodes
}

// GenCodesStatusInp 更新状态
type GenCodesStatusInp struct {
	Id     int64 `json:"id"            description:"生成ID"`
	Status int   `json:"status"        description:"生成状态"`
}
type GenCodesStatusModel struct{}

// GenCodesSelectsInp 选项
type GenCodesSelectsInp struct {
}
type GenCodesSelectsModel struct {
	GenType   GenTypeSelects `json:"genType" dc:"生成类型"`
	Db        form.Selects   `json:"db" dc:"数据库选项"`
	Status    form.Selects   `json:"status" dc:"生成状态"`
	LinkMode  form.Selects   `json:"linkMode" dc:"关联表方式"`
	BuildMeth form.Selects   `json:"buildMeth" dc:"生成方式"`
	// 字段表格选项
	FormMode  form.Selects        `json:"formMode" dc:"表单组件"`
	FormRole  form.Selects        `json:"formRole" dc:"表单验证"`
	DictMode  DictTreeSelectModel `json:"dictMode" dc:"字典类型"`
	WhereMode form.Selects        `json:"whereMode" dc:"查询条件"`
}

type GenTypeSelects []*GenTypeSelect

type GenTypeSelect struct {
	Value     int          `json:"value"`
	Label     string       `json:"label"`
	Name      string       `json:"name"`
	Templates form.Selects `json:"templates"`
}

func (p GenTypeSelects) Len() int {
	return len(p)
}
func (p GenTypeSelects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p GenTypeSelects) Less(i, j int) bool {
	return gconv.Int64(p[j].Value) > gconv.Int64(p[i].Value)
}

// GenCodesTableSelectInp 数据库表选项
type GenCodesTableSelectInp struct {
	Name string `json:"name" dc:"数据库配置名称"`
}
type GenCodesTableSelectModel struct {
	Value           string `json:"value"`
	Label           string `json:"label"`
	Name            string `json:"name"`
	DaoName         string `json:"daoName" dc:"orm模型名称"`
	DefVarName      string `json:"defVarName" dc:"默认实体名称"`
	DefAlias        string `json:"defAlias" dc:"默认关联表别名"`
	DefTableComment string `json:"defTableComment" dc:"默认菜单名称"`
}

// GenCodesColumnSelectInp 表字段选项
type GenCodesColumnSelectInp struct {
	Name  string `json:"name" dc:"数据库配置名称"`
	Table string `json:"table" dc:"表名称"`
}
type GenCodesColumnSelectModel struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Name  string `json:"name"`
}

// GenCodesColumnListInp 表字段列表
type GenCodesColumnListInp struct {
	Name   string `json:"name" dc:"数据库配置名称"`
	Table  string `json:"table" dc:"表名称"`
	IsLink int64  `json:"isLink" dc:"是否是关联表"`
	Alias  string `json:"alias" dc:"关联表别名"`
}
type GenCodesColumnListModel struct {
	model.GenCodesColumn
}

// GenCodesPreviewInp 生成预览
type GenCodesPreviewInp struct {
	entity.SysGenCodes
}

// GenFile 生成文件配置
type GenFile struct {
	Content  string `json:"content" dc:"页面内容"`
	Path     string `json:"path" dc:"生成路径"`
	Meth     int    `json:"meth" dc:"生成方式"`
	Required bool   `json:"required" dc:"是否是必要构建文件"`
}

type GenCodesPreviewModel struct {
	Config *model.GenerateConfig `json:"config"`
	Views  map[string]*GenFile   `json:"views" dc:"页面"`
}

// GenCodesBuildInp 提交生成
type GenCodesBuildInp struct {
	entity.SysGenCodes
}
