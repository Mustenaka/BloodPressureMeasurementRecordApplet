package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

type MedicalReportEcgs struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Data     string `json:"data"`      // 数据
	CreateAt string `json:"create_at"` // 报告时间
}

// 获取表名称
func (MedicalReportEcgs) TableName() string {
	return "medical_report_ecgs"
}

// 判断有效性
func (medicalReportEcgs *MedicalReportEcgs) Validate() error {
	validate := validator.New()
	return validate.Struct(medicalReportEcgs)
}
