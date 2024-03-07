// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/dict"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
)

type sAdminPost struct{}

func NewAdminPost() *sAdminPost {
	return &sAdminPost{}
}

func init() {
	service.RegisterAdminPost(NewAdminPost())
	dict.RegisterFunc("adminPostOption", "岗位选项", service.AdminPost().Option)
}

// Delete 删除
func (s *sAdminPost) Delete(ctx context.Context, in *adminin.PostDeleteInp) (err error) {
	exist, err := dao.AdminMemberPost.Ctx(ctx).Where(dao.AdminMemberPost.Columns().PostId, in.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	if !exist.IsEmpty() {
		return gerror.New("请先解除该岗位下所有已关联用户关联关系！")
	}

	_, err = dao.AdminPost.Ctx(ctx).WherePri(in.Id).Delete()
	return
}

// VerifyUnique 验证部门唯一属性
func (s *sAdminPost) VerifyUnique(ctx context.Context, in *adminin.VerifyUniqueInp) (err error) {
	if in.Where == nil {
		return
	}

	cols := dao.AdminPost.Columns()
	msgMap := g.MapStrStr{
		cols.Name: "岗位名称已存在，请换一个",
		cols.Code: "岗位编码已存在，请换一个",
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf("字段 [ %v ] 未配置唯一属性验证", k)
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.AdminPost, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}

// Edit 修改/新增
func (s *sAdminPost) Edit(ctx context.Context, in *adminin.PostEditInp) (err error) {
	// 验证唯一性
	err = s.VerifyUnique(ctx, &adminin.VerifyUniqueInp{
		Id: in.Id,
		Where: g.Map{
			dao.AdminPost.Columns().Name: in.Name,
			dao.AdminPost.Columns().Code: in.Code,
		},
	})
	if err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		_, err = dao.AdminPost.Ctx(ctx).WherePri(in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = dao.AdminPost.Ctx(ctx).Data(in).Insert()
	return
}

// MaxSort 最大排序
func (s *sAdminPost) MaxSort(ctx context.Context, in *adminin.PostMaxSortInp) (res *adminin.PostMaxSortModel, err error) {
	if in.Id > 0 {
		if err = dao.AdminPost.Ctx(ctx).WherePri(in.Id).OrderDesc(dao.AdminPost.Columns().Sort).Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
	}

	if res == nil {
		res = new(adminin.PostMaxSortModel)
	}
	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取指定岗位信息
func (s *sAdminPost) View(ctx context.Context, in *adminin.PostViewInp) (res *adminin.PostViewModel, err error) {
	err = dao.AdminPost.Ctx(ctx).WherePri(in.Id).Scan(&res)
	return
}

// List 获取列表
func (s *sAdminPost) List(ctx context.Context, in *adminin.PostListInp) (list []*adminin.PostListModel, totalCount int, err error) {
	mod := dao.AdminPost.Ctx(ctx)
	cols := dao.AdminPost.Columns()

	// 访问路径
	if in.Name != "" {
		mod = mod.WhereLike(cols.Name, "%"+in.Name+"%")
	}

	// 模块
	if in.Code != "" {
		mod = mod.Where(cols.Code, in.Code)
	}

	// 请求方式
	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取岗位列表失败！")
		return
	}

	if totalCount == 0 {
		return
	}

	err = mod.Page(in.Page, in.PerPage).OrderAsc(cols.Sort).Scan(&list)
	return
}

// Option 岗位选项
func (s *sAdminPost) Option(ctx context.Context) (opts []*model.Option, err error) {
	var list []*entity.AdminPost
	if err = dao.AdminPost.Ctx(ctx).OrderAsc(dao.AdminPost.Columns().Sort).Scan(&list); err != nil {
		return nil, err
	}

	if len(list) == 0 {
		opts = make([]*model.Option, 0)
		return
	}

	for _, v := range list {
		opts = append(opts, dict.GenHashOption(v.Id, v.Name))
	}
	return
}

// GetMemberByStartName 获取指定用户的第一岗位
func (s *sAdminPost) GetMemberByStartName(ctx context.Context, memberId int64) (name string, err error) {
	// 默认取第一岗位
	postId, err := dao.AdminMemberPost.Ctx(ctx).Fields(dao.AdminMemberPost.Columns().PostId).Where(dao.AdminMemberPost.Columns().MemberId, memberId).Limit(1).Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	val, err := dao.AdminPost.Ctx(ctx).Fields(dao.AdminPost.Columns().Name).WherePri(postId.Int()).OrderDesc(dao.AdminPost.Columns().Id).Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return val.String(), nil
}

// Status 更新状态
func (s *sAdminPost) Status(ctx context.Context, in *adminin.PostStatusInp) (err error) {
	_, err = dao.AdminPost.Ctx(ctx).WherePri(in.Id).Data(dao.AdminPost.Columns().Status, in.Status).Update()
	return
}
