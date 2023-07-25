// Package hgorm
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hgorm

// dao.
import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/utility/convert"
	"hotgo/utility/tree"
)

type daoInstance interface {
	Table() string
	Ctx(ctx context.Context) *gdb.Model
}

// Join 关联表属性
type Join struct {
	Dao    daoInstance                // 关联表dao实例
	Alias  string                     // 别名
	fields map[string]*gdb.TableField // 表字段列表
}

func LeftJoin(m *gdb.Model, masterTable, masterField, joinTable, alias, onField string) *gdb.Model {
	return m.LeftJoin(GenJoinOnRelation(masterTable, masterField, joinTable, alias, onField)...)
}

// GenJoinOnRelation 生成关联表关联条件
func GenJoinOnRelation(masterTable, masterField, joinTable, alias, onField string) []string {
	relation := fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", alias, onField, masterTable, masterField)
	return []string{joinTable, alias, relation}
}

// GenJoinSelect 生成关联表select
// 这里会将实体中的字段驼峰转为下划线于数据库进行匹配，意味着数据库字段必须全部是小写字母+下划线的格式
func GenJoinSelect(ctx context.Context, entity interface{}, dao daoInstance, joins []*Join) (allFields string, err error) {
	var tmpFields []string
	if len(joins) == 0 {
		err = gerror.New("JoinFields joins len = 0")
		return
	}

	for _, v := range joins {
		v.fields, err = v.Dao.Ctx(ctx).TableFields(v.Dao.Table())
		if err != nil {
			return
		}
	}

	masterFields, err := dao.Ctx(ctx).TableFields(dao.Table())
	if err != nil {
		return
	}

	entityFields, err := convert.GetEntityFieldTags(entity)
	if err != nil {
		return
	}

	if len(entityFields) == 0 {
		return "*", nil
	}

	// 是否为关联表字段
	getJoinAttribute := func(field string) (*Join, string) {
		for _, v := range joins {
			if gstr.HasPrefix(field, v.Alias) {
				return v, gstr.CaseSnakeFirstUpper(gstr.StrEx(field, v.Alias))
			}
		}
		return nil, ""
	}

	for _, field := range entityFields {
		// 关联表
		jd, joinField := getJoinAttribute(field)
		if jd != nil {
			if _, ok := jd.fields[joinField]; ok {
				tmpFields = append(tmpFields, fmt.Sprintf("`%s`.`%s` as `%s`", jd.Alias, joinField, field))
				continue
			}
		}

		// 主表
		originalField := gstr.CaseSnakeFirstUpper(field)
		if _, ok := masterFields[originalField]; ok {
			tmpFields = append(tmpFields, fmt.Sprintf("`%s`.`%s`", dao.Table(), originalField))
			continue
		}
	}
	return gstr.Implode(",", convert.UniqueSlice(tmpFields)), nil
}

// GetPkField 获取dao实例中的主键名称
func GetPkField(ctx context.Context, dao daoInstance) (string, error) {
	fields, err := dao.Ctx(ctx).TableFields(dao.Table())
	if err != nil {
		return "", err
	}
	if len(fields) == 0 {
		return "", gerror.New("field not found")
	}

	for _, field := range fields {
		if field.Key == "PRI" {
			return field.Name, nil
		}
	}
	return "", gerror.New("no primary key")
}

// IsUnique 是否唯一
func IsUnique(ctx context.Context, dao daoInstance, where g.Map, message string, pkId ...interface{}) error {
	if len(where) == 0 {
		return gerror.New("where condition cannot be empty")
	}

	m := dao.Ctx(ctx).Where(where)
	if len(pkId) > 0 {
		field, err := GetPkField(ctx, dao)
		if err != nil {
			return err
		}
		m = m.WhereNot(field, pkId[0])
	}

	count, err := m.Count(1)
	if err != nil {
		return err
	}

	if count > 0 {
		if message == "" {
			for k := range where {
				message = fmt.Sprintf("in the table：%s, %v not uniqued", dao.Table(), where[k])
				break
			}
		}
		return gerror.New(message)
	}
	return nil
}

// GenSubTree 生成下级关系树
func GenSubTree(ctx context.Context, dao daoInstance, oldPid int64) (newPid int64, newLevel int, subTree string, err error) {
	// 顶级树
	if oldPid <= 0 {
		return 0, 1, "", nil
	}

	field, err := GetPkField(ctx, dao)
	if err != nil {
		return 0, 0, "", err
	}

	models, err := dao.Ctx(ctx).Where(field, oldPid).One()
	if err != nil {
		return 0, 0, "", err
	}

	if models.IsEmpty() {
		return 0, 0, "", gerror.New("上级信息不存在")
	}

	level, ok := models["level"]
	if !ok {
		return 0, 0, "", gerror.New("表中必须包含`level`字段")
	}

	supTree, ok := models["tree"]
	if !ok {
		return 0, 0, "", gerror.New("表中必须包含`tree`字段")
	}

	newPid = oldPid
	newLevel = level.Int() + 1
	subTree = tree.GenLabel(supTree.String(), oldPid)
	return
}
