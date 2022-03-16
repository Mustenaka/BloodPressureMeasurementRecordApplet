package entity

// 药品供应表
type MedicalSupplie struct {
	MedicalId uint   // 药品使用id
	Metering  int    // 药品使用剂量
	Measuring string // 计量规格说明（克、毫克、毫升、升等）
}
