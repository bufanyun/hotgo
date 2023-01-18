// Package hgorm
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package hgorm

// dao.
import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/utility/convert"
	"hotgo/utility/tree"
)

// GenJoinOnRelation 生成关联表关联条件
func GenJoinOnRelation(masterTable, masterField, joinTable, alias, onField string) []string {
	return []string{
		joinTable,
		alias,
		fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", alias, onField, masterTable, masterField),
	}
}

type daoInstance interface {
	Table() string
	Ctx(ctx context.Context) *gdb.Model
}

// Join 关联表属性
type Join struct {
	Dao    interface{}                // 关联表dao实例
	Alias  string                     // 别名
	fields map[string]*gdb.TableField // 表字段列表
}

// GenJoinSelect 生成关联表select
// 这里会将实体中的字段驼峰转为下划线于数据库进行匹配，意味着数据库字段必须全部是小写字母+下划线的格式
func GenJoinSelect(ctx context.Context, entity interface{}, masterDao interface{}, joins []*Join) (allFields string, err error) {
	var tmpFields []string

	md, ok := masterDao.(daoInstance)
	if !ok {
		err = errors.New("masterDao unimplemented interface format.daoInstance")
		return
	}

	if len(joins) == 0 {
		err = errors.New("JoinFields joins len = 0")
		return
	}

	for _, v := range joins {
		jd, ok := v.Dao.(daoInstance)
		if !ok {
			err = errors.New("joins index unimplemented interface format.daoInstance")
			return
		}
		v.fields, err = jd.Ctx(ctx).TableFields(jd.Table())
		if err != nil {
			return
		}
	}

	masterFields, err := md.Ctx(ctx).TableFields(md.Table())
	if err != nil {
		return
	}

	entityFields, err := convert.GetEntityFieldTags(entity)
	if err != nil {
		return "", err
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
			tmpFields = append(tmpFields, fmt.Sprintf("`%s`.`%s`", md.Table(), originalField))
			continue
		}
	}

	return gstr.Implode(",", convert.UniqueSliceString(tmpFields)), nil
}

// GenSelect 生成select
// 这里会将实体中的字段驼峰转为下划线于数据库进行匹配，意味着数据库字段必须全部是小写字母+下划线的格式
func GenSelect(ctx context.Context, entity interface{}, dao interface{}) (allFields string, err error) {
	var tmpFields []string

	md, ok := dao.(daoInstance)
	if !ok {
		err = errors.New("dao unimplemented interface format.daoInstance")
		return
	}

	fields, err := md.Ctx(ctx).TableFields(md.Table())
	if err != nil {
		return
	}

	entityFields, err := convert.GetEntityFieldTags(entity)
	if err != nil {
		return "", err
	}

	if len(entityFields) == 0 {
		return "*", nil
	}

	for _, field := range entityFields {
		originalField := gstr.CaseSnakeFirstUpper(field)
		if _, ok := fields[originalField]; ok {
			tmpFields = append(tmpFields, fmt.Sprintf("`%s`", originalField))
			continue
		}
	}

	return gstr.Implode(",", convert.UniqueSliceString(tmpFields)), nil
}

// GetPkField 获取dao实例中的主键名称
func GetPkField(ctx context.Context, dao daoInstance) (string, error) {
	fields, err := dao.Ctx(ctx).TableFields(dao.Table())
	if err != nil {
		return "", err
	}
	if len(fields) == 0 {
		return "", errors.New("field not found")
	}

	for _, field := range fields {
		if field.Key == "PRI" {
			return field.Name, nil
		}
	}

	return "", errors.New("no primary key")
}

// IsUnique 是否唯一
func IsUnique(ctx context.Context, dao interface{}, where g.Map, message string, pkId ...interface{}) error {
	d, ok := dao.(daoInstance)
	if !ok {
		return errors.New("IsUnique dao unimplemented interface format.daoInstance")
	}

	if len(where) == 0 {
		return errors.New("where condition cannot be empty")
	}

	m := d.Ctx(ctx).Where(where)
	if len(pkId) > 0 {
		field, err := GetPkField(ctx, d)
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
			for k, _ := range where {
				message = fmt.Sprintf("in the table：%s, %v not uniqued", d.Table(), where[k])
				break
			}
		}
		return errors.New(message)
	}
	return nil
}

// GenSubTree 生成下级关系树
func GenSubTree(ctx context.Context, dao interface{}, oldPid int64) (newPid int64, newLevel int, subTree string, err error) {
	// 顶级树
	if oldPid == 0 {
		return 0, 1, "", nil
	}

	d, ok := dao.(daoInstance)
	if !ok {
		return 0, 0, "", errors.New("GenTree dao unimplemented interface format.daoInstance")
	}
	field, err := GetPkField(ctx, d)
	if err != nil {
		return 0, 0, "", err
	}

	models, err := d.Ctx(ctx).WhereNot(field, oldPid).One()
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
