package routers

type DbConfig struct {
	MysqlConfig MysqlConfig `yaml:"mysql_config" json:"mysql_config"`
}

type MysqlConfig struct {
	Protocol string `yaml:"protocol" json:"protocol"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
}
