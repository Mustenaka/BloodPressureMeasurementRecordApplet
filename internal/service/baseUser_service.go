package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"context"
)

var _ BaseUserService = (*baseUserService)(nil)

// BaseUserService 定义用户操作服务接口
type BaseUserService interface {
	GetByName(ctx context.Context, name string) (*model.BaseUser, error)
	GetById(ctx context.Context, uid uint) (*model.BaseUser, error)
}

// baseUserService 实现UserService接口
type baseUserService struct {
	ur repo.BaseUserRepo
}

// 新用户服务
func NewBaseUserService(_ur repo.BaseUserRepo) *baseUserService {
	return &baseUserService{
		ur: _ur,
	}
}

// GetByName 通过用户名 查找用户
func (us *baseUserService) GetByName(ctx context.Context, name string) (*model.BaseUser, error) {
	if len(name) == 0 {
		return nil, errors.WithCode(code.ValidateErr, "用户名称不能为空")
	}
	return us.ur.GetBaseUserByName(ctx, name)
}

// GetById 根据用户ID查找用户
func (us *baseUserService) GetById(ctx context.Context, uid uint) (*model.BaseUser, error) {
	return us.ur.GetBaseUserById(ctx, uid)
}
