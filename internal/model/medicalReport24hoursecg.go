package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

type MedicalReport24hoursecg struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Data     string `json:"data"`      // 数据
	Average  int    `json:"average"`   // 平均值
	CreateAt string `json:"create_at"` // 报告时间
}

// 获取表名称
func (MedicalReport24hoursecg) TableName() string {
	return "medical_report_24hoursecgs"
}

// 判断有效性
func (medicalReport24hoursecg *MedicalReport24hoursecg) Validate() error {
	validate := validator.New()
	return validate.Struct(medicalReport24hoursecg)
}
