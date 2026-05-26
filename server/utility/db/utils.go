package db

import (
	"hotgo/internal/consts"
	"strings"
	"unicode"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

// 判断字符是否为字母、数字或下划线
func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// 获取一行文本中的第一个完整单词
func getFirstWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !isWordChar(r) })
	if len(fields) > 0 {
		return fields[0]
	}
	return ""
}

// 获取一行文本中的最后一个完整单词
func getLastWord(text string) string {
	fields := strings.FieldsFunc(text, func(r rune) bool { return !isWordChar(r) })
	if len(fields) > 0 {
		return fields[len(fields)-1]
	}
	return ""
}

// GetQuoteChar 获取数据库对应的引号字符
// MySQL使用反引号 `，PostgreSQL使用双引号 "
func GetQuoteChar() string {
	db := g.DB()
	if db == nil {
		return "`" // 默认MySQL
	}
	dbType := db.GetConfig().Type
	if dbType == consts.DBPgsql {
		return "\""
	}
	return "`"
}

// QuoteIdentifier 引用标识符（表名、字段名）
// 自动去除已有的引号，避免重复引用
func QuoteIdentifier(identifier string) string {
	identifier = gstr.Trim(identifier, "`\"")
	q := GetQuoteChar()
	return q + identifier + q
}

// QuoteField 引用完整字段（表名.字段名）
// 自动去除表名和字段名中可能已有的引号
func QuoteField(table, field string) string {
	table = gstr.Trim(table, "`\"")
	field = gstr.Trim(field, "`\"")
	q := GetQuoteChar()
	return q + table + q + "." + q + field + q
}

// QuoteFieldAs 引用字段并添加别名（表名.字段名 as 别名）
// 自动去除表名、字段名、别名中可能已有的引号
func QuoteFieldAs(table, field, alias string) string {
	table = gstr.Trim(table, "`\"")
	field = gstr.Trim(field, "`\"")
	alias = gstr.Trim(alias, "`\"")
	q := GetQuoteChar()
	return q + table + q + "." + q + field + q + " as " + q + alias + q
}
