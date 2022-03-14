package main

import (
	"BloodPressure/model"
	"BloodPressure/tools"
	"fmt"
	"time"
)

// func startServer() {
// 	ExampleRun()
// }

// Test hello
func RunProgram() {
	user := model.BaseUser{
		// UserId:     11,
		OpenId:     tools.RandomString(16),
		UserName:   "李翠莲",
		Tel:        "18778262136",
		Email:      "leecuilian@outlook.com",
		Permission: 3,
		LastTime:   time.Now().Format("2006-01-02 15:04:05"),
		Sex:        "女",
		Status:     "关闭",
	}
	model.DB.Create(&user)
}

func main() {
	fmt.Println("Beginning the Program!")

	RunProgram()

	fmt.Println("Done.")
}
