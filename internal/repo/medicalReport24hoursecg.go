package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// MedicalReport24hoursecgRepo 检验报告-24小时动态心超数据 repo接口
type MedicalReport24hoursecgRepo interface {
	// 添加
	AddEcg(ctx context.Context, id uint, data string, average int) error
	AddEcgWithTime(ctx context.Context, id uint, data string, average int, createAt string) error

	// 获取
	GetEcgById(ctx context.Context, id uint) ([]model.MedicalReport24hoursecg, error)
	GetEcgByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReport24hoursecg, error)

	// 删除
	DeleteEcgByUserId(ctx context.Context, userId uint) error
	DeleteEcgById(ctx context.Context, id uint) error
}
