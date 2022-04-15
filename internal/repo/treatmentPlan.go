package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// PatientBpRecordRepo 血压记录repo接口
type TreatmentPlanRepo interface {
	// 添加血压记录
	AddPlan(ctx context.Context, id uint, plan, note string) error

	// 获取记录组
	GetPlanById(ctx context.Context, id uint) ([]model.TreatmentPlan, error)
	GetPlanByIdLimit(ctx context.Context, id uint, limit int) ([]model.TreatmentPlan, error)
}
