// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"hotgo/internal/consts"
	"hotgo/internal/model"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/file"
	"hotgo/utility/format"
	"os"
	"runtime"
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
		meta     = service.AdminMonitor().GetMeta(client.Context())
		mHost, _ = host.Info()
		pwd, _   = os.Getwd()
		gm       runtime.MemStats
	)

	runtime.ReadMemStats(&gm)

	data := g.Map{
		// 服务器信息
		"hostname":    mHost.Hostname,
		"os":          mHost.OS,
		"arch":        mHost.KernelArch,
		"intranet_ip": meta.IntranetIP,
		"public_ip":   meta.PublicIP,

		// 运行信息
		"version":   runtime.Version(), // GO 版本
		"hgVersion": consts.VersionApp, // HG 版本
		"startTime": gtime.New(meta.STartTime),
		"runTime":   gtime.Now().Timestamp() - meta.STartTime,
		"rootPath":  runtime.GOROOT(),
		"pwd":       pwd,
		"goroutine": runtime.NumGoroutine(),
		"goMem":     format.FileSize(int64(gm.Sys)),
		"goSize":    file.DirSize(pwd),
	}

	isDemo := g.Cfg().MustGet(client.Context(), "hotgo.isDemo", false).Bool()
	if isDemo {
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
		mCpu, cpuErr         = cpu.Info()
		mCpuUsed             float64
		mMem, memErr         = mem.VirtualMemory()
		mMemUsed             float64
		mDisk, diskErr       = disk.Usage("/")
		mProcess, processErr = process.Pids()
		mLoadAvg             = new(model.LoadAvgStats)
		monitorHeads         []MonitorHead
		nets                 []NetC
		meta                 = service.AdminMonitor().GetMeta(client.Context())
	)

	if cpuErr != nil {
		g.Log().Infof(client.Context(), "read CPU info fail:%+v", cpuErr)
		mCpu = []cpu.InfoStat{{VendorID: "", ModelName: ""}}
	}

	if memErr != nil {
		g.Log().Infof(client.Context(), "read mem info fail:%+v", memErr)
		mMem = new(mem.VirtualMemoryStat)
	}

	if diskErr != nil {
		g.Log().Infof(client.Context(), "read disk info fail:%+v", diskErr)
		mDisk = new(disk.UsageStat)
	}

	if processErr != nil {
		g.Log().Infof(client.Context(), "read process.Pids fail:%+v", processErr)
	}

	// cpu使用率
	cu, err := cpu.Percent(time.Second, false)
	if err == nil {
		mCpuUsed = gconv.Float64(fmt.Sprintf("%.2f", cu[0]))
	}

	// 内存使用率
	mMemUsed = gconv.Float64(fmt.Sprintf("%.2f", mMem.UsedPercent))

	// 负载
	if len(meta.LoadAvg) > 0 {
		mLoadAvg = meta.LoadAvg[len(meta.LoadAvg)-1]
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

	for _, v := range meta.NetIO {
		nets = append(nets, NetC{
			Time:      v.Time,
			BytesSent: format.FileSize(int64(v.BytesSent)), // 转换为最大整数单位
			BytesRecv: format.FileSize(int64(v.BytesRecv)),
			Down:      format.Round2Float64(v.Down / 1024), // 转换为kb
			UP:        format.Round2Float64(v.UP / 1024),
		})
	}

	websocket.SendSuccess(client, req.Event, g.Map{
		"head": monitorHeads,
		"load": meta.LoadAvg,
		"net":  nets,
	})
}
