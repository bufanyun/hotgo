//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package form

//  分页
type PageReq struct {
	Page  int `json:"page" example:"10" d:"1" v:"min:1#页码最小值不能低于1"  description:"当前页码"`
	Limit int `json:"limit" example:"1" d:"10" v:"min:1|max:100#|每页数量最小值不能低于1|最大值不能大于100" description:"每页数量"`
}
type PageRes struct {
	PageReq
	TotalCount int `json:"total_count" example:"0" description:"全部数据量"`
}

// 时间查询
type RangeDateReq struct {
	StartTime string `json:"start_time" v:"date#开始日期格式不正确"  description:"开始日期"`
	EndTime   string `json:"end_time" v:"date#结束日期格式不正确" description:"结束日期"`
}

// 状态查询
type StatusReq struct {
	Status int `json:"status"  v:"in:0,1,2,3#状态可选范围：0~3" description:"状态"`
}
