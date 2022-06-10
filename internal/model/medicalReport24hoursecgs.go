package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

type MedicalReport24hoursecgs struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Data     string `json:"data"`      // 数据
	Average  int    `json:"average"`   // 平均值
	CreateAt string `json:"create_at"` // 报告时间
}

// 获取表名称
func (MedicalReport24hoursecgs) TableName() string {
	return "medical_report_24hoursecgs"
}

// 判断有效性
func (medicalReport24hoursecgs *MedicalReport24hoursecgs) Validate() error {
	validate := validator.New()
	return validate.Struct(medicalReport24hoursecgs)
}
