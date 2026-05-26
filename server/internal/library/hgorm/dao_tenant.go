// Package hgorm
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hgorm

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

// TenantRelation 租户关系
type TenantRelation struct {
	DeptType   string // 部门类型
	TenantId   int64  // 租户ID
	MerchantId int64  // 商户ID
	UserId     int64  // 用户ID
}

// GetTenantRelation 获取租户关系
func GetTenantRelation(ctx context.Context, memberId int64) (tr *TenantRelation, err error) {

	data, err := g.Model(gstr.Join([]string{dao.AdminMember.Table(), "u"}, " ")).Ctx(ctx).
		LeftJoin(gstr.Join([]string{dao.AdminDept.Table(), "d"}, " "), "u.dept_id=d.id").
		Fields("u.tree,d.type").
		Where("u.id", memberId).One()
	if err != nil {
		return nil, err
	}

	if data.IsEmpty() {
		err = gerror.Newf("未找到用户[%v]的租户关系,该用户不存在", memberId)
		return
	}

	ids := tree.GetIds(data["tree"].String())

	getRelationId := func(deptType string) (int64, error) {
		id, err := g.Model(gstr.Join([]string{dao.AdminMember.Table(), "u"}, " ")).Ctx(ctx).
			LeftJoin(gstr.Join([]string{dao.AdminDept.Table(), "d"}, " "), "u.dept_id=d.id").
			Fields("u.id").
			WhereIn("u.id", ids).Where("d.type", deptType).
			OrderAsc("u.level"). // 确保是第一关系
			Limit(1).
			Value()
		if err != nil {
			return 0, err
		}
		if id.Int64() < 1 {
			err = gerror.Newf("未找到有效的租户关系,memberId:%v,deptType:%v", memberId, deptType)
			return 0, err
		}
		return id.Int64(), nil
	}

	tr = new(TenantRelation)
	tr.DeptType = data["type"].String()
	switch tr.DeptType {
	// 公司
	case consts.DeptTypeCompany:
		return
	// 租户
	case consts.DeptTypeTenant:
		tr.TenantId = memberId

		// 商户
	case consts.DeptTypeMerchant:
		tr.TenantId, err = getRelationId(consts.DeptTypeTenant)
		if err != nil {
			return nil, err
		}
		tr.MerchantId = memberId
		// 用户
	case consts.DeptTypeUser:
		tr.TenantId, err = getRelationId(consts.DeptTypeTenant)
		if err != nil {
			return nil, err
		}
		tr.MerchantId, err = getRelationId(consts.DeptTypeMerchant)
		if err != nil {
			return nil, err
		}
		tr.UserId = memberId
	default:
		err = gerror.Newf("未找到用户[%v]的租户关系,部门类型[%v] 无效", memberId, tr.DeptType)
		return nil, err
	}
	return
}
