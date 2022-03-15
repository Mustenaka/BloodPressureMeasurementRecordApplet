package copyright

import (
	"BloodPressure/pkg/global"
	"BloodPressure/pkg/log"
	"fmt"
)

var (
	AppName    = global.GetInstance().GetConfigValue("basicinfo", "appName")
	Author     = global.GetInstance().GetConfigValue("basicinfo", "author")
	AppCompany = global.GetInstance().GetConfigValue("basicinfo", "appCompany")
	Version    = global.GetInstance().GetConfigValue("basicinfo", "version")
	Copyright  = global.GetInstance().GetConfigValue("basicinfo", "copyright")
)

// fmt 直接控制台输出软件版本、版权等信息
func FmtPrintCopyright() {
	fmt.Println("软件名称: ", AppName)
	fmt.Println("软件作者: ", Author)
	fmt.Println("软件开发公司: ", AppCompany)
	fmt.Println("软件版本: ", Version)
	fmt.Println("软件版权: ", Copyright)
}

// 日志模块 输出软件版本、版权等信息
func LogPrintCopyright() {
	log.Info("版权信息", log.WithPair("软件名称", AppName))
	log.Info("版权信息", log.WithPair("软件作者", Author))
	log.Info("版权信息", log.WithPair("软件开发公司", AppCompany))
	log.Info("版权信息", log.WithPair("软件版本", Version))
	log.Info("版权信息", log.WithPair("软件版权", Copyright))
}
