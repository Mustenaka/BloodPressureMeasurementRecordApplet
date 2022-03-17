package global

import (
	"strconv"
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

// 获取config的值 并转换为int
func (*SingletonData) GetConfigValueInt(selection string, key string) int {
	result, _ := strconv.Atoi(singletonData.config.GetValue(selection, key))
	return result
}

// 获取config的值 并转换为bool
func (*SingletonData) GetConfigValueBool(selection string, key string) bool {
	result, _ := strconv.ParseBool(singletonData.config.GetValue(selection, key))
	return result
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
