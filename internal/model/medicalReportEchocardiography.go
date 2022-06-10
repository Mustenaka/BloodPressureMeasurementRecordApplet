package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 心超报告
type MedicalReportEchocardiography struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Ef       int    `json:"ef"`
	Lvidd    int    `json:"lvidd"`
	Lvids    int    `json:"lvids"`
	CreateAt string `json:"create_at"` // 报告时间
}

// 获取表名称
func (MedicalReportEchocardiography) TableName() string {
	return "medical_report_echocardiographys"
}

// 判断有效性
func (medicalReportEchocardiography *MedicalReportEchocardiography) Validate() error {
	validate := validator.New()
	return validate.Struct(medicalReportEchocardiography)
}
