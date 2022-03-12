package main

import (
	"sync"

	"github.com/widuu/goini"
)

type SingletonData struct {
	config goini.Config
}

// 单例对象
var singletonData *SingletonData

// 动态锁
var mutx sync.Mutex

func GetInstance(iniFileName string) *SingletonData {
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
