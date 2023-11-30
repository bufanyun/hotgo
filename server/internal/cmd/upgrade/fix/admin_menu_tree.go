// Package fix
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package fix

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
)

// UpdateAdminMenuTree 更新菜单关系树
// 根据树等级从上到下依次检查，将无效的关系树进行修复更新
func UpdateAdminMenuTree(ctx context.Context) {
	var list []*entity.AdminMenu
	err := dao.AdminMenu.Ctx(ctx).OrderAsc("level").Scan(&list)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	genUpdateData := func(v *entity.AdminMenu) g.Map {
		update := g.Map{"updated_at": gtime.Now()}
		if v.Pid <= 0 {
			update["level"] = 1
			update["tree"] = ""
			return update
		}

		// 生成下级关系树
		update["pid"], update["level"], update["tree"], err = hgorm.GenSubTree(ctx, &dao.AdminMenu, v.Pid)
		return update
	}

	for _, v := range list {
		update := genUpdateData(v)
		if v.Level == update["level"] && v.Tree == update["tree"] {
			continue
		}
		if _, err = dao.AdminMenu.Ctx(ctx).WherePri(v.Id).Data(genUpdateData(v)).Update(); err != nil {
			g.Log().Fatal(ctx, err)
		}
	}
}
