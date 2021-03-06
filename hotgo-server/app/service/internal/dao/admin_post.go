// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao/internal"
	"github.com/gogf/gf/v2/errors/gerror"
)

// adminPostDao is the data access object for table hg_admin_post.
// You can define custom methods on it to extend its functionality as you wish.
type adminPostDao struct {
	*internal.AdminPostDao
}

var (
	// AdminPost is globally public accessible object for table hg_admin_post operations.
	AdminPost = adminPostDao{
		internal.NewAdminPostDao(),
	}
)

// Fill with you ideas below.

//
//  @Title  判断名称是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   id
//  @Param   name
//  @Return  bool
//  @Return  error
//
func (dao *adminPostDao) IsUniqueName(ctx context.Context, id int64, name string) (bool, error) {
	var data *entity.AdminPost
	m := dao.Ctx(ctx).Where("name", name)

	if id > 0 {
		m = m.WhereNot("id", id)
	}

	if err := m.Scan(&data); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return false, err
	}

	if data == nil {
		return true, nil
	}

	return false, nil
}

//
//  @Title  判断编码是否唯一
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   id
//  @Param   code
//  @Return  bool
//  @Return  error
//
func (dao *adminPostDao) IsUniqueCode(ctx context.Context, id int64, code string) (bool, error) {
	var data *entity.AdminPost
	m := dao.Ctx(ctx).Where("code", code)

	if id > 0 {
		m = m.WhereNot("id", id)
	}

	if err := m.Scan(&data); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return false, err
	}

	if data == nil {
		return true, nil
	}

	return false, nil
}
