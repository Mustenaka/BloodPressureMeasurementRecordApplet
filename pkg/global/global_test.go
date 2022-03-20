package global

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	conf := GetInstance()
	value := conf.GetConfigValue("database", "hostname")
	fmt.Println("1", value)
	t.Log(value)

	emptyValue := conf.GetConfigValue("UnknownList", "UnknownData")
	fmt.Println("2", emptyValue)
	t.Log(emptyValue)
}
