package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// PatientInfoRepo 患者病历信息repo接口
type PatientInfoRepo interface {
	// 添加病历信息记录
	AddInfo(ctx context.Context, patientInfo *model.PatientInfo) error

	// 获取记录
	GetInfoById(ctx context.Context, id uint) (*model.PatientInfo, error)

	// 更新记录
	UpdateInfoById(ctx context.Context, id uint, patientInfo *model.PatientInfo) error

	// 删除记录
	DeleteInfoByID(ctx context.Context, id uint) error
}
