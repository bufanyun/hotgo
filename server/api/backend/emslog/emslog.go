package emslog

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
)

// ListReq 查询列表
type ListReq struct {
	g.Meta `path:"/emsLog/list" method:"get" tags:"邮件记录" summary:"获取邮件记录列表"`
	sysin.EmsLogListInp
}

type ListRes struct {
	List []*sysin.EmsLogListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

// ViewReq 获取指定信息
type ViewReq struct {
	g.Meta `path:"/emsLog/view" method:"get" tags:"邮件记录" summary:"获取指定信息"`
	sysin.EmsLogViewInp
}
type ViewRes struct {
	*sysin.EmsLogViewModel
}

// EditReq 修改/新增数据
type EditReq struct {
	g.Meta `path:"/emsLog/edit" method:"post" tags:"邮件记录" summary:"修改/新增邮件记录"`
	sysin.EmsLogEditInp
}
type EditRes struct{}

// DeleteReq 删除
type DeleteReq struct {
	g.Meta `path:"/emsLog/delete" method:"post" tags:"邮件记录" summary:"删除邮件记录"`
	sysin.EmsLogDeleteInp
}
type DeleteRes struct{}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/emsLog/status" method:"post" tags:"邮件记录" summary:"更新邮件记录状态"`
	sysin.EmsLogStatusInp
}
type StatusRes struct{}

// SendTestReq 更新状态
type SendTestReq struct {
	g.Meta `path:"/emsLog/sendTest" method:"post" tags:"邮件记录" summary:"发送测试邮件"`
	sysin.SendEmsInp
}
type SendTestRes struct{}
