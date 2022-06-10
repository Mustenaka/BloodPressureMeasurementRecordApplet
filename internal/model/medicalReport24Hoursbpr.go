package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

type MedicalReport24hoursbpr struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	DayHigh   int    `json:"day_high"`
	DayLow    int    `json:"day_low"`
	NightHigh int    `json:"night_high"`
	NightLow  int    `json:"night_low"`
	CreateAt  string `json:"create_at"`
	// 最大高血压
	MaxHigh     int    `json:"max_high"`
	MaxHighTime string `json:"max_high_time"`
	MaxLow      int    `json:"max_low"`
	MaxLowTime  string `json:"max_low_time"`
	// 最低高血压
	MinHigh     int    `json:"min_high"`
	MinHighTime string `json:"min_high_time"`
	MinLow      int    `json:"min_low"`
	MinLowTime  string `json:"min_low_time"`
}

// 获取表名称
func (MedicalReport24hoursbpr) TableName() string {
	return "medical_report_24hoursbprs"
}

// 判断有效性
func (medicalReport24hoursbpr *MedicalReport24hoursbpr) Validate() error {
	validate := validator.New()
	return validate.Struct(medicalReport24hoursbpr)
}
