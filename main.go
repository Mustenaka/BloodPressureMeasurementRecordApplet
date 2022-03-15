package main

import (
	"BloodPressure/pkg/log"
	"fmt"
	"os"
)

// Test hello
func RunProgram() {
	fmt.Println(os.Getwd())
	log.Infof("Test info", log.WithPair("age", 20), log.WithPair("name", "小明"))
	// user := model.BaseUser{
	// 	// UserId:     11,
	// 	OpenId:     random.RandomUpperString(16),
	// 	UserName:   "刘头",
	// 	Tel:        "18278262188",
	// 	Email:      "liutou@outlook.com",
	// 	Permission: 3,
	// 	LastTime:   time.Now().Format("2006-01-02 15:04:05"),
	// 	Sex:        "男",
	// 	Status:     "开启",
	// }
	// if err := model.DB.Create(&user); err.Error != nil {
	// 	// 错误处理
	// 	fmt.Println("无法插入数据")
	// }

	// 查询
	// var users []model.BaseUser
	// if err := model.DB.Where(&model.BaseUser{UserName: "李翠花"}).Find(&users); err.Error != nil {
	// 	// 错误处理
	// 	fmt.Println("没有找到该数据333")
	// }

	// 删除
	// model.DB.Where(&model.BaseUser{UserName: "翠花"}).Delete(&model.BaseUser{})
	// 修改
	// model.DB.Model(&model.BaseUser{}).Where(&model.BaseUser{UserName: "李翠花"}).Updates(model.BaseUser{Sex: "其他"})

	// 通过查询的用户插入高血压数据
	// for _, user := range users {
	// 	fmt.Println(user)

	// 	// bpRecord := model.PatientBpRecord{
	// 	// 	UserId:       user.UserId,
	// 	// 	RecordDate:   time.Now().Format("2006-01-02"),
	// 	// 	RecordTime:   time.Now().Format("15:04:05"),
	// 	// 	LowPressure:  int16(tools.RandomInt(60, 90)),
	// 	// 	HighPressure: int16(tools.RandomInt(90, 150)),
	// 	// }
	// 	// if err := model.DB.Create(&bpRecord); err.Error != nil {
	// 	// 	// 错误处理
	// 	// 	fmt.Println("无法插入数据")
	// 	// }
	// }

}

func main() {
	fmt.Println("Beginning the Program!")

	// 先加载package 因此数据库初始化先于这里的代码执行
	RunProgram()

	fmt.Println("Done.")
}
