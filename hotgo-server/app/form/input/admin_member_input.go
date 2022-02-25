package input

import (
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

//  更新会员资料
type AdminMemberUpdateProfileInp struct {
	Mobile   int
	Email    string
	Realname string
}

//  获取指定会员资料
type AdminMemberProfileInp struct {
	Id int64
}
type AdminMemberProfileModel struct {
	PostGroup string                `json:"postGroup" description:"岗位名称"`
	RoleGroup string                `json:"roleGroup" description:"角色名称"`
	User      *AdminMemberViewModel `json:"user" description:"用户基本信息"`
	SysDept   *AdminDeptViewModel   `json:"sysDept" description:"部门信息"`
	SysRoles  []*AdminRoleListModel `json:"sysRoles" description:"角色列表"`
	PostIds   int64                 `json:"postIds" description:"当前岗位"`
	RoleIds   int64                 `json:"roleIds" description:"当前角色"`
}

//  更新会员资料
type MemberUpdateProfileInp struct {
	Mobile   int
	Email    string
	Realname string
}

//  修改登录密码
type AdminMemberUpdatePwdInp struct {
	Id          int64
	OldPassword string
	NewPassword string
}

//  重置密码
type AdminMemberResetPwdInp struct {
	Password string
	Id       int64
}

// 邮箱是否唯一
type AdminMemberEmailUniqueInp struct {
	Email string
	Id    int64
}

type AdminMemberEmailUniqueModel struct {
	IsUnique bool
}

// 手机号是否唯一
type AdminMemberMobileUniqueInp struct {
	Mobile string
	Id     int64
}

type AdminMemberMobileUniqueModel struct {
	IsUnique bool
}

// 名称是否唯一
type AdminMemberNameUniqueInp struct {
	Username string
	Id       int64
}

type AdminMemberNameUniqueModel struct {
	IsUnique bool
}

// 最大排序
type AdminMemberMaxSortInp struct {
	Id int64
}

type AdminMemberMaxSortModel struct {
	Sort int
}

//  修改/新增字典数据
type AdminMemberEditInp struct {
	Id                 int64       `json:"id"                   description:""`
	PostIds            []int64     `json:"postIds"       v:"required#岗位不能为空"           description:"岗位ID"`
	DeptId             int64       `json:"dept_id"       v:"required#部门不能为空"           description:"部门ID"`
	Username           string      `json:"username"   v:"required#账号不能为空"           description:"帐号"`
	Password           string      `json:"password"            description:"密码"`
	Realname           string      `json:"realname"             description:"真实姓名"`
	Avatar             string      `json:"avatar"               description:"头像"`
	Sex                string      `json:"sex"                  description:"性别[0:未知;1:男;2:女]"`
	Qq                 string      `json:"qq"                   description:"qq"`
	Email              string      `json:"email"                description:"邮箱"`
	Birthday           *gtime.Time `json:"birthday"             description:"生日"`
	ProvinceId         int         `json:"province_id"          description:"省"`
	CityId             int         `json:"city_id"              description:"城市"`
	AreaId             int         `json:"area_id"              description:"地区"`
	Address            string      `json:"address"              description:"默认地址"`
	Mobile             string      `json:"mobile"               description:"手机号码"`
	HomePhone          string      `json:"home_phone"           description:"家庭号码"`
	DingtalkRobotToken string      `json:"dingtalk_robot_token" description:"钉钉机器人token"`
	Role               int         `json:"role"            v:"required#角色不能为空"         description:"权限"`
	Remark             string      `json:"remark"               description:"备注"`
	Status             string      `json:"status"               description:"状态"`
	CreatedAt          *gtime.Time `json:"created_at"           description:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updated_at"           description:"修改时间"`
}

type AdminMemberAddInp struct {
	AdminMemberEditInp
	PasswordHash string `json:"password_hash"            description:"密码hash"`
	Salt         string `json:"salt"            description:"密码盐"`
}

type AdminMemberEditModel struct{}

//  删除字典类型
type AdminMemberDeleteInp struct {
	Id interface{}
}
type AdminMemberDeleteModel struct{}

// 获取信息
type AdminMemberViewInp struct {
	Id int64
}

type AdminMemberViewModel struct {
	entity.AdminMember
}

//  获取列表
type AdminMemberListInp struct {
	Page      int
	Limit     int
	Name      string
	Code      string
	DeptId    int
	Mobile    int
	Username  string
	Realname  string
	StartTime string
	EndTime   string
	Status    int
}

type AdminMemberListModel struct {
	entity.AdminMember
	DeptName string `json:"dept_name"`
	RoleName string `json:"role_name"`
}

// 登录
type AdminMemberLoginSignInp struct {
	Username string
	Password string
	Device   string
	Cid      string
	Code     string
}
type AdminMemberLoginSignModel struct {
	model.Identity
	Token string `json:"token" v:""  description:"登录token"`
}
