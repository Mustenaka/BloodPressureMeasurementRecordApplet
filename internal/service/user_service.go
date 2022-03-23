package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"context"
)

var _ UserService = (*userService)(nil)

// UserService 定义用户操作服务接口
type UserService interface {
	GetByName(ctx context.Context, name string) (*model.BaseUser, error)
	GetById(ctx context.Context, uid uint) (*model.BaseUser, error)
}

// userService 实现UserService接口
type userService struct {
	ur repo.BaseUserRepo
}

// 新用户服务
func NewUserService(_ur repo.BaseUserRepo) *userService {
	return &userService{
		ur: _ur,
	}
}

// GetByName 通过用户名 查找用户
func (us *userService) GetByName(ctx context.Context, name string) (*model.BaseUser, error) {
	if len(name) == 0 {
		return nil, errors.WithCode(code.ValidateErr, "用户名称不能为空")
	}
	return us.ur.GetUserByName(ctx, name)
}

// GetById 根据用户ID查找用户
func (us *userService) GetById(ctx context.Context, uid uint) (*model.BaseUser, error) {
	return us.ur.GetUserById(ctx, uid)
}
