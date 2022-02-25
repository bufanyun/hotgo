package adminService

import (
	"context"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/bufanyun/hotgo/app/service/internal/dao"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var MemberPost = new(memberPost)

type memberPost struct{}

func (service *memberPost) UpdatePostIds(ctx context.Context, member_id int64, post_ids []int64) (err error) {
	_, err = dao.AdminMemberPost.Ctx(ctx).
		Where("member_id", member_id).
		Delete()
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
			err = gerror.Wrap(err, "插入会员岗位失败")
			return err
		}
	}

	return nil
}

//
//  @Title  获取指定会员的岗位ids
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   member_id
//  @Return  post_ids
//  @Return  err
//
func (service *memberPost) GetMemberByIds(ctx context.Context, member_id int64) (post_ids []int64, err error) {

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

	g.Log().Print(ctx, "post_ids:", post_ids)
	return post_ids, nil
}
