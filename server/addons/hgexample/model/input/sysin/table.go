// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

// TableEditInp 修改/新增
type TableEditInp struct {
	entity.AddonHgexampleTable
}

type TableEditModel struct{}

func (in *TableEditInp) Filter(ctx context.Context) (err error) {
	if in.Map.IsNil() {
		in.Map = gjson.New(consts.NilJsonToString)
	}
	if in.Flag.IsNil() {
		in.Flag = gjson.New(consts.NilJsonToString)
	}
	if in.Images.IsNil() {
		in.Images = gjson.New(consts.NilJsonToString)
	}
	if in.Attachfiles.IsNil() {
		in.Attachfiles = gjson.New(consts.NilJsonToString)
	}
	if in.Hobby.IsNil() {
		in.Hobby = gjson.New(consts.NilJsonToString)
	}

	if in.Title == "" {
		return errors.New("标题不能为空")
	}

	if in.Email != "" && !validate.IsEmail(in.Email) {
		return errors.New("邮箱格式不正确")
	}

	if err := g.Validator().Rules("float|between:0,5").Messages("请输入一个浮点数|推荐星只能是0~5星").Data(in.Star).Run(ctx); err != nil {
		return err.Current()
	}
	return
}

// TableDeleteInp 删除类型
type TableDeleteInp struct {
	Id interface{} `json:"id" v:"required#表格ID不能为空" dc:"表格ID"`
}

type TableDeleteModel struct{}

// TableViewInp 获取信息
type TableViewInp struct {
	Id int64 `json:"id" v:"required#表格ID不能为空" dc:"表格ID"`
}

type TableViewModel struct {
	entity.AddonHgexampleTable
}

// TableListInp 获取列表
type TableListInp struct {
	form.PageReq
	Id         int64         `json:"id"          description:""`
	Flag       *gjson.Json   `json:"flag"        description:"标签"`
	Title      string        `json:"title"       description:"标题"`
	Content    string        `json:"content"     description:"内容"`
	Price      []float64     `json:"price"       description:"价格"`
	ActivityAt *gtime.Time   `json:"activityAt"  description:"活动时间"`
	Switch     int           `json:"switch"      description:"开关"`
	Hobby      *gjson.Json   `json:"hobby"       description:"爱好"`
	Status     int           `json:"status"      description:"状态"`
	CreatedAt  []*gtime.Time `json:"createdAt"   description:"创建时间"`
}

type TableListModel struct {
	entity.AddonHgexampleTable
	TableCategoryName        string `json:"TableCategoryName" description:"分类名称"`
	TableCategoryDescription string `json:"TableCategoryDescription" description:"分类描述"`
	TableCategoryRemark      string `json:"TableCategoryRemark" description:"分类备注"`
	SysProvincesTitle        string `json:"sysProvincesTitle" description:""`
}

func (in *TableListInp) Filter(ctx context.Context) (err error) {
	if !in.Flag.IsNil() {
		in.Flag = gjson.New(in.Flag.Var().Ints())
	}
	if !in.Hobby.IsNil() {
		in.Hobby = gjson.New(in.Hobby.Var().Ints())
	}
	return
}

type TableExportModel struct {
	Id         int64       `json:"id"          description:""`
	CategoryId int64       `json:"categoryId"  description:"分类ID"`
	Flag       *gjson.Json `json:"flag"        description:"标签"`
	Title      string      `json:"title"       description:"标题"`
	Star       float64     `json:"star"        description:"推荐星"`
	Price      float64     `json:"price"       description:"价格"`
	Views      int64       `json:"views"       description:"浏览次数"`
	ActivityAt *gtime.Time `json:"activityAt"  description:"活动时间"`
	StartAt    *gtime.Time `json:"startAt"     description:"开启时间"`
	EndAt      *gtime.Time `json:"endAt"       description:"结束时间"`
	Switch     int         `json:"switch"      description:"开关"`
	Sort       int         `json:"sort"        description:"排序"`
	Avatar     string      `json:"avatar"      description:"头像"`
	Sex        int         `json:"sex"         description:"性别"`
	Qq         string      `json:"qq"          description:"qq"`
	Email      string      `json:"email"       description:"邮箱"`
	Mobile     string      `json:"mobile"      description:"手机号码"`
	Hobby      *gjson.Json `json:"hobby"       description:"爱好"`
	Channel    int         `json:"channel"     description:"渠道"`
	Pid        int64       `json:"pid"         description:"上级ID"`
	Level      int         `json:"level"       description:"树等级"`
	Tree       string      `json:"tree"        description:"关系树"`
	Remark     string      `json:"remark"      description:"备注"`
	Status     int         `json:"status"      description:"状态"`
	CreatedBy  int64       `json:"createdBy"   description:"创建者"`
	UpdatedBy  int64       `json:"updatedBy"   description:"更新者"`
	CreatedAt  *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"   description:"修改时间"`
	DeletedAt  *gtime.Time `json:"deletedAt"   description:"删除时间"`
}

// TableMaxSortInp 最大排序
type TableMaxSortInp struct{}

type TableMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// TableStatusInp 更新状态
type TableStatusInp struct {
	Id     int64 `json:"id" v:"required#表格ID不能为空" dc:"表格ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *TableStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type TableStatusModel struct{}

// TableSwitchInp 更新开关状态
type TableSwitchInp struct {
	form.SwitchReq
	Id int64 `json:"id" v:"required#表格ID不能为空" dc:"表格ID"`
}

type TableSwitchModel struct{}
