// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/internal/model/input/sysin"
	"os"
	"strings"
)

// parseServFunName 解析业务服务名称
func (l *gCurd) parseServFunName(templateGroup, varName string) string {
	templateGroup = gstr.UcFirst(templateGroup)
	if gstr.HasPrefix(varName, templateGroup) && varName != templateGroup {
		return varName
	}
	return templateGroup + varName
}

// getPkField 获取主键
func (l *gCurd) getPkField(in *CurdPreviewInput) *sysin.GenCodesColumnListModel {
	if len(in.masterFields) == 0 {
		panic("getPkField masterFields uninitialized.")
	}
	for _, field := range in.masterFields {
		if IsIndexPK(field.Index) {
			return field
		}
	}
	return nil
}

// hasEffectiveJoin 存在有效的关联表
func hasEffectiveJoins(joins []*CurdOptionsJoin) bool {
	for _, join := range joins {
		if isEffectiveJoin(join) {
			return true
		}
	}
	return false
}

func isEffectiveJoin(join *CurdOptionsJoin) bool {
	return join.Alias != "" && join.Field != "" && join.LinkTable != "" && join.MasterField != "" && join.DaoName != "" && join.LinkMode > 0
}

// formatComment formats the comment string to fit the golang code without any lines.
func formatComment(comment string) string {
	comment = gstr.ReplaceByArray(comment, g.SliceStr{
		"\n", " ",
		"\r", " ",
	})
	comment = gstr.Replace(comment, `\n`, " ")
	comment = gstr.Trim(comment)
	return comment
}

// 移除末尾的换行符
func removeEndWrap(comment string) string {
	if len(comment) > 2 && comment[len(comment)-2:] == " \n" {
		comment = comment[:len(comment)-2]
	}
	return comment
}

// ImportSql 导出sql文件
func ImportSql(ctx context.Context, path string) error {
	rows, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	sqlArr := strings.Split(string(rows), "\n")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" || strings.HasPrefix(sql, "--") {
			continue
		}
		exec, err := g.DB().Exec(ctx, sql)
		g.Log().Infof(ctx, "views.ImportSql sql:%v, exec:%+v, err:%+v", sql, exec, err)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkCurdPath(temp *model.GenerateAppCrudTemplate, addonName string) (err error) {
	if temp == nil {
		return gerror.New("生成模板配置不能为空")
	}

	if temp.IsAddon {
		temp.TemplatePath = gstr.Replace(temp.TemplatePath, "{$name}", addonName)
		temp.ApiPath = gstr.Replace(temp.ApiPath, "{$name}", addonName)
		temp.InputPath = gstr.Replace(temp.InputPath, "{$name}", addonName)
		temp.ControllerPath = gstr.Replace(temp.ControllerPath, "{$name}", addonName)
		temp.LogicPath = gstr.Replace(temp.LogicPath, "{$name}", addonName)
		temp.RouterPath = gstr.Replace(temp.RouterPath, "{$name}", addonName)
		temp.SqlPath = gstr.Replace(temp.SqlPath, "{$name}", addonName)
		temp.WebApiPath = gstr.Replace(temp.WebApiPath, "{$name}", addonName)
		temp.WebViewsPath = gstr.Replace(temp.WebViewsPath, "{$name}", addonName)
	}

	tip := `生成模板配置参数'%s'路径不存在，请先创建路径:%s`

	if !gfile.Exists(temp.TemplatePath) {
		return gerror.Newf(tip, "TemplatePath", temp.TemplatePath)
	}
	if !gfile.Exists(temp.ApiPath) {
		return gerror.Newf(tip, "ApiPath", temp.ApiPath)
	}
	if !gfile.Exists(temp.InputPath) {
		return gerror.Newf(tip, "InputPath", temp.InputPath)
	}
	if !gfile.Exists(temp.ControllerPath) {
		return gerror.Newf(tip, "ControllerPath", temp.ControllerPath)
	}
	if !gfile.Exists(temp.LogicPath) {
		return gerror.Newf(tip, "LogicPath", temp.LogicPath)
	}
	if !gfile.Exists(temp.RouterPath) {
		return gerror.Newf(tip, "RouterPath", temp.RouterPath)
	}
	if !gfile.Exists(temp.SqlPath) {
		return gerror.Newf(tip, "SqlPath", temp.SqlPath)
	}
	if !gfile.Exists(temp.WebApiPath) {
		return gerror.Newf(tip, "WebApiPath", temp.WebApiPath)
	}
	if !gfile.Exists(temp.WebViewsPath) {
		return gerror.Newf(tip, "WebViewsPath", temp.WebViewsPath)
	}
	return
}

// GetModName 获取主包名
func GetModName(ctx context.Context) (modName string, err error) {
	if !gfile.Exists("go.mod") {
		err = gerror.New("go.mod does not exist in current working directory")
		return
	}

	var (
		goModContent = gfile.GetContents("go.mod")
		match, _     = gregex.MatchString(`^module\s+(.+)\s*`, goModContent)
	)

	if len(match) > 1 {
		modName = gstr.Trim(match[1])
	} else {
		err = gerror.New("module name does not found in go.mod")
		return
	}
	return
}

// IsIndexPK 是否是主键
func IsIndexPK(index string) bool {
	return gstr.ToUpper(index) == gstr.ToUpper(consts.GenCodesIndexPK)
}

// IsIndexUNI 是否是唯一索引
func IsIndexUNI(index string) bool {
	return gstr.ToUpper(index) == gstr.ToUpper(consts.GenCodesIndexUNI)
}
