package model

// BasicConfig 基础配置
type BasicConfig struct {
	CaptchaSwitch  int    `json:"basicCaptchaSwitch"`
	CloseText      string `json:"basicCloseText"`
	Copyright      string `json:"basicCopyright"`
	IcpCode        string `json:"basicIcpCode"`
	Logo           string `json:"basicLogo"`
	Name           string `json:"basicName"`
	Domain         string `json:"basicDomain"`
	RegisterSwitch int    `json:"basicRegisterSwitch"`
	SystemOpen     bool   `json:"basicSystemOpen"`
}

// EmailTemplate 邮件模板
type EmailTemplate struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// EmailConfig 邮箱配置
type EmailConfig struct {
	User         string           `json:"smtpUser"`
	Password     string           `json:"smtpPass"`
	Addr         string           `json:"smtpAddr"`
	Host         string           `json:"smtpHost"`
	Port         int64            `json:"smtpPort"`
	SendName     string           `json:"smtpSendName"`
	AdminMailbox string           `json:"smtpAdminMailbox"`
	MinInterval  int              `json:"smtpMinInterval"`
	MaxIpLimit   int              `json:"smtpMaxIpLimit"`
	CodeExpire   int              `json:"smtpCodeExpire"`
	Template     []*EmailTemplate `json:"smtpTemplate"`
}

// CashConfig 提现配置
type CashConfig struct {
	Switch      bool    `json:"cashSwitch"`
	MinFee      float64 `json:"cashMinFee"`
	MinFeeRatio float64 `json:"cashMinFeeRatio"`
	MinMoney    float64 `json:"cashMinMoney"`
	Tips        string  `json:"cashTips"`
}

// UploadConfig 上传配置
type UploadConfig struct {
	Drive string `json:"uploadDrive"`
	// 基本配置
	FileSize  int64  `json:"uploadFileSize"`
	FileType  string `json:"uploadFileType"`
	ImageSize int64  `json:"uploadImageSize"`
	ImageType string `json:"uploadImageType"`
	// 本地存储配置
	LocalPath string `json:"uploadLocalPath"`
	// UCloud对象存储配置
	UCloudBucketHost string `json:"uploadUCloudBucketHost"`
	UCloudBucketName string `json:"uploadUCloudBucketName"`
	UCloudEndpoint   string `json:"uploadUCloudEndpoint"`
	UCloudFileHost   string `json:"uploadUCloudFileHost"`
	UCloudPath       string `json:"uploadUCloudPath"`
	UCloudPrivateKey string `json:"uploadUCloudPrivateKey"`
	UCloudPublicKey  string `json:"uploadUCloudPublicKey"`
}

// GeoConfig 地理配置
type GeoConfig struct {
	GeoAmapWebKey string `json:"geoAmapWebKey"`
}

// SmsTemplate 短信模板
type SmsTemplate struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SmsConfig 短信配置
type SmsConfig struct {
	// 基础
	SmsDrive       string `json:"smsDrive"`
	SmsMinInterval int    `json:"smsMinInterval"`
	SmsMaxIpLimit  int    `json:"smsMaxIpLimit"`
	SmsCodeExpire  int    `json:"smsCodeExpire"`
	// 阿里云
	SmsAliyunAccessKeyID     string         `json:"smsAliyunAccessKeyID"`
	SmsAliyunAccessKeySecret string         `json:"smsAliyunAccessKeySecret"`
	SmsAliyunSign            string         `json:"smsAliyunSign"`
	SmsAliyunTemplate        []*SmsTemplate `json:"smsAliyunTemplate"`
}

///////////// 以下是本地配置

// SSLConfig https配置
type SSLConfig struct {
	Switch  bool   `json:"switch"`
	CrtPath string `json:"crtPath"`
	KeyPath string `json:"keyPath"`
}

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
	Delimiters    []string `json:"delimiters"`
	DevPath       string   `json:"devPath"`
	DisableTables []string `json:"disableTables"`
	SelectDbs     []string `json:"selectDbs"`
}
