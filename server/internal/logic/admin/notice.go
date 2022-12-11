// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/charset"
	"hotgo/utility/convert"
	"hotgo/utility/simple"
	"strings"
)

type sAdminNotice struct{}

func NewAdminNotice() *sAdminNotice {
	return &sAdminNotice{}
}

func init() {
	service.RegisterAdminNotice(NewAdminNotice())
}

// Delete 删除
func (s *sAdminNotice) Delete(ctx context.Context, in adminin.NoticeDeleteInp) error {
	_, err := dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// Edit 修改/新增
func (s *sAdminNotice) Edit(ctx context.Context, in adminin.NoticeEditInp) (err error) {
	if in.Title == "" {
		err = gerror.New("标题不能为空")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	if in.Id > 0 {
		_, err = dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return err
		}

		return nil
	}

	// 新增
	in.CreatedAt = gtime.Now()
	_, err = dao.AdminNotice.Ctx(ctx).Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	// 推送通知
	memberIds := charset.SplitMemberIds(in.Receiver, ",")
	response := &websocket.WResponse{
		Event: "notice",
		Data:  in,
	}
	simple.SafeGo(ctx, func(ctx context.Context) {
		if len(memberIds) == 0 {
			websocket.SendToAll(response)
		} else {
			for _, memberId := range memberIds {
				websocket.SendToUser(memberId, response)
			}
		}
	})

	return nil
}

// Status 更新部门状态
func (s *sAdminNotice) Status(ctx context.Context, in adminin.NoticeStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return err
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return err
	}

	if !convert.InSliceInt(consts.StatusMap, in.Status) {
		err = gerror.New("状态不正确")
		return err
	}

	// 修改
	in.UpdatedAt = gtime.Now()
	_, err = dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// MaxSort 最大排序
func (s *sAdminNotice) MaxSort(ctx context.Context, in adminin.NoticeMaxSortInp) (*adminin.NoticeMaxSortModel, error) {
	var res adminin.NoticeMaxSortModel
	if in.Id > 0 {
		if err := dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Order("sort desc").Scan(&res); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return nil, err
		}
	}
	res.Sort = res.Sort + 10
	return &res, nil
}

// View 获取指定字典类型信息
func (s *sAdminNotice) View(ctx context.Context, in adminin.NoticeViewInp) (res *adminin.NoticeViewModel, err error) {
	if err = dao.AdminNotice.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	return res, nil
}

// List 获取列表
func (s *sAdminNotice) List(ctx context.Context, in adminin.NoticeListInp) (list []*adminin.NoticeListModel, totalCount int64, err error) {
	mod := dao.AdminNotice.Ctx(ctx)

	// 访问路径
	if in.Title != "" {
		mod = mod.WhereLike("title", "%"+in.Title+"%")
	}

	// 模块
	if in.Content != "" {
		mod = mod.WhereLike("content", "%"+in.Content+"%")
	}

	// 请求方式
	if in.Status > 0 {
		mod = mod.Where("status", in.Status)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, nil
	}

	if err = mod.Page(int(in.Page), int(in.PerPage)).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	for k, v := range list {
		list[k].ReceiveNum = len(strings.Split(v.Reader, ","))
	}
	return list, totalCount, err
}
