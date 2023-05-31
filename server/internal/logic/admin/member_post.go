// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
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

// UpdatePostIds 更新用户岗位
func (s *sAdminMemberPost) UpdatePostIds(ctx context.Context, memberId int64, postIds []int64) (err error) {
	if _, err = dao.AdminMemberPost.Ctx(ctx).Where(dao.AdminMemberPost.Columns().MemberId, memberId).Delete(); err != nil {
		err = gerror.Wrap(err, "清理用户旧岗位数据失败，请稍后重试！")
		return
	}

	for i := 0; i < len(postIds); i++ {
		_, err = dao.AdminMemberPost.Ctx(ctx).Insert(entity.AdminMemberPost{
			MemberId: memberId,
			PostId:   postIds[i],
		})
		if err != nil {
			err = gerror.Wrap(err, "加入用户岗位数据失败，请稍后重试！")
			return err
		}
	}
	return
}
