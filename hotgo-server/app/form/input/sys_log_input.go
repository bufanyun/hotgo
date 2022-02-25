package input

import (
	"github.com/bufanyun/hotgo/app/model/entity"
)

//  获取菜单列表
type LogListInp struct {
	Page       int
	Limit      int
	Module     string
	MemberId   int
	TakeUpTime int
	Method     string
	Url        string
	Ip         string
	ErrorCode  string
	StartTime  string
	EndTime    string
}

type LogListModel struct {
	entity.SysLog
	MemberName string `json:"member_name"`
	Region     string `json:"region"`
}
