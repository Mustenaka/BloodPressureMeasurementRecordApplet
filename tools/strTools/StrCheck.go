package strtools

import "strings"

// 如果src为空则返回dst，如果不是则返回src本身
func UpdateNotNullStirng(src, dst string) string {
	if src == "" {
		return dst
	}
	return src
}

// 通过传入的DateTime切割成两个date和time字符串
func SplitDateTime(dateTimeString string) (date, time string) {
	context := strings.Fields(dateTimeString)
	date = context[0]
	time = context[1]
	return date, time
}
