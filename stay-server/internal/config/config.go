package config

type AppConfig struct {
	MysqlConfig MysqlConfig `yaml:"mysql_config" json:"mysql_config"`
	WxApp       WxApp       `yaml:"wx_app"`
}

type MysqlConfig struct {
	Protocol string `yaml:"protocol" json:"protocol"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
}

type WxApp struct {
	AppId     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}
