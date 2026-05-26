// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

// 后台配置.

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
	// minio配置
	MinioAccessKey string `json:"uploadMinioAccessKey"`
	MinioSecretKey string `json:"uploadMinioSecretKey"`
	MinioEndpoint  string `json:"uploadMinioEndpoint"`
	MinioUseSSL    int    `json:"uploadMinioUseSSL"`
	MinioPath      string `json:"uploadMinioPath"`
	MinioBucket    string `json:"uploadMinioBucket"`
	MinioDomain    string `json:"uploadMinioDomain"`
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

// LoginConfig 登录配置
type LoginConfig struct {
	RegisterSwitch int     `json:"loginRegisterSwitch"`
	CaptchaSwitch  int     `json:"loginCaptchaSwitch"`
	CaptchaType    int     `json:"loginCaptchaType"`
	Avatar         string  `json:"loginAvatar"`
	RoleId         int64   `json:"loginRoleId"`
	DeptId         int64   `json:"loginDeptId"`
	PostIds        []int64 `json:"loginPostIds"`
	Protocol       string  `json:"loginProtocol"`
	Policy         string  `json:"loginPolicy"`
	AutoOpenId     int     `json:"loginAutoOpenId"`
	ForceInvite    int     `json:"loginForceInvite"`
}
