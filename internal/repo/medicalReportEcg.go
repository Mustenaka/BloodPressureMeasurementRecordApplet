package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// MedicalReportEcg 体检报告-心电图参数 repo接口
type MedicalReportEcg interface {
	// 添加
	AddEcg(ctx context.Context, id uint, data string) error
	AddEcgWithTime(ctx context.Context, id uint, data string, createAt string) error

	// 获取
	GetEcgById(ctx context.Context, id uint) ([]model.MedicalReportEcg, error)
	GetEcgByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReportEcg, error)

	// 删除
	DeleteEcgByUserId(ctx context.Context, userId uint) error
	DeleteEcgById(ctx context.Context, id uint) error
}
