package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// PatientInfoRepo 患者病历信息repo接口
type PatientInfoRepo interface {
	// 添加病历信息记录
	AddInfo(ctx context.Context, id uint) error

	// 获取记录
	GetInfoById(ctx context.Context, id uint) (*model.PatientInfo, error)

	// 更新记录
	UpdateInfoById(ctx context.Context, id uint) (*model.PatientInfo, error)
}
