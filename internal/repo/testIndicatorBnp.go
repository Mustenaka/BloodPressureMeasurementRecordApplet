package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// TestIndicatorBnp 检验报告-BNP参数 repo接口
type TestIndicatorBnp interface {
	// 添加
	AddBnp(ctx context.Context, id uint, data int) error
	AddBnpWithTime(ctx context.Context, id uint, data int, createAt string) error

	// 获取
	GetBnpById(ctx context.Context, id uint) ([]model.TestIndicatorBnp, error)
	GetBnpByIdLimit(ctx context.Context, id uint, limit int) ([]model.TestIndicatorBnp, error)

	// 删除
	DeleteBnpByUserId(ctx context.Context, userId uint) error
	DeleteBnpById(ctx context.Context, id uint) error
}
