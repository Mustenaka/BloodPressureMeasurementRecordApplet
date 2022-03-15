package log

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
