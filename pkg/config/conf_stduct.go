package config

// Config is application global config
type Config struct {
	BasicinfoConfig BasicinfoConfig `mapstructure:"basicinfo"` // 项目基本配置信息
	ServerConfig    ServerConfig    `mapstructure:"server"`    // 服务信息
	DBConfig        DBConfig        `mapstructure:"database"`  // 数据库信息
	RedisConfig     RedisConfig     `mapstructure:"redis"`     // redis
	LogConfig       LogConfig       `mapstructure:"logconfig"` // uber zap
}

// Basic information is used to version control.
type BasicinfoConfig struct {
	AppName    string `mapstructure:"appName"` // 应用名称
	Author     string `mapstructure:"author"`
	AppCompany string `mapstructure:"appCompany"`
	Version    string `mapstructure:"version"`
	Copyright  string `mapstructure:"copyright"`
}

// Basic information is used to version control.
type ServerConfig struct {
	Mode         string `mapstructure:"mode"`           // gin启动模式
	Port         string `mapstructure:"port"`           // 启动端口
	Url          string `mapstructure:"url"`            // 应用地址,用于自检 eg. http://127.0.0.1
	MaxPingCount int    `mapstructure:"max-ping-count"` // 最大自检次数，用户健康检查
	JwtSecret    string `mapstructure:"jwt-secret"`
}

// DBConfig is used to configure mysql database
type DBConfig struct {
	Dbname          string `mapstructure:"dbname"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	MaximumPoolSize int    `mapstructure:"maximum-pool-size"`
	MaximumIdleSize int    `mapstructure:"maximum-idle-size"`
	LogMode         bool   `mapstructure:"log-mode"`
}

// RedisConfig is used to configure redis
type RedisConfig struct {
	Addr         string `mapstructure:"address"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool-size"`
	MinIdleConns int    `mapstructure:"min-idle-conns"`
	IdleTimeout  int    `mapstructure:"idle-timeout"`
}

// LogConfig is used to configure uber zap
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file-name"`
	TimeFormat string `mapstructure:"time-format"`
	MaxSize    int    `mapstructure:"max-size"`
	MaxBackups int    `mapstructure:"max-backups"`
	MaxAge     int    `mapstructure:"max-age"`
	Compress   bool   `mapstructure:"compress"`
	LocalTime  bool   `mapstructure:"local-time"`
	Console    bool   `mapstructure:"console"`
}
