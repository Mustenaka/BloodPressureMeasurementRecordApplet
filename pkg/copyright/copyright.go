package copyright

import (
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/log"
	"fmt"
)

// fmt 直接控制台输出软件版本、版权等信息
func FmtPrintCopyright(conf config.BasicinfoConfig) {
	if conf == (config.BasicinfoConfig{}) {
		panic("GlobalConfig Basicinfo config is empty.")
	}
	fmt.Println("软件名称: ", conf.AppName)
	fmt.Println("软件作者: ", conf.Author)
	fmt.Println("软件开发公司: ", conf.AppCompany)
	fmt.Println("软件版本: ", conf.Version)
	fmt.Println("软件版权: ", conf.Copyright)
}

// 日志模块 输出软件版本、版权等信息
func LogPrintCopyright(conf config.BasicinfoConfig) {
	if conf == (config.BasicinfoConfig{}) {
		panic("GlobalConfig Basicinfo config is empty.")
	}
	log.Info("版权信息", log.WithPair("软件名称", conf.AppName))
	log.Info("版权信息", log.WithPair("软件作者", conf.Author))
	log.Info("版权信息", log.WithPair("软件开发公司", conf.AppCompany))
	log.Info("版权信息", log.WithPair("软件版本", conf.Version))
	log.Info("版权信息", log.WithPair("软件版权", conf.Copyright))
}
