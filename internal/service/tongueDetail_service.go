package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ TongueDetailService = (*tongueDetailService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type TongueDetailService interface {
	// 添加一条记录
	AddById(ctx context.Context, uid uint, tongue, tongueCoating, pulse string) error
	AddByIdWithTime(ctx context.Context, uid uint, tongue, tongueCoating, pulse string, createAt string) error
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.TongueDetail, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TongueDetail, error)
}

// treatmentPlanService 实现UserService接口
type tongueDetailService struct {
	tr repo.TongueDetailRepo
}

// 新血压记录服务
func NewTongueDetailService(_ur repo.TongueDetailRepo) *tongueDetailService {
	return &tongueDetailService{
		tr: _ur,
	}
}

// 添加一条舌苔脉象记录
func (us *tongueDetailService) AddById(ctx context.Context, uid uint, tongue, tongueCoating, pulse string) error {
	return us.tr.AddTongue(ctx, uid, tongue, tongueCoating, pulse)
}

// 添加一条带时间的舌苔脉象记录
func (us *tongueDetailService) AddByIdWithTime(ctx context.Context, uid uint, tongue, tongueCoating, pulse string, createAt string) error {
	return us.tr.AddTongueWithTime(ctx, uid, tongue, tongueCoating, pulse, createAt)
}

// 查询
func (us *tongueDetailService) GetById(ctx context.Context, uid uint) ([]model.TongueDetail, error) {
	return us.tr.GetTongueById(ctx, uid)
}

// 通过limit限制查询
func (us *tongueDetailService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TongueDetail, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.tr.GetTongueById(ctx, uid)
	}
	return us.tr.GetTongueByIdLimit(ctx, uid, limit)
}
