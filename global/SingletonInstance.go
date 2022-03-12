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
func GetInstance(iniFileName string) *SingletonData {
	if singletonData == nil {
		mutx.Lock()
		if singletonData == nil {
			singletonData = &SingletonData{
				Config: *goini.SetConfig(iniFileName),
			}
			mutx.Unlock()
			return singletonData
		}
	}
	mutx.Unlock()
	return singletonData
}
