package model

// EmailConfig 邮箱配置
type EmailConfig struct {
	User         string `json:"smtpUser"`
	Password     string `json:"smtpPass"`
	Addr         string `json:"smtpAddr"`
	Host         string `json:"smtpHost"`
	Port         int64  `json:"smtpPort"`
	SendName     string `json:"smtpSendName"`
	AdminMailbox string `json:"smtpAdminMailbox"`
}

// CashConfig 提现配置
type CashConfig struct {
	Switch      bool    `json:"cashSwitch"`
	MinFee      float64 `json:"cashMinFee"`
	MinFeeRatio float64 `json:"cashMinFeeRatio"`
	MinMoney    float64 `json:"cashMinMoney"`
	Tips        string  `json:"cashTips"`
}

///////////// 以下是本地配置

// SSLConfig https配置
type SSLConfig struct {
	Switch  bool   `json:"switch"`
	CrtPath string `json:"crtPath"`
	KeyPath string `json:"keyPath"`
}
