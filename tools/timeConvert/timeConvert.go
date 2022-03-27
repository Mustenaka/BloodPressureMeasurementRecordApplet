package timeconvert

import (
	"BloodPressure/pkg/constant"
	ti "time"
)

type EnumTimeType int

const (
	_ EnumTimeType = iota
	datetime
	time
	date
)

// 获取当前时间字符串格式
func NowTimeString() string {
	return TimeConvertString(ti.Now(), time)
}

// 获取当前日期时间字符串格式
func NowDateTimeString() string {
	return TimeConvertString(ti.Now(), datetime)
}

// 获取当前日期字符串格式
func NowDateString() string {
	return TimeConvertString(ti.Now(), date)
}

// 传入date，datetime，time返回当前时间参数
func TimeConvertString(t ti.Time, enumTimeType EnumTimeType) string {
	return t.Format(checkTimeType(enumTimeType))
}

func checkTimeType(enumTimeType EnumTimeType) string {
	var typeResult string
	// 根据当前类型获取字符常量
	if enumTimeType == datetime {
		typeResult = constant.DateTimeLayout
	} else if enumTimeType == time {
		typeResult = constant.TimeLayout
	} else if enumTimeType == date {
		typeResult = constant.DateLayout
	} else {
		return ""
	}
	return typeResult
}
