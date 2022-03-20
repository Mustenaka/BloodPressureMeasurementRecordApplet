package main

import (
	"BloodPressure/pkg/global"
	"BloodPressure/pkg/log"
)

// Test hello
func RunProgram() {
	// // 解析服务器启动参数
	// appOpt := &server.AppOptions{}
	// server.ResolveAppOptions(appOpt)
	// if appOpt.PrintVersion {
	// 	version.PrintVersion()
	// }

	// 加载配置参数
	conf := global.GetInstance()
	// logConfig := log.LogConfig{
	// 	Level:      conf.GetConfigValue("logconfig", "level"),
	// 	FileName:   conf.GetConfigValue("logconfig", "file-name"),
	// 	TimeFormat: constant.TimeLayout,
	// 	MaxSize:    conf.GetConfigValueInt("logconfig", "max-size"),
	// 	MaxBackups: conf.GetConfigValueInt("logconfig", "max-backups"),
	// 	MaxAge:     conf.GetConfigValueInt("logconfig", "max-age"),
	// 	Compress:   conf.GetConfigValueBool("logconfig", "compress"),
	// 	LocalTime:  conf.GetConfigValueBool("logconfig", "local-time"),
	// 	Console:    conf.GetConfigValueBool("logconfig", "console"),
	// }
	// log.InitLogger(&logConfig, "aaaa")
	log.InitLogger(log.InitLoggerWithConfig(), "asd")
	log.Info("basicinfo: ", log.WithPair("AppName", conf.GetConfigValue("basicinfo", "appName")))
	log.Info("basicinfo: ", log.WithPair("Version", conf.GetConfigValue("basicinfo", "version")))
	log.Info("basicinfo: ", log.WithPair("Copyright", conf.GetConfigValue("basicinfo", "copyright")))

	// 加载数据库
	// model.Connect()
}

func main() {
	RunProgram()
}
