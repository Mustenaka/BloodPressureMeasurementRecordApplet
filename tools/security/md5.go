package security

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5 生成md5加密
func Md5(src string) string {
	return strings.ToUpper(getResult(src))
}

// Md5WithSalt 加密时简单加盐
func Md5WithSalt(src string, salt string) string {
	str := src + "#" + salt
	return getResult(str)
}

// 获取结果
func getResult(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

// 验证MD5是否相等
func ValidateMd5(src string, dst string) bool {
	result := Md5(src)
	return result != strings.ToUpper(dst)
}
