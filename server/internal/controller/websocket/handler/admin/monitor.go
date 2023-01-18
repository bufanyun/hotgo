// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package admin

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"hotgo/internal/consts"
	"hotgo/internal/global"
	"hotgo/internal/model"
	"hotgo/internal/websocket"
	"hotgo/utility/file"
	"hotgo/utility/format"
	"os"
	"runtime"
	"strconv"
	"time"
)

var (
	Monitor = cMonitor{}
)

type cMonitor struct{}

type MonitorHeadExtra struct {
	Data  interface{} `json:"data"`
	Data1 string      `json:"data1,omitempty"`
}
type MonitorHead struct {
	Title       string           `json:"title"`
	Data        string           `json:"data"`
	BottomTitle string           `json:"bottomTitle"`
	TotalSum    string           `json:"totalSum"`
	IconClass   string           `json:"iconClass"`
	Extra       MonitorHeadExtra `json:"extra"`
}

// RunInfo 运行信息
func (c *cMonitor) RunInfo(client *websocket.Client, req *websocket.WRequest) {
	var (
		data     = g.Map{}
		mHost, _ = host.Info()
		pwd, _   = os.Getwd()
		gm       runtime.MemStats
	)

	runtime.ReadMemStats(&gm)

	data = g.Map{
		// 服务器信息
		"hostname":    mHost.Hostname,
		"os":          mHost.OS,
		"arch":        mHost.KernelArch,
		"intranet_ip": global.MonitorData.IntranetIP,
		"public_ip":   global.MonitorData.PublicIP,

		// GO运行信息
		"goName":    "Golang",
		"version":   runtime.Version(),
		"startTime": global.MonitorData.STartTime,
		"runTime":   gtime.Now().Timestamp() - global.MonitorData.STartTime.Timestamp(),
		"rootPath":  runtime.GOROOT(),
		"pwd":       pwd,
		"goroutine": runtime.NumGoroutine(),
		"goMem":     format.FileSize(int64(gm.Sys)),
		"goSize":    file.DirSize(pwd),
	}

	isDemo := g.Cfg().MustGet(client.Context(), "hotgo.isDemo", false)
	if isDemo.Bool() {
		data["rootPath"] = consts.DemoTips
		data["pwd"] = consts.DemoTips
		data["intranet_ip"] = consts.DemoTips
		data["public_ip"] = consts.DemoTips
	}

	websocket.SendSuccess(client, req.Event, data)
}

// Trends 实时数据
func (c *cMonitor) Trends(client *websocket.Client, req *websocket.WRequest) {

	type NetC struct {
		Time      *gtime.Time `json:"time"`
		BytesSent string      `json:"bytesSent"` // number of bytes sent
		BytesRecv string      `json:"bytesRecv"` // number of bytes received
		Down      float64     `json:"down"`
		UP        float64     `json:"up"`
	}

	var (
		mCpu, _      = cpu.Info()
		mCpuUsed     float64
		mMem, _      = mem.VirtualMemory()
		mMemUsed     float64
		mDisk, _     = disk.Usage("/")
		mProcess, _  = process.Pids()
		mLoadAvg     *model.LoadAvgStats
		data         = g.Map{}
		monitorHeads []MonitorHead
		nets         []NetC
	)

	// cpu使用率
	cu, err := cpu.Percent(time.Second, false)
	if err == nil {
		mCpuUsed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cu[0]), 64)
	}

	// 内存使用率
	mMemUsed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", mMem.UsedPercent), 64)

	// 负载
	if len(global.MonitorData.LoadAvg) > 0 {
		mLoadAvg = global.MonitorData.LoadAvg[len(global.MonitorData.LoadAvg)-1]
	}

	monitorHeads = append(monitorHeads, MonitorHead{
		Title:       "CPU",
		Data:        "使用率 " + format.Round2String(mCpuUsed) + "%",
		BottomTitle: "CPU数量",
		TotalSum:    gconv.String(runtime.NumCPU()) + "核心",
		IconClass:   "HardwareChip",
		Extra: MonitorHeadExtra{
			Data:  mCpu[0].VendorID,
			Data1: mCpu[0].ModelName,
		}},
		MonitorHead{
			Title:       "内存",
			Data:        "使用率 " + format.Round2String(mMemUsed) + "%",
			BottomTitle: "总内存",
			TotalSum:    format.FileSize(int64(mMem.Total)),
			IconClass:   "AppsSharp",
			Extra: MonitorHeadExtra{
				Data:  format.FileSize(int64(mMem.Used)),
				Data1: format.FileSize(int64(mMem.Total - mMem.Used)),
			}},
		MonitorHead{
			Title:       "磁盘",
			Data:        "已用 " + format.FileSize(int64(mDisk.Used)),
			BottomTitle: "总容量",
			TotalSum:    format.FileSize(int64(mDisk.Total)),
			IconClass:   "PieChart",
			Extra: MonitorHeadExtra{
				Data: format.Round2String(mDisk.UsedPercent),
			}},
		MonitorHead{
			Title:       "负载",
			Data:        format.Round2String(mLoadAvg.Ratio) + "%",
			BottomTitle: "总进程数",
			TotalSum:    gconv.String(len(mProcess)) + "个",
			IconClass:   "Analytics",
		})

	for _, v := range global.MonitorData.NetIO {
		nets = append(nets, NetC{
			Time:      v.Time,
			BytesSent: format.FileSize(int64(v.BytesSent)), // 转换为最大整数单位
			BytesRecv: format.FileSize(int64(v.BytesRecv)),
			Down:      format.Round2Float64(v.Down / 1024), // 转换为kb
			UP:        format.Round2Float64(v.UP / 1024),
		})
	}

	data = g.Map{
		"head": monitorHeads,
		"load": global.MonitorData.LoadAvg,
		"net":  nets,
	}

	websocket.SendSuccess(client, req.Event, data)
}
