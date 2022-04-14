package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ PatientInfoService = (*patientInfoService)(nil)

// PatientInfoService 定义用户操作服务接口
type PatientInfoService interface {
	// 添加患者信息
	Add(ctx context.Context, patientInfo *model.PatientInfo) error
	// 查询患者信息
	GetById(ctx context.Context, uid uint) (*model.PatientInfo, error)
	// 更新患者信息
	Update(ctx context.Context, patientInfo *model.PatientInfo) error
	// 删除患者信息
	DeleteById(ctx context.Context, uid uint) error
}

// patientInfoService 实现UserService接口
type patientInfoService struct {
	pir repo.PatientInfoRepo
}

// 新患者信息服务(患者信息需要和用户基本信息组合)
func NewPatientInfoService(_pir repo.PatientInfoRepo) *patientInfoService {
	return &patientInfoService{
		pir: _pir,
	}
}

// 添加患者信息
func (pis *patientInfoService) Add(ctx context.Context, patientInfo *model.PatientInfo) error {
	return pis.pir.AddInfo(ctx, patientInfo)
}

// 获取患者信息
func (pis *patientInfoService) GetById(ctx context.Context, uid uint) (*model.PatientInfo, error) {
	return pis.pir.GetInfoById(ctx, uid)
}

// 更新患者信息
func (pis *patientInfoService) Update(ctx context.Context, patientInfo *model.PatientInfo) error {
	return pis.pir.UpdateInfoById(ctx, patientInfo.UserId, patientInfo)
}

// 删除患者信息
func (pis *patientInfoService) DeleteById(ctx context.Context, uid uint) error {
	return pis.pir.DeleteInfoByID(ctx, uid)
}
