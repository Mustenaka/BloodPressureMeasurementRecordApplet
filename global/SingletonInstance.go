package global

import (
	"sync"

	"github.com/widuu/goini"
)

// 单例对象
var singletonData *SingletonData

// 动态锁
var mutx sync.Mutex

// 获取单例模式
func GetInstance() *SingletonData {
	// ini 文件地址
	iniFileName := "./config/config.ini"
	// 双重锁，延迟获取
	if singletonData == nil {
		mutx.Lock()
		if singletonData == nil {
			singletonData = &SingletonData{
				config: *goini.SetConfig(iniFileName),
			}
			mutx.Unlock()
			return singletonData
		}
	}
	mutx.Unlock()
	return singletonData
}

// 获取config的值
func (*SingletonData) GetConfigValue(selection string, key string) string {
	return singletonData.config.GetValue(selection, key)
}

// 设置config的值
func (*SingletonData) SetConfigValue(selection string, key string, value string) bool {
	return singletonData.config.SetValue(selection, key, value)
}

// 删除config的值
func (*SingletonData) DeleteConfigValue(selection string, key string) {
	singletonData.config.DeleteValue(selection, key)
}

// 获取全部的config的值
func (*SingletonData) ReadConfigList() []map[string]map[string]string {
	return singletonData.config.ReadList()
}
