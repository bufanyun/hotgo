package db

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// 获取数据库表字段及注释
// usage:
//
// fields, err := db.GetFieldsWithComment(ctx, in.Table, in.Name)
//
//	if err != nil {
//		return
//	}
//
// for _, v := range fields {}
func GetFieldsWithComment(ctx context.Context, tableName, dbTag string) (fields map[string]*gdb.TableField, err error) {
	db := g.DB(dbTag)
	fields, err = db.TableFields(ctx, tableName) // 使用 goframe 框架本身已完美支持 mysql 获取表字段及注释
	dbConf := db.GetConfig()
	switch dbConf.Type {
	case "sqlite":
		fields, err = fixSqliteFieldsComment(ctx, tableName, db, fields)
	}
	return
}

type TableComment struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// 获取数据库表字段及注释
func GetTablesWithComment(ctx context.Context, dbTag string) (tables []*TableComment, err error) {
	db := g.DB(dbTag)
	dbConf := db.GetConfig()
	switch dbConf.Type {
	case "mysql":
		sql := "SELECT TABLE_NAME as name, TABLE_COMMENT as comment FROM information_schema.`TABLES` WHERE TABLE_SCHEMA = '%s'"
		if err = db.Ctx(ctx).Raw(sql).Scan(&tables); err != nil {
			return
		}
	case "pgsql":
		sql := `SELECT c.relnamespace, c.table_name, c.oid, obj_description(c.oid) AS comment 
			FROM
			(SELECT a.table_name, p.oid, p.relnamespace 
			FROM information_schema.tables  a 
			left join pg_class p on p.relname = a."table_name"
			WHERE a.table_catalog = '%s') c
			JOIN pg_namespace n ON n.oid = c.relnamespace`
		if err = db.Ctx(ctx).Raw(fmt.Sprintf(sql, dbConf.Name)).Scan(&tables); err != nil {
			return
		}
	case "sqlite":
		var tableNames []string
		tableNames, err = db.Tables(ctx)
		if err != nil {
			return
		}
		tables, err = transSqliteTablesComment(ctx, tableNames, db)
	}
	return
}
