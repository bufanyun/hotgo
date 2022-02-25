package input

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
)

//  获取列表
type AdminRoleListInp struct {
	Page  int
	Limit int
}

type AdminRoleListModel struct {
	entity.AdminRole
}

//  查询列表
type AdminRoleMemberListInp struct {
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	Role      int    `json:"role"   description:"角色ID"`
	DeptId    int    `json:"dept_id"   description:"部门ID"`
	Mobile    int    `json:"mobile"   description:"手机号"`
	Username  string `json:"username"   description:"用户名"`
	Realname  string `json:"realname"   description:"真实姓名"`
	StartTime string `json:"start_time"   description:"开始时间"`
	EndTime   string `json:"end_time"   description:"结束时间"`
	Name      string `json:"name"   description:"岗位名称"`
	Code      string `json:"code"   description:"岗位编码"`
}

type AdminRoleMemberListModel []*AdminMemberListModel

//  查询角色菜单列表
type MenuRoleListInp struct {
	RoleId int64
}
type MenuRoleListModel struct {
	Menus       []*model.LabelTreeMenu `json:"menus"   description:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   description:"选择的菜单ID"`
}
