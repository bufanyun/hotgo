// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/library/hggen/views/gohtml"
	"hotgo/internal/model"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/convert"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/gogf/gf/v2/util/gutil"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/tools/imports"
)

// parseServFunName 解析业务服务名称
func (l *gCurd) parseServFunName(templateGroup, varName string) string {
	templateGroup = gstr.UcFirst(templateGroup)
	if gstr.HasPrefix(varName, templateGroup) && varName != templateGroup {
		return varName
	}
	return templateGroup + varName
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
	sqlContent := string(rows)
	config := g.DB("default").GetConfig()
	if config.Type == consts.DBPgsql {
		return importSqlPgsql(ctx, sqlContent)
	}
	return importSqlMysql(ctx, sqlContent)
}

// importSqlMysql 导出mysql文件
func importSqlMysql(ctx context.Context, sqlContent string) error {
	sqlArr := strings.Split(sqlContent, "\n")
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

// importSqlPgsql 导出pgsql文件
func importSqlPgsql(ctx context.Context, sqlContent string) error {
	lines := strings.Split(sqlContent, "\n")
	var currentStmt strings.Builder
	inDoBlock := false
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "--") {
			continue
		}

		currentStmt.WriteString(line)
		currentStmt.WriteString("\n")

		if strings.HasPrefix(trimmedLine, "DO $$") || strings.HasPrefix(trimmedLine, "DO $") {
			inDoBlock = true
			continue
		}

		if inDoBlock && (strings.HasPrefix(trimmedLine, "END $$;") || strings.HasPrefix(trimmedLine, "END $")) {
			inDoBlock = false
			stmt := currentStmt.String()
			exec, err := g.DB().Exec(ctx, stmt)
			g.Log().Infof(ctx, "importSqlPgsql DO block executed, exec:%+v, err:%+v", exec, err)
			if err != nil {
				g.Log().Errorf(ctx, "importSqlPgsql error: %v, sql: %s", err, stmt)
				return err
			}
			currentStmt.Reset()
			continue
		}

		if !inDoBlock && strings.HasSuffix(trimmedLine, ";") {
			stmt := currentStmt.String()
			exec, err := g.DB().Exec(ctx, stmt)
			g.Log().Infof(ctx, "importSqlPgsql sql executed, exec:%+v, err:%+v", exec, err)
			if err != nil {
				g.Log().Errorf(ctx, "importSqlPgsql error: %v, sql: %s", err, stmt)
				return err
			}
			currentStmt.Reset()
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

// ParseDBConfigNodeLink 解析数据库连接配置
func ParseDBConfigNodeLink(node *gdb.ConfigNode) *gdb.ConfigNode {
	const linkPattern = `(\w+):([\w\-\$]*):(.*?)@(\w+?)\((.+?)\)/{0,1}([^\?]*)\?{0,1}(.*)`
	const defaultCharset = `utf8`
	const defaultProtocol = `tcp`

	var match []string
	if node.Link != "" {
		match, _ = gregex.MatchString(linkPattern, node.Link)
		if len(match) > 5 {
			node.Type = match[1]
			node.User = match[2]
			node.Pass = match[3]
			node.Protocol = match[4]
			array := gstr.Split(match[5], ":")
			if len(array) == 2 && node.Protocol != "file" {
				node.Host = array[0]
				node.Port = array[1]
				node.Name = match[6]
			} else {
				node.Name = match[5]
			}
			if len(match) > 6 && match[7] != "" {
				node.Extra = match[7]
			}
			node.Link = ""
		}
	}
	if node.Extra != "" {
		if m, _ := gstr.Parse(node.Extra); len(m) > 0 {
			_ = gconv.Struct(m, &node)
		}
	}
	// Default value checks.
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	if node.Protocol == "" {
		node.Protocol = defaultProtocol
	}
	return node
}

// ImportWebMethod 导入前端方法
func ImportWebMethod(vs []string) string {
	vs = convert.UniqueSlice(vs)
	str := "{ " + strings.Join(vs, ", ") + " }"
	str = strings.TrimSuffix(str, ", ")
	return str
}

// CheckTreeTableFields 检查树表字段
func CheckTreeTableFields(columns []*sysin.GenCodesColumnListModel) (err error) {
	var fields = gutil.Copy(defaultTreeFields).([]string)
	for _, v := range columns {
		if validate.InSlice(fields, v.Name) {
			fields = convert.RemoveSlice(fields, v.Name)
		}
	}

	if len(fields) > 0 {
		err = gerror.Newf("树表必须包含[%v]字段", strings.Join(fields, "、"))
		return err
	}
	return
}

// CheckIllegalName 检查命名是否合理
func CheckIllegalName(errPrefix string, names ...string) (err error) {
	reg, _ := regexp.Compile("^[a-z_][a-z0-9_]*$")
	for _, name := range names {
		name = strings.ToLower(name)
		match := reg.MatchString(name)
		if !match {
			err = gerror.Newf("%v存在格式不正确，必须全部小写且由字母、数字和下划线组成:%v", errPrefix, name)
			return
		}
		if strings.HasSuffix(name, "test") {
			err = gerror.Newf("%v当中不能以`test`结尾:%v", errPrefix, name)
			return
		}
		if StartsWithDigit(name) {
			err = gerror.Newf("%v当中不能以阿拉伯数字开头:%v", errPrefix, name)
			return
		}
	}
	return
}

func StartsWithDigit(s string) bool {
	r := []rune(s)
	if len(r) > 0 {
		return unicode.IsDigit(r[0])
	}
	return false
}

// IsPidName 是否是树表的pid字段
func IsPidName(name string) bool {
	return name == "pid"
}

func ToTSArray(vs []string) string {
	formattedStrings := make([]string, len(vs))
	for i, str := range vs {
		formattedStrings[i] = fmt.Sprintf("'%s'", str)
	}
	return fmt.Sprintf("[%s]", strings.Join(formattedStrings, ", "))
}

func FormatGo(ctx context.Context, name, code string) (string, error) {
	path := GetTempGeneratePath(ctx) + "/" + name
	if err := gfile.PutContents(path, code); err != nil {
		return "", err
	}
	res, err := imports.Process(path, []byte(code), nil)
	if err != nil {
		err = gerror.Newf(`FormatGo error format "%s" go files: %v`, path, err)
		return "", err
	}
	return string(res), nil
}

func FormatVue(code string) string {
	if formatted, ok := tryPrettierFormat(code, "vue"); ok {
		return formatted
	}

	endTag := `</template>`
	vueLen := gstr.PosR(code, endTag)
	vueCode := code[:vueLen+len(endTag)]
	tsCode := code[vueLen+len(endTag):]

	vueCode = gohtml.Format(vueCode)
	vueCode = formatVueTemplate(vueCode)
	tsCode = FormatTs(tsCode)
	return vueCode + tsCode
}

func FormatTs(code string) string {
	if formatted, ok := tryPrettierFormat(code, "typescript"); ok {
		return formatted
	}
	code = replaceEmptyLinesWithSpace(code)
	code = formatTypeScript(code)
	return code + "\n"
}

// tryPrettierFormat 尝试使用 Prettier 格式化代码
func tryPrettierFormat(code string, parser string) (string, bool) {
	webDir := gfile.Abs("./web")
	prettierBin := gfile.Join(webDir, "node_modules", ".bin", "prettier")

	if !gfile.Exists(prettierBin) {
		return "", false
	}

	tmpFile := gfile.Temp(gtime.TimestampNanoStr()) + "." + getFileExt(parser)
	defer gfile.Remove(tmpFile)

	if err := gfile.PutContents(tmpFile, code); err != nil {
		return "", false
	}

	cmd := fmt.Sprintf("cd %s && %s --write %s 2>/dev/null", webDir, prettierBin, tmpFile)
	_, err := gproc.ShellExec(context.Background(), cmd)
	if err != nil {
		return "", false
	}

	formatted := gfile.GetContents(tmpFile)
	if formatted == "" {
		return "", false
	}
	return formatted, true
}

// getFileExt 根据 parser 类型返回文件扩展名
func getFileExt(parser string) string {
	switch parser {
	case "vue":
		return "vue"
	case "typescript":
		return "ts"
	default:
		return "txt"
	}
}

// formatVueTemplate 格式化 Vue 模板代码
func formatVueTemplate(code string) string {
	lines := strings.Split(code, "\n")
	var result []string
	lastLineWasOpenTag := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if lastLineWasOpenTag && trimmed == "" {
			continue
		}

		if strings.HasPrefix(trimmed, "<") && !strings.HasPrefix(trimmed, "</") &&
			!strings.HasSuffix(trimmed, "/>") && !strings.HasSuffix(trimmed, ">") &&
			!strings.HasPrefix(trimmed, "<!--") {
			lastLineWasOpenTag = true
			result = append(result, line)
			continue
		}

		lastLineWasOpenTag = false
		if trimmed == "/>" {
			if len(result) > 0 {
				prevLine := result[len(result)-1]
				re := regexp.MustCompile(`^(\s*)`)
				matches := re.FindStringSubmatch(prevLine)
				if len(matches) > 1 {
					line = matches[1] + "/>"
				}
			}
		} else if strings.Contains(line, "/>") {
			re := regexp.MustCompile(`\s*/>`)
			line = re.ReplaceAllString(line, " />")
		}

		if strings.HasPrefix(trimmed, "<") && !strings.HasPrefix(trimmed, "</") &&
			!strings.HasPrefix(trimmed, "<!--") && len(trimmed) > 100 {
			line = formatLongVueTag(line)
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

// formatLongVueTag 格式化过长的 Vue 标签
func formatLongVueTag(line string) string {
	re := regexp.MustCompile(`\s+`)
	parts := strings.SplitN(line, ">", 2)
	if len(parts) == 2 {
		tagPart := re.ReplaceAllString(parts[0], " ")
		tagPart = strings.TrimSpace(tagPart)
		return tagPart + ">" + parts[1]
	}
	return line
}

// formatTypeScript 格式化 TypeScript 代码
func formatTypeScript(code string) string {
	lines := strings.Split(code, "\n")
	var result []string
	inClass := false
	lastWasField := false

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "export class") {
			inClass = true
			result = append(result, line)
			continue
		}

		if inClass && (strings.HasPrefix(trimmed, "constructor(") ||
			strings.HasPrefix(trimmed, "public ") && strings.Contains(trimmed, "(")) {
			if lastWasField && len(result) > 0 {
				result = append(result, "")
			}
			inClass = false
			lastWasField = false
			result = append(result, line)
			continue
		}

		if inClass && strings.HasPrefix(trimmed, "public ") && strings.Contains(trimmed, "=") {
			if lastWasField && len(result) > 0 && strings.TrimSpace(result[len(result)-1]) == "" {
				result = result[:len(result)-1]
			}
			result = append(result, line)
			lastWasField = true
			continue
		}

		if trimmed == "}" && lastWasField {
			result = append(result, "")
			lastWasField = false
		}

		if trimmed == "" {
			if inClass && lastWasField {
				continue
			}
			if i > 0 && len(result) > 0 && strings.TrimSpace(result[len(result)-1]) == "" {
				continue
			}
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

func replaceEmptyLinesWithSpace(input string) string {
	re := regexp.MustCompile(`\n\s*\n\s*\n+`)
	result := re.ReplaceAllString(input, "\n\n")
	return result
}

func GetTempGeneratePath(ctx context.Context) string {
	return gfile.Abs(gfile.Temp() + "/hotgo-generate/" + simple.AppName(ctx))
}
