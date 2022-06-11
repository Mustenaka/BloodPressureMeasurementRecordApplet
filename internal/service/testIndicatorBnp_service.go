package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ TestIndicatorBnpService = (*testIndicatorBnpService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type TestIndicatorBnpService interface {
	// 添加一条记录
	AddById(ctx context.Context, uid uint, data int) error
	AddByIdWithTime(ctx context.Context, uid uint, data int, createAt string) error
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.TestIndicatorBnp, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TestIndicatorBnp, error)
}

// testIndicatorBnpService 实现UserService接口
type testIndicatorBnpService struct {
	tr repo.TestIndicatorBnpRepo
}

// 新血压记录服务
func NewTestIndicatorBnpService(_ur repo.TestIndicatorBnpRepo) *testIndicatorBnpService {
	return &testIndicatorBnpService{
		tr: _ur,
	}
}

// 添加Bnp记录
func (us *testIndicatorBnpService) AddById(ctx context.Context, uid uint, data int) error {
	return us.tr.AddBnp(ctx, uid, data)
}

// 自定义时间添加Bnp记录
func (us *testIndicatorBnpService) AddByIdWithTime(ctx context.Context, uid uint, data int, createAt string) error {
	return us.tr.AddBnpWithTime(ctx, uid, data, createAt)
}

// 查询
func (us *testIndicatorBnpService) GetById(ctx context.Context, uid uint) ([]model.TestIndicatorBnp, error) {
	return us.tr.GetBnpById(ctx, uid)
}

// 通过limit限制查询
func (us *testIndicatorBnpService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TestIndicatorBnp, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.tr.GetBnpById(ctx, uid)
	}
	return us.tr.GetBnpByIdLimit(ctx, uid, limit)
}
