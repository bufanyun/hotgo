// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package consts

// 生成代码类型
const (
	GenCodesTypeCurd  = 10 // 增删改查列表
	GenCodesTypeTree  = 11 // 树列表
	GenCodesTypeQueue = 20 // 队列消费者
	GenCodesTypeCron  = 30 // 定时任务
)

var GenCodesTypeNameMap = map[int]string{
	GenCodesTypeCurd:  "增删改查列表",
	GenCodesTypeTree:  "关系树列表(未实现)",
	GenCodesTypeQueue: "队列消费者(未实现)",
	GenCodesTypeCron:  "定时任务(未实现)",
}

var GenCodesTypeConfMap = map[int]string{
	GenCodesTypeCurd:  "crud",
	GenCodesTypeTree:  "tree",
	GenCodesTypeQueue: "queue",
	GenCodesTypeCron:  "cron",
}

// 生成代码状态
const (
	GenCodesStatusOk   = 1 // 生成成功
	GenCodesStatusWait = 2 // 未生成
	GenCodesStatusFail = 3 // 生成失败
)

var GenCodesStatusNameMap = map[int]string{
	GenCodesStatusOk:   "生成成功",
	GenCodesStatusWait: "未生成",
	GenCodesStatusFail: "生成失败",
}

// 生成代码关联表方式
const (
	GenCodesJoinLeft  = 1 // 左关联
	GenCodesJoinRight = 2 // 右关联
	GenCodesJoinInner = 3 // 内关联
)

var GenCodesJoinNameMap = map[int]string{
	GenCodesJoinLeft:  "左关联",
	GenCodesJoinRight: "右关联",
	GenCodesJoinInner: "内关联",
}

var GenCodesJoinLinkMap = map[int]string{
	GenCodesJoinLeft:  "LeftJoin",
	GenCodesJoinRight: "RightJoin",
	GenCodesJoinInner: "InnerJoin",
}

// 生成代码的生成方式
const (
	GenCodesBuildMethCreate = 1 // 创建
	GenCodesBuildMethCover  = 2 // 覆盖
	GenCodesBuildMethSkip   = 3 // 跳过
	GenCodesBuildIgnore     = 4 // 不生成
)

var GenCodesBuildMethNameMap = map[int]string{
	GenCodesBuildMethCreate: "创建文件",
	GenCodesBuildMethCover:  "强制覆盖",
	GenCodesBuildMethSkip:   "已存在跳过",
	GenCodesBuildIgnore:     "不生成",
}

const (
	GenCodesIndexPK  = "PRI" // 主键索引
	GenCodesIndexUNI = "UNI" // 唯一索引
)
