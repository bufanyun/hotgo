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
	WsAddr         string `json:"basicWsAddr"`
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
	// 通用配置
	Drive     string `json:"uploadDrive"`
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
	// 腾讯云cos配置
	CosSecretId  string `json:"uploadCosSecretId"`
	CosSecretKey string `json:"uploadCosSecretKey"`
	CosBucketURL string `json:"uploadCosBucketURL"`
	CosPath      string `json:"uploadCosPath"`
	// 阿里云oss配置
	OssSecretId  string `json:"uploadOssSecretId"`
	OssSecretKey string `json:"uploadOssSecretKey"`
	OssEndpoint  string `json:"uploadOssEndpoint"`
	OssBucketURL string `json:"uploadOssBucketURL"`
	OssPath      string `json:"uploadOssPath"`
	OssBucket    string `json:"uploadOssBucket"`
	// 七牛云对象存储配置
	QiNiuAccessKey string `json:"uploadQiNiuAccessKey"`
	QiNiuSecretKey string `json:"uploadQiNiuSecretKey"`
	QiNiuDomain    string `json:"uploadQiNiuDomain"`
	QiNiuPath      string `json:"uploadQiNiuPath"`
	QiNiuBucket    string `json:"uploadQiNiuBucket"`
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
	AliYunAccessKeyID     string         `json:"smsAliYunAccessKeyID"`
	AliYunAccessKeySecret string         `json:"smsAliYunAccessKeySecret"`
	AliYunSign            string         `json:"smsAliYunSign"`
	AliYunTemplate        []*SmsTemplate `json:"smsAliYunTemplate"`
	// 腾讯云
	TencentSecretId  string         `json:"smsTencentSecretId"`
	TencentSecretKey string         `json:"smsTencentSecretKey"`
	TencentEndpoint  string         `json:"smsTencentEndpoint"`
	TencentRegion    string         `json:"smsTencentRegion"`
	TencentAppId     string         `json:"smsTencentAppId"`
	TencentSign      string         `json:"smsTencentSign"`
	TencentTemplate  []*SmsTemplate `json:"smsTencentTemplate"`
}

type PayConfig struct {
	Debug bool `json:"payDebug"`
	// 支付宝
	AliPayAppId             string `json:"payAliPayAppId"`
	AliPayPrivateKey        string `json:"payAliPayPrivateKey"`
	AliPayAppCertPublicKey  string `json:"payAliPayAppCertPublicKey"`
	AliPayRootCert          string `json:"payAliPayRootCert"`
	AliPayCertPublicKeyRSA2 string `json:"payAliPayCertPublicKeyRSA2"`
	// 微信支付
	WxPayAppId      string `json:"payWxPayAppId"`
	WxPayMchId      string `json:"payWxPayMchId"`
	WxPaySerialNo   string `json:"payWxPaySerialNo"`
	WxPayAPIv3Key   string `json:"payWxPayAPIv3Key"`
	WxPayPrivateKey string `json:"payWxPayPrivateKey"`
	// QQ支付
	QQPayAppId  string `json:"payQQPayAppId"`
	QQPayMchId  string `json:"payQQPayMchId"`
	QQPayApiKey string `json:"payQQPayApiKey"`
}

// WechatOfficialAccountConfig 微信公众号配置
type WechatOfficialAccountConfig struct {
	OfficialAppID          string `json:"officialAccountAppId"`          // appid
	OfficialAppSecret      string `json:"officialAccountAppSecret"`      // app secret
	OfficialToken          string `json:"officialAccountToken"`          // token
	OfficialEncodingAESKey string `json:"officialAccountEncodingAESKey"` // EncodingAESKey
}

// WechatOpenPlatformConfig 微信开放平台配置
type WechatOpenPlatformConfig struct {
	OpenAppID          string `json:"openPlatformAppId"`          // appid
	OpenAppSecret      string `json:"openPlatformAppSecret"`      // app secret
	OpenToken          string `json:"openPlatformToken"`          // token
	OpenEncodingAESKey string `json:"openPlatformEncodingAESKey"` // EncodingAESKey
}

// WechatConfig 微信配置
type WechatConfig struct {
	*WechatOfficialAccountConfig
	*WechatOpenPlatformConfig
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
	TemplatePath string `json:"templatePath"`
	WebApiPath   string `json:"webApiPath"`
	WebViewsPath string `json:"webViewsPath"`
}

// CacheConfig 缓存配置
type CacheConfig struct {
	Adapter string `json:"adapter"`
	FileDir string `json:"fileDir"`
}

type TCPServerConfig struct {
	Address string `json:"address"`
}

type TCPClientConfig struct {
	Cron *TCPClientConnConfig `json:"cron"`
	Auth *TCPClientConnConfig `json:"auth"`
}

type TCPClientConnConfig struct {
	Group     string `json:"group"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	AppId     string `json:"appId"`
	SecretKey string `json:"secretKey"`
}

type TCPConfig struct {
	Server *TCPServerConfig `json:"server"`
	Client *TCPClientConfig `json:"client"`
}
