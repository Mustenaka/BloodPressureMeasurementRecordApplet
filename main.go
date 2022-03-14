package main

import (
	"BloodPressure/model"
	"BloodPressure/tools"
	"fmt"
	"time"
)

func startServer() {
	ExampleRun()
}

// Test hello
func RunProgram() {
	user := model.BaseUser{
		// UserId:     11,
		OpenId:     tools.RandomString(16),
		UserName:   "翠花",
		Tel:        "18278362137",
		Email:      "cuihua@outlook.com",
		Permission: 3,
		LastTime:   time.Now().Format("2006-01-02 15:04:05"),
		Sex:        "女",
		Status:     "开启",
	}
	model.Create(&user)
}

func main() {
	fmt.Println("Beginning the Program!")

	RunProgram()

	fmt.Println("Done.")
}
