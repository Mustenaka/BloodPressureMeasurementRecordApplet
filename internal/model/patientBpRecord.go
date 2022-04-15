package model

import validator "gopkg.in/go-playground/validator.v9"

// 血压记录表
type PatientBpRecord struct {
	RecordId     uint   `json:"record_id"`     // 血压记录id
	UserId       uint   `json:"user_id"`       // 用户id
	RecordDate   string `json:"record_date"`   // 血压记录日期
	RecordTime   string `json:"record_time"`   // 血压记录时间
	LowPressure  int    `json:"low_pressure"`  // 低压
	HighPressure int    `json:"high_pressure"` // 高压
	HeartRate    int    `json:"heart_rate"`    // 心率
}

// 获取表名称
func (PatientBpRecord) TableName() string {
	return "patient_bp_records"
}

// 判断有效性
func (patientBpRecord *PatientBpRecord) Validate() error {
	validate := validator.New()
	return validate.Struct(patientBpRecord)
}
