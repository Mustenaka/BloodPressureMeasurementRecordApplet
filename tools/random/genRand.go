package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 生成随机数字Returns an int >= min, < max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// 生成随机大写字符串
func RandomUpperString(len int) string {
	// len 不能低于 0
	if len <= 0 {
		return ""
	}
	// 生成随机字符串-大写
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

// 生成随机小写字符串
func RandomLowerString(len int) string {
	// len 不能低于 0
	if len <= 0 {
		return ""
	}
	// 生成随机字符串-大写
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomInt(97, 122))
	}
	return string(bytes)
}
