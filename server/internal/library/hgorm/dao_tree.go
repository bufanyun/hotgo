// Package hgorm
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hgorm

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model"
	"hotgo/utility/tree"
	"hotgo/utility/validate"
)

// GenSubTree 生成下级关系树
func GenSubTree(ctx context.Context, dao daoInstance, oldPid int64) (newPid int64, newLevel int, subTree string, err error) {
	if err = CheckTreeTable(ctx, dao); err != nil {
		return
	}

	if oldPid <= 0 {
		return 0, 1, "", nil
	}

	var models *model.DefaultTree
	if err = dao.Ctx(ctx).WherePri(oldPid).Scan(&models); err != nil {
		return 0, 0, "", err
	}

	if models == nil {
		return 0, 0, "", gerror.New("上级信息不存在")
	}

	newPid = oldPid
	newLevel = models.Level + 1
	subTree = tree.GenLabel(models.Tree, oldPid)
	return
}

// CheckTreeTable 检查树表
func CheckTreeTable(ctx context.Context, dao daoInstance) (err error) {
	fields, err := GetFieldsToSlice(ctx, dao)
	if err != nil {
		return err
	}

	if !validate.InSlice(fields, "pid") {
		return gerror.New("树表必须包含`pid`字段")
	}

	if !validate.InSlice(fields, "level") {
		return gerror.New("树表必须包含`level`字段")
	}

	if !validate.InSlice(fields, "tree") {
		return gerror.New("树表必须包含`tree`字段")
	}
	return
}

// AutoUpdateTree 自动更新关系树
func AutoUpdateTree(ctx context.Context, dao daoInstance, id, pid int64) (newPid int64, newLevel int, newTree string, err error) {
	if err = CheckTreeTable(ctx, dao); err != nil {
		return
	}

	if pid <= 0 {
		newPid = 0
		newLevel = 1
		newTree = ""
	} else {
		var pd *model.DefaultTree
		if err = dao.Ctx(ctx).WherePri(pid).Scan(&pd); err != nil {
			return 0, 0, "", err
		}

		if pd == nil {
			return 0, 0, "", gerror.New("未查询到树表上级信息，请检查！")
		}

		if id > 0 && validate.InSlice(tree.GetIds(pd.Tree), id) {
			return 0, 0, "", gerror.New("上级不能设为自己当前的子级！")
		}

		newPid = pid
		newLevel = pd.Level + 1
		newTree = tree.GenLabel(pd.Tree, pid)
	}

	if id > 0 {
		if pid == id {
			return 0, 0, "", gerror.New("上级不能是自己！")
		}

		var models *model.DefaultTree
		if err = dao.Ctx(ctx).WherePri(id).Scan(&models); err != nil {
			return 0, 0, "", err
		}

		if models == nil {
			return 0, 0, "", gerror.New("树表信息不存在，请检查！")
		}

		// 上级发生变化时，遍历修改其所有的下级关系树
		if models.Pid != pid {
			if err = updateChildrenTree(ctx, dao, models.Id, newLevel, newTree); err != nil {
				return
			}
		}
	}
	return
}

// updateChildrenTree 更新下级关系树
func updateChildrenTree(ctx context.Context, dao daoInstance, pid int64, pLevel int, pTree string) (err error) {
	var list []*model.DefaultTree
	if err = dao.Ctx(ctx).Where("pid", pid).Scan(&list); err != nil {
		return
	}

	if len(list) == 0 {
		return
	}

	newLevel := pLevel + 1
	newTree := tree.GenLabel(pTree, pid)

	var updateIds []int64
	for _, v := range list {
		updateIds = append(updateIds, v.Id)
		if err = updateChildrenTree(ctx, dao, v.Id, newLevel, newTree); err != nil {
			return
		}
	}

	if len(updateIds) > 0 {
		update := g.Map{
			"level": newLevel,
			"tree":  newTree,
		}
		_, err = dao.Ctx(ctx).WhereIn("id", updateIds).Data(update).Update()
	}
	return
}
