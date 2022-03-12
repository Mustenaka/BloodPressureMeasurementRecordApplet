package main

import (
	"sync"

	"github.com/widuu/goini"
)

type SingleData struct {
	config goini.Config
}

var singleData *SingleData
var mutx sync.Mutex

func GetInstance4() *SingleData {
	if singleData == nil {
		mutx.Lock()
		if singleData == nil {
			singleData = &SingleData{}
			mutx.Unlock()
			return singleData
		}

	}
	mutx.Unlock()
	return singleData
}
