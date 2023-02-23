// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
)

type sAdminMemberPost struct{}

func NewAdminMemberPost() *sAdminMemberPost {
	return &sAdminMemberPost{}
}

func init() {
	service.RegisterAdminMemberPost(NewAdminMemberPost())
}

func (s *sAdminMemberPost) UpdatePostIds(ctx context.Context, member_id int64, post_ids []int64) (err error) {
	_, err = dao.AdminMemberPost.Ctx(ctx).Where("member_id", member_id).Delete()
	if err != nil {
		err = gerror.Wrap(err, "删除失败")
		return err
	}

	for i := 0; i < len(post_ids); i++ {
		_, err = dao.AdminMemberPost.Ctx(ctx).
			Insert(entity.AdminMemberPost{
				MemberId: member_id,
				PostId:   post_ids[i],
			})
		if err != nil {
			err = gerror.Wrap(err, "插入用户岗位失败")
			return err
		}
	}

	return nil
}

// GetMemberByIds 获取指定用户的岗位ids
func (s *sAdminMemberPost) GetMemberByIds(ctx context.Context, member_id int64) (post_ids []int64, err error) {
	var list []*entity.AdminMemberPost
	err = dao.AdminMemberPost.Ctx(ctx).
		Fields("post_id").
		Where("member_id", member_id).
		Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return post_ids, err
	}

	for i := 0; i < len(list); i++ {
		post_ids = append(post_ids, list[i].PostId)
	}

	return post_ids, nil
}
