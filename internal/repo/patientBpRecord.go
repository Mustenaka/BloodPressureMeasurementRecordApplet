package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// PatientBpRecordRepo 血压记录repo接口
type PatientBpRecordRepo interface {
	// 添加血压记录
	AddRecord(ctx context.Context, id uint, low, high int) error

	// 获取全部记录组
	GetRecordById(ctx context.Context, id uint) ([]model.PatientBpRecord, error)
	// // 获取指定日期（当日）记录组
	// GetRecordByIdLimitDate(ctx context.Context, id uint) ([]*model.PatientBpRecord, error)

}
