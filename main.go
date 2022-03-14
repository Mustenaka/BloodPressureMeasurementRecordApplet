package main

import (
	"BloodPressure/model"
	"BloodPressure/tools"
	"fmt"
	"time"
)

// Test hello
func RunProgram() {
	user := model.BaseUser{
		// UserId:     11,
		OpenId:     tools.RandomString(16),
		UserName:   "张测试",
		Tel:        "18783612206",
		Email:      "amum123@outlook.com",
		Permission: 2,
		LastTime:   time.Now().Format("2006-01-02 15:04:05"),
	}
	model.Create(&user)
}

func main() {
	fmt.Println("Beginning the Program!")

	RunProgram()

	fmt.Println("Done.")
}
