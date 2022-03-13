package tools

import (
	"math/rand"
	"time"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// 生成随机字符串
func RandomString(len int) string {
	rand.Seed(time.Now().UnixNano())
	// len 不能低于 0
	if len <= 0 {
		return ""
	}
	// 生成随机字符串
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
