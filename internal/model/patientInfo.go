package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// 患者病历信息
type PatientInfo struct {
	// patient_infos 内容
	PatientId                 uint   // 病历信息
	UserId                    uint   // 对应用户
	RealName                  string // 真实姓名
	Sex                       string // 性别
	Birthday                  string // 生日
	Tel                       string // 电话号码
	IsMarried                 bool   // 0-未婚、1-已婚
	HbpYears                  int    // 高血压患病时间（年）
	Anamnesis                 string // 既往病史(对应表格1~12)
	IsSmoking                 bool   // 是否吸烟
	SmokingHistory            int    // 吸烟史（年）
	SmokingDaily              int    // 日吸烟数
	IsDrink                   bool   // 是否饮酒
	DrinkHistory              int    // 饮酒史（年）
	DrinkDaily                int    // 每日饮酒量
	PatientHeight             int    // 身高
	PatientWeight             int    // 体重
	PatientWaistCircumference int    // 腰围
	PatientChestCircumference int    // 胸围
	PatientHipCircumference   int    // 臀围
	IsTakeChineseMedicine     bool   // 是否服用中药
	AntihypertensivePlan      string // 降压方案
	IsNondrugControlPlan      bool   // 是否非药物控制手段
	NondrugControlPlan        string // 非药物控制手段内容
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
