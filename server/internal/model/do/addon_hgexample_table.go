// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTable is the golang structure of table hg_addon_hgexample_table for DAO operations like Where/Data.
type AddonHgexampleTable struct {
	g.Meta      `orm:"table:hg_addon_hgexample_table, do:true"`
	Id          any         // ID
	Pid         any         // 上级ID
	Level       any         // 树等级
	Tree        any         // 关系树
	CategoryId  any         // 分类ID
	Flag        *gjson.Json // 标签
	Title       any         // 标题
	Description any         // 描述
	Content     any         // 内容
	Image       any         // 单图
	Images      *gjson.Json // 多图
	Attachfile  any         // 附件
	Attachfiles *gjson.Json // 多附件
	Map         *gjson.Json // 动态键值对
	Star        any         // 推荐星
	Price       any         // 价格
	Views       any         // 浏览次数
	ActivityAt  *gtime.Time // 活动时间
	StartAt     *gtime.Time // 开启时间
	EndAt       *gtime.Time // 结束时间
	Switch      any         // 开关
	Sort        any         // 排序
	Avatar      any         // 头像
	Sex         any         // 性别
	Qq          any         // qq
	Email       any         // 邮箱
	Mobile      any         // 手机号码
	Hobby       *gjson.Json // 爱好
	Channel     any         // 渠道
	CityId      any         // 所在城市
	Remark      any         // 备注
	Status      any         // 状态
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
