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
		OpenId:     tools.RandomUpperString(16),
		UserName:   "王五",
		Tel:        "18278266123",
		Email:      "wangwu@outlook.com",
		Permission: 3,
		LastTime:   time.Now().Format("2006-01-02 15:04:05"),
		Sex:        "男",
		Status:     "开启",
	}
	if err := model.DB.Create(&user); err.Error != nil {
		// 错误处理
		fmt.Println("无法插入数据")
	}

	var users []model.BaseUser
	if err := model.DB.Where(&model.BaseUser{UserName: "李翠花"}).Find(&users); err.Error != nil {
		// 错误处理
		fmt.Println("没有找到该数据333")
	}

	// 删除
	// model.DB.Where(&model.BaseUser{UserName: "翠花"}).Delete(&model.BaseUser{})
	// 修改
	// model.DB.Model(&model.BaseUser{}).Where(&model.BaseUser{UserName: "李翠花"}).Updates(model.BaseUser{Sex: "其他"})

	for _, user := range users {
		fmt.Println(user)

		// bpRecord := model.PatientBpRecord{
		// 	UserId:       user.UserId,
		// 	RecordDate:   time.Now().Format("2006-01-02"),
		// 	RecordTime:   time.Now().Format("15:04:05"),
		// 	LowPressure:  int16(tools.RandomInt(60, 90)),
		// 	HighPressure: int16(tools.RandomInt(90, 150)),
		// }
		// if err := model.DB.Create(&bpRecord); err.Error != nil {
		// 	// 错误处理
		// 	fmt.Println("无法插入数据")
		// }
	}

}

func main() {
	fmt.Println("Beginning the Program!")

	// 先加载package 因此数据库初始化先于这里的代码执行
	RunProgram()

	fmt.Println("Done.")
}
