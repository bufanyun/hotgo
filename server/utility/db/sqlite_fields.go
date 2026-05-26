package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
)

// func getSQLiteSchemaByCli(tableName string, db gdb.DB) (string, error) {
// 	dbConf := db.GetConfig()
// 	cmd := exec.Command("sqlite3", dbConf.Name, fmt.Sprintf(".schema %s", tableName))
// 	glog.Info(context.TODO(), "sqlite3", dbConf.Name, fmt.Sprintf("'.schema %s'", tableName))
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}
// 	return strings.TrimSpace(string(output)), nil
// }

func getSQLiteSchemaBySql(ctx context.Context, tableName string, db gdb.DB) (string, error) {
	schemaRes, err := db.GetValue(ctx, fmt.Sprintf(`SELECT sql FROM sqlite_master WHERE type='table' AND name='%s';`, tableName))
	if err != nil {
		return "", err
	}
	return schemaRes.String(), nil
}

func getSQliteTableComments(createTableSql string) (tableComment, tableName string) {
	// 按照换行符分割文本
	lines := strings.Split(createTableSql, "\n")

	// 循环输出每一行
	for _, line := range lines {
		// 检查 createTableSql 是否包含 comment
		if strings.Contains(line, "--") {
			if strings.Contains(line, "CREATE TABLE") {
				tableName = getLastWord(strings.Split(line, "(")[0])
				tableComment = strings.Split(line, "--")[1]
			}
		}
	}
	return
}

func getSQliteFieldsComments(createTableSql string) (fieldCommentMap gmap.Map, tableComment, tableName string) {
	// 按照换行符分割文本
	lines := strings.Split(createTableSql, "\n")

	// 循环输出每一行
	for _, line := range lines {
		// 检查 createTableSql 是否包含 comment
		if strings.Contains(line, "--") {
			if strings.Contains(line, "CREATE TABLE") {
				tableName = getLastWord(strings.Split(line, "(")[0])
				tableComment = strings.Split(line, "--")[1]
			} else {
				firstWord := getFirstWord(line)
				lastWord := getLastWord(line)
				fieldCommentMap.Set(firstWord, lastWord)
			}
		}
	}
	return
}

func transSqliteTablesComment(ctx context.Context, tableNames []string, db gdb.DB) (tables []*TableComment, err error) {
	schemaStr := ""
	eleIgnore := "sqlite_sequence"
	for _, v := range tableNames {
		if v != eleIgnore {
			schemaStr, err = getSQLiteSchemaBySql(ctx, v, db)
			if err != nil {
				return
			}
			comment, _ := getSQliteTableComments(schemaStr)
			tables = append(tables, &TableComment{
				Name:    v,
				Comment: comment,
			})
		}
	}
	return
}

func fixSqliteFieldsComment(ctx context.Context, tableName string, db gdb.DB, fields map[string]*gdb.TableField) (map[string]*gdb.TableField, error) {
	// 记录: db.DoSelect 无法执行 .开头的命令
	// schemaRes, err := db.DoSelect(ctx, dbConf.Link, fmt.Sprintf(`.schema %s`, d.QuoteWord(table)))
	// 记录: 查询 sqlite_master 不响应任何结果
	// s, err := db.Query(ctx, `select * from sqlite_master WHERE name="%s";`, tableName)
	// schemaStr, err := getSQLiteSchemaBySql(tableName, db)
	schemaStr, err := getSQLiteSchemaBySql(ctx, tableName, db)
	if err != nil {
		return fields, err
	}
	comments, _, _ := getSQliteFieldsComments(schemaStr)
	for i := range fields {
		if comments.Contains(i) {
			fields[i].Comment = comments.Get(i).(string)
		}
	}
	return fields, nil
}
