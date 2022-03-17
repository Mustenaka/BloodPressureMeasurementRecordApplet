package entity

// 血压记录表
type PatientBpRecord struct {
	RecordId     uint   // 血压记录id
	UserId       uint   // 用户id
	RecordDate   string // 血压记录日期
	RecordTime   string // 血压记录时间
	LowPressure  int16  // 低压
	HighPressure int16  // 高压
}
