package conf

//应用端口
type AppConfig struct {
	Port string
}

//数据库配置信息
type MySqlConfig struct {
	Username string
	Password string
	FdnUrl string
	GtyUrl string
}

type LoggerConfig struct {
	LoggerPath string
	LoggerLevel string
}

type ProjectConfig struct {
	Mysql MySqlConfig
	LoggerInfo LoggerConfig
	AppConfig AppConfig
}
