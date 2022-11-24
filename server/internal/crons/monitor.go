// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package crons

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
	"hotgo/internal/global"
	"hotgo/internal/model"
	"hotgo/utility/format"
	"runtime"
	"sync"
)

// Monitor 监控
var Monitor = &cMonitor{name: "monitor"}

type cMonitor struct {
	name string
	sync.RWMutex
}

func (c *cMonitor) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cMonitor) Execute(ctx context.Context) {
	c.Lock()
	defer c.Unlock()
	c.NetIO()
	c.loadAvg()
}

func (c *cMonitor) loadAvg() {
	pl, _ := load.Avg()
	counter := model.LoadAvgStats{
		Time:  gtime.Now(),
		Avg:   pl.Load1,
		Ratio: pl.Load1 / (float64(runtime.NumCPU()) * 2) * 100,
	}

	global.MonitorData.LoadAvg = append(global.MonitorData.LoadAvg, &counter)
	if len(global.MonitorData.LoadAvg) > 10 {
		global.MonitorData.LoadAvg = append(global.MonitorData.LoadAvg[:0], global.MonitorData.LoadAvg[(1):]...)
	}
}

func (c *cMonitor) NetIO() {
	var counter model.NetIOCounters
	ni, _ := net.IOCounters(true)
	counter.Time = gtime.Now()
	for _, v := range ni {
		counter.BytesSent += v.BytesSent
		counter.BytesRecv += v.BytesRecv
	}

	if len(global.MonitorData.NetIO) > 0 {
		lastNetIO := global.MonitorData.NetIO[len(global.MonitorData.NetIO)-1]
		sub := counter.Time.Sub(lastNetIO.Time).Seconds()
		counter.Down = format.Round2Float64((float64(counter.BytesRecv - lastNetIO.BytesRecv)) / sub)
		counter.UP = format.Round2Float64((float64(counter.BytesSent - lastNetIO.BytesSent)) / sub)
	}

	global.MonitorData.NetIO = append(global.MonitorData.NetIO, &counter)
	if len(global.MonitorData.NetIO) > 10 {
		global.MonitorData.NetIO = append(global.MonitorData.NetIO[:0], global.MonitorData.NetIO[(1):]...)
	}
}
