package openid

import (
	"fmt"
	"testing"
)

func TestGetOpenidByCode(t *testing.T) {
	/*
		这个测试意义不是很大，因为code每一次获取使用之后都会更改，而且超过5分钟不使用就会作废
		所以只能通过复制小程序端的code代码，粘贴到code := "xxxx"中，手动测一下
		这样测试的效率不仅缓慢而且没有意义
	*/
	code := ""
	openId, err := GetOpenidByCode(code)
	if err != nil {
		t.Error("get openid failed")
	}
	fmt.Println(openId)
	t.Log(openId)
}
