package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// PatientBpRecordRepo 血压记录repo接口
type PatientBpRecordRepo interface {
	// 添加血压记录
	AddRecord(ctx context.Context, id uint, low, high, heartRate int) error

	// 获取记录组
	GetRecordById(ctx context.Context, id uint) ([]model.PatientBpRecord, error)
	GetRecordByIdLimitDays(ctx context.Context, id uint, limitdays int) ([]model.PatientBpRecord, error)
}
