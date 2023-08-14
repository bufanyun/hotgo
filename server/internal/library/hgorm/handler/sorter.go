// Package handler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package handler

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gutil"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"regexp"
)

// ISorter 排序器接口，实现该接口即可使用Handler匹配排序，支持多字段排序
type ISorter interface {
	GetSorters() []*form.Sorter
}

// Sorter 排序器
func Sorter(in ISorter) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		masterTable, ts := convert.GetModelTable(m)
		fields, err := m.TableFields(masterTable)
		if err != nil {
			g.Log().Panicf(m.GetCtx(), "failed to sorter TableFields err:%+v", err)
		}

		sorters := in.GetSorters()
		aliases := extractTableAliases(ts)
		if len(aliases) > 0 {
			var newSorters []*form.Sorter
			var removeIndex []int
			for as, table := range aliases {
				// 关联表
				fds, err := m.TableFields(table)
				if err != nil {
					g.Log().Panicf(m.GetCtx(), "failed to sorter TableFields err2:%+v", err)
				}

				var sorter2 []*form.Sorter
				for k, sorter := range sorters {
					if gstr.HasPrefix(sorter.ColumnKey, as) {
						sorter2 = append(sorter2, &form.Sorter{
							ColumnKey: gstr.Replace(sorter.ColumnKey, as, ""),
							Order:     sorter.Order,
						})
						removeIndex = append(removeIndex, k)
					}
				}

				if len(sorter2) > 0 {
					sorter2 = mappingAndFilterToTableFields(fds, sorter2)
					for _, v := range sorter2 {
						v.ColumnKey = fmt.Sprintf("`%v`.`%v`", as, v.ColumnKey)
					}
					newSorters = append(newSorters, sorter2...)
				}
			}

			// 移除关联表字段
			sorters = mappingAndFilterToTableFields(fields, removeSorterIndexes(sorters, removeIndex))
			for _, v := range sorters {
				v.ColumnKey = fmt.Sprintf("`%v`.`%v`", masterTable, v.ColumnKey)
			}

			sorters = append(newSorters, sorters...)
		} else {
			// 单表
			sorters = mappingAndFilterToTableFields(fields, sorters)
			for _, v := range sorters {
				v.ColumnKey = fmt.Sprintf("`%v`.`%v`", masterTable, v.ColumnKey)
			}
		}

		hasSort := false
		for _, sorter := range sorters {
			if len(sorter.ColumnKey) == 0 {
				continue
			}

			switch sorter.Order {
			case "descend": // 降序
				hasSort = true
				m = m.OrderDesc(sorter.ColumnKey)
			case "ascend": // 升序
				hasSort = true
				m = m.OrderAsc(sorter.ColumnKey)
			default:
				continue
			}
		}

		if hasSort {
			return m
		}

		// 不存在排序条件，默认使用主表主键做降序排序
		var pk string
		for name, field := range fields {
			if gstr.ContainsI(field.Key, consts.GenCodesIndexPK) {
				pk = name
				break
			}
		}

		// 没有主键
		if len(pk) == 0 {
			return m
		}

		// 存在别名，优先匹配别名
		if len(aliases) > 0 {
			for as, table := range aliases {
				if table == masterTable {
					return m.OrderDesc(fmt.Sprintf("`%v`.`%v`", as, pk))
				}
			}
		}
		return m.OrderDesc(fmt.Sprintf("`%v`.`%v`", masterTable, pk))
	}
}

// extractTableAliases 解析关联条件中的关联表别名
func extractTableAliases(ts string) map[string]string {
	re := regexp.MustCompile("`?([^`\\s]+)`?\\s+AS\\s+`?([^`\\s]+)`?\\s")
	matches := re.FindAllStringSubmatch(ts, -1)

	result := make(map[string]string)
	for _, match := range matches {
		result[match[2]] = match[1]
	}
	return result
}

// removeSorterIndexes 移除指定索引的排序器
func removeSorterIndexes(slice []*form.Sorter, indexes []int) []*form.Sorter {
	removed := make([]*form.Sorter, 0)
	indexMap := make(map[int]bool)

	for _, index := range indexes {
		indexMap[index] = true
	}

	for i, value := range slice {
		if !indexMap[i] {
			removed = append(removed, value)
		}
	}
	return removed
}

// mappingAndFilterToTableFields 将排序字段映射为实际的表字段
func mappingAndFilterToTableFields(fieldsMap map[string]*gdb.TableField, sorters []*form.Sorter) (ser []*form.Sorter) {
	if len(fieldsMap) == 0 {
		return
	}

	var fields []string
	for _, v := range sorters {
		fields = append(fields, v.ColumnKey)
	}

	fieldsKeyMap := make(map[string]interface{}, len(fieldsMap))
	for k := range fieldsMap {
		fieldsKeyMap[k] = nil
	}

	var inputFieldsArray = gstr.SplitAndTrim(gstr.Join(fields, ","), ",")
	for _, field := range inputFieldsArray {
		if _, ok := fieldsKeyMap[field]; ok {
			continue
		}

		if !gregex.IsMatchString(`^[\w\-]+$`, field) {
			continue
		}

		if foundKey, _ := gutil.MapPossibleItemByKey(fieldsKeyMap, field); foundKey != "" {
			for _, v := range sorters {
				if v.ColumnKey == field {
					v.ColumnKey = foundKey
				}
			}
		}
	}
	return sorters
}
