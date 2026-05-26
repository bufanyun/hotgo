// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// GenCodesMaxSortInp 最大排序
type GenCodesMaxSortInp struct {
	Id int64 `json:"id" dc:"生成代码ID"`
}

type GenCodesMaxSortModel struct {
	Sort int `json:"sort" dc:"排序"`
}

// GenCodesEditInp 修改/新增数据
type GenCodesEditInp struct {
	entity.SysGenCodes
}

func (in *GenCodesEditInp) Filter(ctx context.Context) (err error) {
	return genFilter(ctx, in.SysGenCodes)
}

type GenCodesEditModel struct {
	entity.SysGenCodes
}

// GenCodesDeleteInp 删除
type GenCodesDeleteInp struct {
	Id interface{} `json:"id" v:"required#生成代码ID不能为空" dc:"生成代码ID"`
}

type GenCodesDeleteModel struct{}

// GenCodesViewInp 获取信息
type GenCodesViewInp struct {
	Id int64 `json:"id" v:"required#生成代码ID不能为空" dc:"生成代码ID"`
}

type GenCodesViewModel struct {
	entity.SysGenCodes
}

// GenCodesListInp 获取列表
type GenCodesListInp struct {
	form.PageReq
	form.StatusReq
	GenType int    `json:"genType" dc:"生成类型"`
	VarName string `json:"varName" dc:"实体"`
}

type GenCodesListModel struct {
	entity.SysGenCodes
	GenTemplateGroup string `json:"genTemplateGroup" dc:"生成模板组名"`
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
	FormMode      form.Selects    `json:"formMode"  dc:"表单组件"`
	FormRole      form.Selects    `json:"formRole"  dc:"表单验证"`
	DictMode      []*DictTypeTree `json:"dictMode"  dc:"字典类型"`
	WhereMode     form.Selects    `json:"whereMode" dc:"查询条件"`
	Addons        form.Selects    `json:"addons"    dc:"插件选项"`
	TableAlign    form.Selects    `json:"tableAlign"    dc:"表格排列方式"`
	TreeStyleType []*model.Option `json:"treeStyleType" dc:"树表样式选项"`
}

type GenTypeSelects []*GenTypeSelect

type GenTypeSelect struct {
	Value     int                `json:"value"`
	Label     string             `json:"label"`
	Name      string             `json:"name"`
	Templates GenTemplateSelects `json:"templates"`
}

type GenTemplateSelects []*GenTemplateSelect

type GenTemplateSelect struct {
	Value   interface{} `json:"value"`
	Label   string      `json:"label"`
	Name    string      `json:"name"`
	IsAddon bool        `json:"isAddon"`
}

func (p GenTemplateSelects) Len() int {
	return len(p)
}

func (p GenTemplateSelects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p GenTemplateSelects) Less(i, j int) bool {
	return gconv.Int64(p[j].Value) > gconv.Int64(p[i].Value)
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

func (in *GenCodesPreviewInp) Filter(ctx context.Context) (err error) {
	return genFilter(ctx, in.SysGenCodes)
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

func (in *GenCodesBuildInp) Filter(ctx context.Context) (err error) {
	return genFilter(ctx, in.SysGenCodes)
}

func genFilter(ctx context.Context, in entity.SysGenCodes) (err error) {
	if in.VarName == "" {
		err = gerror.New("实体命名不能为空")
		return
	}

	if !gregex.IsMatchString(`^[a-zA-Z]{1}\w{1,23}$`, in.VarName) {
		err = gerror.New("实体命名格式不正确，字母开头，只能包含字母、数字和下划线，长度在2~24之间")
		return
	}

	if in.GenType == consts.GenCodesTypeCurd || in.GenType == consts.GenCodesTypeTree {
		if in.DbName == "" {
			err = gerror.New("数据库不能为空")
			return
		}
		if in.TableName == "" {
			err = gerror.New("数据库表不能为空")
			return
		}
		if in.TableComment == "" {
			err = gerror.New("菜单名称不能为空")
			return
		}
	}
	return
}
