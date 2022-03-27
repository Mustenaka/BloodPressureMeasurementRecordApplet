package model

import validator "gopkg.in/go-playground/validator.v9"

// 治疗计划使用药物表
type TreatmentUseMedical struct {
	Id          uint
	TreatmentId uint
	MedicalId   uint
}

// 获取表名称
func (TreatmentUseMedical) TableName() string {
	return "treatment_use_medicals"
}

// 判断有效性
func (treatmentUseMedical *TreatmentUseMedical) Validate() error {
	validate := validator.New()
	return validate.Struct(treatmentUseMedical)
}
