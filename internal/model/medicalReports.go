package model

import validator "gopkg.in/go-playground/validator.v9"

// 体检报告表
type MedicalReports struct {
	Id     uint // 检查报告id
	UserId uint // 用户id
}

// 获取表名称
func (MedicalReports) TableName() string {
	return "medical_reports"
}

// 判断有效性
func (medicalReports *MedicalReports) Validate() error {
	validate := validator.New()
	return validate.Struct(medicalReports)
}
