// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

// 本地配置.

// LogConfig 日志配置
type LogConfig struct {
	Switch   bool     `json:"switch"`
	Queue    bool     `json:"queue"`
	Module   []string `json:"module"`
	SkipCode []string `json:"skipCode"`
}

// ServeLogConfig 服务日志配置
type ServeLogConfig struct {
	Switch      bool     `json:"switch"`
	Queue       bool     `json:"queue"`
	LevelFormat []string `json:"levelFormat"`
}

// GenerateAppCrudTemplate curd模板
type GenerateAppCrudTemplate struct {
	Group          string `json:"group"`
	IsAddon        bool   `json:"isAddon"`
	MasterPackage  string `json:"masterPackage"`
	TemplatePath   string `json:"templatePath"`
	ApiPath        string `json:"apiPath"`
	InputPath      string `json:"inputPath"`
	ControllerPath string `json:"controllerPath"`
	LogicPath      string `json:"logicPath"`
	RouterPath     string `json:"routerPath"`
	SqlPath        string `json:"sqlPath"`
	WebApiPath     string `json:"webApiPath"`
	WebViewsPath   string `json:"webViewsPath"`
}

// GenerateAppQueueTemplate 消息队列模板
type GenerateAppQueueTemplate struct {
	Group        string `json:"group"`
	TemplatePath string `json:"templatePath"`
}

// GenerateAppTreeTemplate 关系树列表模板
type GenerateAppTreeTemplate struct {
	Group        string `json:"group"`
	TemplatePath string `json:"templatePath"`
}

// GenerateConfig 生成代码配置
type GenerateConfig struct {
	AllowedIPs  []string `json:"allowedIPs"`
	Application struct {
		Crud struct {
			Templates []*GenerateAppCrudTemplate `json:"templates"`
		} `json:"crud"`
		Queue struct {
			Templates []*GenerateAppQueueTemplate `json:"templates"`
		} `json:"queue"`
		Tree struct {
			Templates []*GenerateAppTreeTemplate `json:"templates"`
		} `json:"tree"`
	} `json:"application"`
	Delimiters    []string          `json:"delimiters"`
	DevPath       string            `json:"devPath"`
	DisableTables []string          `json:"disableTables"`
	SelectDbs     []string          `json:"selectDbs"`
	Addon         *BuildAddonConfig `json:"addon"`
}

// BuildAddonConfig 构建插件模块配置
type BuildAddonConfig struct {
	SrcPath      string `json:"srcPath"`
	WebApiPath   string `json:"webApiPath"`
	WebViewsPath string `json:"webViewsPath"`
}

// TCPServerConfig tcp服务器配置
type TCPServerConfig struct {
	Address string `json:"address"`
}

// TCPClientConfig tcp客户端配置
type TCPClientConfig struct {
	Cron *TCPClientConnConfig `json:"cron"`
	Auth *TCPClientConnConfig `json:"auth"`
}

// TCPClientConnConfig tcp客户端认证
type TCPClientConnConfig struct {
	Group     string `json:"group"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	AppId     string `json:"appId"`
	SecretKey string `json:"secretKey"`
}

// TCPConfig tcp服务器配置
type TCPConfig struct {
	Server *TCPServerConfig `json:"server"`
	Client *TCPClientConfig `json:"client"`
}

// TokenConfig 登录令牌配置
type TokenConfig struct {
	SecretKey       string `json:"secretKey"`
	Expires         int64  `json:"expires"`
	AutoRefresh     bool   `json:"autoRefresh"`
	RefreshInterval int64  `json:"refreshInterval"`
	MaxRefreshTimes int64  `json:"maxRefreshTimes"`
	MultiLogin      bool   `json:"multiLogin"`
}
