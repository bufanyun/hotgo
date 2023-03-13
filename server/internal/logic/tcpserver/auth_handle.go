package tcpserver

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/msgin"
)

// onAuthSummary 获取授权信息
func (s *sTCPServer) onAuthSummary(args ...interface{}) {
	var (
		in     *msgin.AuthSummary
		client = args[1].(*tcp.ClientConn)
		res    = new(msgin.ResponseAuthSummary)
		models *entity.SysServeLicense
	)

	if err := gconv.Scan(args, &in); err != nil {
		s.serv.Logger.Infof(s.serv.Ctx, "onAuthSummary message Scan failed:%+v, args:%+v", err, args)
		return
	}

	if client.Auth == nil {
		res.Code = 1
		res.Message = "登录信息获取失败，请重新登录"
		s.serv.Write(client.Conn, res)
		return
	}

	if err := dao.SysServeLicense.Ctx(s.serv.Ctx).Where("appid = ?", client.Auth.AppId).Scan(&models); err != nil {
		res.Code = 2
		res.Message = err.Error()
		s.serv.Write(client.Conn, res)
		return
	}

	if models == nil {
		res.Code = 3
		res.Message = "授权信息不存在"
		s.serv.Write(client.Conn, res)
		return
	}

	if models.Status != consts.StatusEnabled {
		res.Code = 4
		res.Message = "授权已禁用，请联系管理员"
		s.serv.Write(client.Conn, res)
		return
	}

	if models.Group != client.Auth.Group {
		res.Code = 5
		res.Message = "你登录的授权分组未得到授权，请联系管理员"
		s.serv.Write(client.Conn, res)
		return
	}

	if models.EndAt.Before(gtime.Now()) {
		res.Code = 6
		res.Message = "授权已过期，请联系管理员"
		s.serv.Write(client.Conn, res)
		return
	}

	res.Data = new(msgin.AuthSummaryData)
	res.Data.EndAt = models.EndAt
	res.Data.Online = models.Online
	s.serv.Write(client.Conn, res)
}
