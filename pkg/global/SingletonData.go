package global

import (
	"github.com/widuu/goini"
)

// 单例数据
type SingletonData struct {
	config          goini.Config    // Config - ini的配置文件数据读取
	Basicinfo       Basicinfo       // 软件基本信息
	ServerConf      ServerConf      // 服务器信息
	DatabaseConf    DatabaseConf    // 数据库配置信息
	DBpoolConf      DBpoolConf      // 数据库连接池配置信息
	LoggerConf      LoggerConf      // 日志模块配置信息
	SupureAdminConf SupureAdminConf // 超级用户（管理员信息）
}

type Basicinfo struct {
	AppName    string
	Author     string
	AppCompany string
	Version    string
	Copyright  string
}

type ServerConf struct {
	Mode      string
	Port      int
	Url       string
	MaxPing   int
	JwtSecret string
}

type DatabaseConf struct {
	Username  string
	Password  string
	Hostname  string
	Function  string
	Port      string
	Dbname    string
	Charset   string
	ParseTime string
	Loc       string
}

type DBpoolConf struct {
	MaxIdleConns int
	MaxOpenCoons int
}

type LoggerConf struct {
	Level      string
	FileName   string
	TimeFormat string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   string
	LocalTime  string
	Console    string
}

type SupureAdminConf struct {
	Username string
	Password string
}
