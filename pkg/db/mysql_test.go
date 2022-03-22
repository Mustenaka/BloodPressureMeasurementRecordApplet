package db

import (
	"BloodPressure/internal/model"
	"BloodPressure/pkg/config"
	"fmt"
	"testing"
)

// 测试链接并寻找特定用户
func TestConnect(t *testing.T) {
	c := config.Load("./../config/config_back.ini")
	db := NewDefaultMysql(c.DBConfig)
	var users []model.BaseUser
	if err := db.Master().Where(&model.BaseUser{UserName: "李翠花"}).Find(&users); err.Error != nil {
		// 错误处理
		fmt.Println("没有找到该数据333")
	}
	for _, value := range users {
		fmt.Println(value.UserId, value.UserName)
	}
}
