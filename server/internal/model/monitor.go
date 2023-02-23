// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type MonitorData struct {
	// STartTime 启动时间
	STartTime *gtime.Time
	// 内网IP
	IntranetIP string
	// 公网IP
	PublicIP string
	// NetIO 网络流量统计
	NetIO []*NetIOCounters
	// LoadAvg 负载统计
	LoadAvg []*LoadAvgStats
}

type NetIOCounters struct {
	Time      *gtime.Time `json:"time"`
	BytesSent uint64      `json:"bytesSent"` // number of bytes sent
	BytesRecv uint64      `json:"bytesRecv"` // number of bytes received
	Down      float64     `json:"down"`
	UP        float64     `json:"up"`
}

type LoadAvgStats struct {
	Time  *gtime.Time `json:"time"`
	Avg   float64     `json:"avg"`
	Ratio float64     `json:"ratio"`
}
