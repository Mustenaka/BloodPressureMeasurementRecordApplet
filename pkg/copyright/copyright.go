package copyright

// 打印版权信息
// 直接单例调用版权信息，版权信息加载来自于config.ini信息

import (
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/log"
	"fmt"
	"sync"
)

// 版权结构体
type copyright struct {
	AppName    string `mapstructure:"appName"` // 应用名称
	Author     string `mapstructure:"author"`
	AppCompany string `mapstructure:"appCompany"`
	Version    string `mapstructure:"version"`
	Copyright  string `mapstructure:"copyright"`
}

// 单例模式加载版权信息
type Singleton struct {
	*copyright
}

var single *Singleton
var mut sync.Mutex

// 单例加载版权信息(双重锁)
func GetInstance() *Singleton {
	if single == nil {
		mut.Lock()
		defer mut.Unlock()
		if single == nil {
			single = &Singleton{}
		}
	}
	return single
}

// 加载版权信息
func (s *Singleton) LoadCopyright(conf config.BasicinfoConfig) {
	// 空数据排查
	if conf == (config.BasicinfoConfig{}) {
		panic("GlobalConfig Basicinfo config is empty.")
	}

	// 加载数据
	s.copyright = &copyright{
		AppName:    conf.AppName,
		Author:     conf.Author,
		AppCompany: conf.AppCompany,
		Version:    conf.Version,
		Copyright:  conf.Copyright,
	}
}

func (s *Singleton) GetCopyright() copyright {
	return *s.copyright
}

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
