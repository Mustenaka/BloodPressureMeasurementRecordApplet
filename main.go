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
	// user := model.BaseUser{
	// 	// UserId:     11,
	// 	OpenId:     tools.RandomString(16),
	// 	UserName:   "李翠莲",
	// 	Tel:        "18778262136",
	// 	Email:      "leecuilian@outlook.com",
	// 	Permission: 3,
	// 	LastTime:   time.Now().Format("2006-01-02 15:04:05"),
	// 	Sex:        "女",
	// 	Status:     "关闭",
	// }

	var users []model.BaseUser
	model.DB.Where(&model.BaseUser{UserName: "李翠莲"}).Find(&users)

	fmt.Println(users[0])

	bpRecord := model.PatientBpRecord{
		UserId:       users[0].UserId,
		RecordDate:   time.Now().Format("2006-01-02"),
		RecordTime:   time.Now().Format("15:04:05"),
		LowPressure:  int16(tools.RandomInt(60, 90)),
		HighPressure: int16(tools.RandomInt(90, 150)),
	}
	model.DB.Create(&bpRecord)
}

func main() {
	fmt.Println("Beginning the Program!")

	RunProgram()

	fmt.Println("Done.")
}
