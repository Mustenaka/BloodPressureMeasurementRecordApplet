package model

import validator "gopkg.in/go-playground/validator.v9"

// 药品供应表
type MedicalSupplie struct {
	MedicalId uint   // 药品使用id
	Metering  int    // 药品使用剂量
	Measuring string // 计量规格说明（克、毫克、毫升、升等）
}

// 获取表名称
func (MedicalSupplie) TableName() string {
	return "medical_supplies"
}

// 判断有效性
func (medicalSupplie *MedicalSupplie) Validate() error {
	validate := validator.New()
	return validate.Struct(medicalSupplie)
}
