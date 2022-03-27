package log

import (
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/constant"
	"BloodPressure/tools/uuid"
	"context"
	"testing"
)

func init() {
	defer Sync()
	c := config.Config{
		LogConfig: config.LogConfig{
			Level:      "debug",
			FileName:   "test.log",
			TimeFormat: constant.DateTimeLayout,
			MaxSize:    1,
			MaxBackups: 5,
			MaxAge:     2,
			Compress:   false,
			LocalTime:  true,
			Console:    true,
		},
		BasicinfoConfig: config.BasicinfoConfig{
			AppName: "zapTest",
		},
	}
	InitLogger(&(c.LogConfig), c.BasicinfoConfig.AppName)
}

func TestInfo(t *testing.T) {
	Info("test info", WithPair("age", 20), WithPair("name", "小明"))
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("test info", WithPair("age", 20), WithPair("name", "小明"))
	}
}

func TestTempLogger_Debug(t *testing.T) {
	c := context.WithValue(context.TODO(), constant.RequestId, uuid.GenUUID16())
	RID(c).Debug("test log Request ID", WithPair("age", 20), WithPair("name", "小明"))
	// 在包外使用时, 可以把web框架比如*gin.Context实例直接传入
	// log.RID(c).Debug("test log Request ID", WithPair("age", 20), WithPair("name", "小明"))
}

func TestTempLogger_Debugf(t *testing.T) {
	c := context.WithValue(context.TODO(), constant.RequestId, uuid.GenUUID16())
	RID(c).Debugf("age=%d,name=%s\r\n", 20, "小明")
}
