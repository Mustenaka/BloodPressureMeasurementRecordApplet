package log

import (
	"BloodPressure/pkg/global"
	"BloodPressure/pkg/log/constant"
	"BloodPressure/tools/uuid"
	"context"
	"testing"
)

func init() {
	defer Sync()
	conf := global.GetInstance()
	logConfig := LogConfig{
		Level:      conf.GetConfigValue("logconfig", "level"),
		FileName:   conf.GetConfigValue("logconfig", "file-name"),
		TimeFormat: constant.TimeLayout,
		MaxSize:    conf.GetConfigValueInt("logconfig", "max-size"),
		MaxBackups: conf.GetConfigValueInt("logconfig", "max-backups"),
		MaxAge:     conf.GetConfigValueInt("logconfig", "max-age"),
		Compress:   conf.GetConfigValueBool("logconfig", "compress"),
		LocalTime:  conf.GetConfigValueBool("logconfig", "local-time"),
		Console:    conf.GetConfigValueBool("logconfig", "console"),
	}
	InitLogger(&(logConfig), conf.GetConfigValue("basicinfo", "appName"))
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
