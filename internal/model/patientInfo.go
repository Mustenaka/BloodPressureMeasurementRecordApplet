package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 患者病历信息
type PatientInfo struct {
	// patient_infos 内容
	PatientId                 uint   `json:"patient_id"`                  // 病历信息
	UserId                    uint   `json:"user_id"`                     // 对应用户
	RealName                  string `json:"real_name"`                   // 真实姓名
	Sex                       string `json:"sex"`                         // 性别
	Birthday                  string `json:"birthday"`                    // 生日
	Tel                       string `json:"tel"`                         // 电话号码
	IsMarried                 bool   `json:"is_married"`                  // 0-未婚、1-已婚
	HbpYears                  int    `json:"hbp_years"`                   // 高血压患病时间（年）
	Anamnesis                 string `json:"anamnesis"`                   // 既往病史(对应表格1~12)
	IsSmoking                 bool   `json:"is_smoking"`                  // 是否吸烟
	SmokingHistory            int    `json:"smoking_history"`             // 吸烟史（年）
	SmokingDaily              int    `json:"smoking_daily"`               // 日吸烟数
	IsDrink                   bool   `json:"is_drink"`                    // 是否饮酒
	DrinkHistory              int    `json:"drink_history"`               // 饮酒史（年）
	DrinkDaily                int    `json:"drink_daily"`                 // 每日饮酒量
	PatientHeight             int    `json:"patient_height"`              // 身高
	PatientWeight             int    `json:"patient_weight"`              // 体重
	PatientWaistCircumference int    `json:"patient_waist_circumference"` // 腰围
	PatientChestCircumference int    `json:"patient_chest_circumference"` // 胸围
	PatientHipCircumference   int    `json:"patient_hip_circumference"`   // 臀围
	IsTakeChineseMedicine     bool   `json:"is_take_chinese_medicine"`    // 是否服用中药
	AntihypertensivePlan      string `json:"antihypertensive_plan"`       // 降压方案
	IsNondrugControlPlan      bool   `json:"is_nondrug_control_plan"`     // 是否非药物控制手段
	NondrugControlPlan        string `json:"nondrug_control_plan"`        // 非药物控制手段内容
}

// 获取表名称
func (PatientInfo) TableName() string {
	return "patient_infos"
}

// 判断有效性
func (patientInfo *PatientInfo) Validate() error {
	validate := validator.New()
	return validate.Struct(patientInfo)
}
