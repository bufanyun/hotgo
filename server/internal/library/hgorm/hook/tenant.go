// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hook

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/utility/convert"
)

// SaveTenant 自动维护更新租户关系字段
// 根据部门类型识别当前租户、商户、用户身份
var SaveTenant = gdb.HookHandler{
	Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
		h, err := newHookSaveTenant(ctx, in)
		if err != nil {
			return nil, err
		}
		return h.handle()
	},
	Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
		h, err := newHookSaveTenant(ctx, in)
		if err != nil {
			return nil, err
		}
		return h.handle()
	},
}

type hookSaveTenant struct {
	ctx         context.Context
	in          any
	isNewRecord bool
	relations   map[int64]*hgorm.TenantRelation
}

func newHookSaveTenant(ctx context.Context, in any) (*hookSaveTenant, error) {
	h := new(hookSaveTenant)
	h.ctx = ctx
	h.in = in
	_, h.isNewRecord = in.(*gdb.HookInsertInput)
	h.relations = make(map[int64]*hgorm.TenantRelation)
	return h, nil
}

// getFields 获取表字段
func (h *hookSaveTenant) getFields() []string {
	if h.isNewRecord {
		in := h.in.(*gdb.HookInsertInput)
		return convert.EscapeFieldsToSlice(in.Model.GetFieldsStr())
	}
	in := h.in.(*gdb.HookUpdateInput)
	return convert.EscapeFieldsToSlice(in.Model.GetFieldsStr())
}

// getRelation 获取指定用户的租户关系
func (h *hookSaveTenant) getRelation(id int64) (tr *hgorm.TenantRelation, err error) {
	v, ok := h.relations[id]
	if ok {
		return v, nil
	}

	h.relations[id], err = hgorm.GetTenantRelation(h.ctx, id)
	if err != nil {
		return nil, err
	}
	return h.relations[id], nil
}

// getData 获取更新数据
func (h *hookSaveTenant) getData() any {
	if h.isNewRecord {
		in := h.in.(*gdb.HookInsertInput)
		return in.Data
	}
	in := h.in.(*gdb.HookUpdateInput)
	return in.Data
}

// setData 修改更新数据
func (h *hookSaveTenant) setData(data any) {
	if h.isNewRecord {
		in := h.in.(*gdb.HookInsertInput)
		in.Data = data.(gdb.List)
		return
	}
	in := h.in.(*gdb.HookUpdateInput)
	in.Data = data
}

func (h *hookSaveTenant) next() (result sql.Result, err error) {
	if h.isNewRecord {
		in := h.in.(*gdb.HookInsertInput)
		return in.Next(h.ctx)
	}
	in := h.in.(*gdb.HookUpdateInput)
	return in.Next(h.ctx)
}

// checkRelationConsistent 检查关系是否一致
func (h *hookSaveTenant) checkRelationConsistent(tid, mid, uid any) error {
	var (
		tenantId   = gconv.Int64(tid)
		merchantId = gconv.Int64(mid)
		userId     = gconv.Int64(uid)
	)

	// 存在用户，优先用用户开始检查
	if userId > 0 {
		tr, err := h.getRelation(userId)
		if err != nil {
			return err
		}
		if tenantId > 0 && tr.TenantId != tenantId {
			err = gerror.Newf("租户[%v]与用户[%v]关系不匹配", tenantId, userId)
			return err
		}
		if merchantId > 0 && tr.MerchantId != merchantId {
			err = gerror.Newf("商户[%v]与用户[%v]关系不匹配", merchantId, userId)
			return err
		}
		return nil
	}

	if merchantId > 0 {
		tr, err := h.getRelation(merchantId)
		if err != nil {
			return err
		}
		if tenantId > 0 && tr.TenantId != tenantId {
			err = gerror.Newf("租户[%v]与商户[%v]关系不匹配", tenantId, userId)
			return err
		}
		return nil
	}
	return nil
}

// checkRelationSingle 检查单个用户关系
func (h *hookSaveTenant) checkRelationSingle(idx any, relation, limitType string) (err error) {
	id := gconv.Int64(idx)
	if id < 1 {
		return
	}

	ok := false
	tr, err := h.getRelation(id)
	if err != nil {
		return err
	}

	if tr.DeptType != limitType {
		err = gerror.Newf("用户[%v]关系验证不通过,类型身份不匹配[%v != %v]", id, tr.DeptType, limitType)
		return
	}

	relationId := contexts.GetUserId(h.ctx)
	switch relation {
	case consts.DeptTypeTenant:
		if ok = tr.TenantId == relationId; !ok {
			err = gerror.Newf("%v的租户不是%v", id, relationId)
			return
		}
	case consts.DeptTypeMerchant:
		if ok = tr.MerchantId == relationId; !ok {
			err = gerror.Newf("%v的商户不是%v", id, relationId)
			return
		}
	}
	return
}

// 检查用户关系是否有效
func (h *hookSaveTenant) checkRelation(deptType string, data gdb.Map) (err error) {
	switch deptType {
	// 公司和用户，检查关系是否一致
	case consts.DeptTypeCompany, consts.DeptTypeUser:
		if err = h.checkRelationConsistent(data[consts.TenantId], data[consts.MerchantId], data[consts.UserId]); err != nil {
			return
		}
	// 租户，检查商户和用户是否属于自己
	case consts.DeptTypeTenant:
		if err = h.checkRelationSingle(data[consts.MerchantId], consts.DeptTypeTenant, consts.DeptTypeMerchant); err != nil {
			return
		}
		if err = h.checkRelationSingle(data[consts.UserId], consts.DeptTypeTenant, consts.DeptTypeUser); err != nil {
			return
		}
		// 商户，检查用户是否属于自己
	case consts.DeptTypeMerchant:
		if err = h.checkRelationSingle(data[consts.UserId], consts.DeptTypeMerchant, consts.DeptTypeUser); err != nil {
			return
		}
	}
	return
}

func (h *hookSaveTenant) handle() (result sql.Result, err error) {
	var (
		update   = make(g.Map)
		fields   = h.getFields()
		memberId = contexts.GetUserId(h.ctx)
		deptType = contexts.GetDeptType(h.ctx)
		tr       *hgorm.TenantRelation
	)

	if memberId == 0 || len(deptType) == 0 {
		err = gerror.New("缺少用户上下文数据")
		return nil, err
	}

	// 非公司类型，加载自己的租户关系，用于重写关系
	if !contexts.IsCompanyDept(h.ctx) {
		tr, err = h.getRelation(memberId)
		if err != nil {
			return nil, err
		}
	}

	switch deptType {
	// 公司
	case consts.DeptTypeCompany:

		// 租户
	case consts.DeptTypeTenant:
		if gstr.InArray(fields, consts.TenantId) {
			update[consts.TenantId] = tr.TenantId
		}

		// 商户
	case consts.DeptTypeMerchant:
		if gstr.InArray(fields, consts.TenantId) {
			update[consts.TenantId] = tr.TenantId
		}
		if gstr.InArray(fields, consts.MerchantId) {
			update[consts.MerchantId] = tr.MerchantId
		}

		// 用户
	case consts.DeptTypeUser:
		if gstr.InArray(fields, consts.TenantId) {
			update[consts.TenantId] = tr.TenantId
		}
		if gstr.InArray(fields, consts.MerchantId) {
			update[consts.MerchantId] = tr.MerchantId
		}
		if gstr.InArray(fields, consts.UserId) {
			update[consts.UserId] = tr.UserId
		}
	default:
		err = gerror.Newf("当前用户部门类型[%v] 找到有效的hook，请检查！", deptType)
		return nil, err
	}

	switch value := h.getData().(type) {
	case gdb.List:
		for i, data := range value {
			if err = h.checkRelation(deptType, data); err != nil {
				return nil, err
			}
			for k, v := range update {
				data[k] = v
			}
			value[i] = data
		}
		h.setData(value)
	case gdb.Map:
		if err = h.checkRelation(deptType, value); err != nil {
			return nil, err
		}
		for k, v := range update {
			value[k] = v
		}
		h.setData(value)
	}
	return h.next()
}
