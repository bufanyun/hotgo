// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminin

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

// TestEditInp 修改/新增
type TestEditInp struct {
	entity.Test
}

type TestEditModel struct{}

func (in *TestEditInp) Filter(ctx context.Context) (err error) {
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

// TestDeleteInp 删除类型
type TestDeleteInp struct {
	Id interface{} `json:"id" v:"required#测试ID不能为空" dc:"测试ID"`
}

type TestDeleteModel struct{}

// TestViewInp 获取信息
type TestViewInp struct {
	Id int64 `json:"id" v:"required#测试ID不能为空" dc:"测试ID"`
}

type TestViewModel struct {
	entity.Test
}

// TestListInp 获取列表
type TestListInp struct {
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

type TestListModel struct {
	entity.Test
	TestCategoryName        string `json:"testCategoryName" description:"分类名称"`
	TestCategoryDescription string `json:"testCategoryDescription" description:"分类描述"`
	TestCategoryRemark      string `json:"testCategoryRemark" description:"分类备注"`
	SysProvincesTitle       string `json:"sysProvincesTitle" description:""`
}

func (in *TestListInp) Filter(ctx context.Context) (err error) {
	if !in.Flag.IsNil() {
		in.Flag = gjson.New(in.Flag.Var().Ints())
	}
	if !in.Hobby.IsNil() {
		in.Hobby = gjson.New(in.Hobby.Var().Ints())
	}
	return
}

type TestExportModel struct {
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

// TestMaxSortInp 最大排序
type TestMaxSortInp struct{}

type TestMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// TestStatusInp 更新状态
type TestStatusInp struct {
	Id     int64 `json:"id" v:"required#测试ID不能为空" dc:"测试ID"`
	Status int   `json:"status" dc:"状态"`
}

type TestStatusModel struct{}

// TestSwitchInp 更新开关状态
type TestSwitchInp struct {
	form.SwitchReq
	Id int64 `json:"id" v:"required#测试ID不能为空" dc:"测试ID"`
}

type TestSwitchModel struct{}
