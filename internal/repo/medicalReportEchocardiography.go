package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// MedicalReportEchocardiography 检验报告-心超参数 repo接口
type MedicalReportEchocardiographyRepo interface {
	// 添加
	AddEchocardiography(ctx context.Context, id uint, ef, lvidd, lvids int) error
	AddEchocardiographyWithTime(ctx context.Context, id uint, ef, lvidd, lvids int, createAt string) error

	// 获取
	GetEchocardiographyById(ctx context.Context, id uint) ([]model.MedicalReportEchocardiography, error)
	GetEchocardiographyByIdLimit(ctx context.Context, id uint, limit int) ([]model.MedicalReportEchocardiography, error)

	// 删除
	DeleteEchocardiographyByUserId(ctx context.Context, userId uint) error
	DeleteEchocardiographyById(ctx context.Context, id uint) error
}
