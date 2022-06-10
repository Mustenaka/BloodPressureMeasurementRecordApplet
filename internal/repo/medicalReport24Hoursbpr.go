package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// MedicalReport24HoursbprRepo 检验报告- 24小时动态血压报告参数 repo接口
type MedicalReport24HoursbprRepo interface {
	// 添加
	Addbpr(ctx context.Context, id uint, report *model.MedicalReport24hoursbpr) error

	// 获取
	GetbprById(ctx context.Context, id uint) ([]model.MedicalReport24hoursbpr, error)
	GetbprByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReport24hoursbpr, error)

	// 删除
	DeletebprByUserId(ctx context.Context, userId uint) error
	DeletebprById(ctx context.Context, id uint) error
}
