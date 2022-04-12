package model

import validator "gopkg.in/go-playground/validator.v9"

// 治疗计划表
type TreatmentPlan struct {
	TreatmentId    uint   // 治疗方案id
	UserId         uint   // 用户id
	Plan           string // 治疗计划
	Note           string // 治疗计划备注（如禁忌症）
	CreateDatetime string // 创建时间
	Status         string // 状态{“生效”，“失效”}
}

// 获取表名称
func (TreatmentPlan) TableName() string {
	return "treatment_plans"
}

// 判断有效性
func (treatmentPlan *TreatmentPlan) Validate() error {
	validate := validator.New()
	return validate.Struct(treatmentPlan)
}
