package tcpserver

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/msgin"
)

// OnAuthSummary 获取授权信息
func (s *sTCPServer) OnAuthSummary(ctx context.Context, args ...interface{}) {
	var (
		in     *msgin.AuthSummary
		user   = tcp.GetCtx(ctx)
		res    = new(msgin.ResponseAuthSummary)
		models *entity.SysServeLicense
	)

	if err := gconv.Scan(args, &in); err != nil {
		res.Code = 1
		res.Message = err.Error()
		s.serv.Reply(ctx, res)
		return
	}

	if user.Auth == nil {
		res.Code = 2
		res.Message = "登录信息获取失败，请重新登录"
		s.serv.Reply(ctx, res)
		return
	}

	if err := dao.SysServeLicense.Ctx(ctx).Where("appid = ?", user.Auth.AppId).Scan(&models); err != nil {
		res.Code = 3
		res.Message = err.Error()
		s.serv.Reply(ctx, res)
		return
	}

	if models == nil {
		res.Code = 4
		res.Message = "授权信息不存在"
		s.serv.Reply(ctx, res)
		return
	}

	if models.Status != consts.StatusEnabled {
		res.Code = 5
		res.Message = "授权已禁用，请联系管理员"
		s.serv.Reply(ctx, res)
		return
	}

	if models.Group != user.Auth.Group {
		res.Code = 6
		res.Message = "你登录的授权分组未得到授权，请联系管理员"
		s.serv.Reply(ctx, res)
		return
	}

	if models.EndAt.Before(gtime.Now()) {
		res.Code = 7
		res.Message = "授权已过期，请联系管理员"
		s.serv.Reply(ctx, res)
		return
	}

	res.Data = new(msgin.AuthSummaryData)
	res.Data.EndAt = models.EndAt
	res.Data.Online = models.Online
	s.serv.Reply(ctx, res)
}
