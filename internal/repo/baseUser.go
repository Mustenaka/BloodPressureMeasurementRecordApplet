package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// baseUser 用户repo接口
type BaseUserRepo interface {
	GetUserByName(ctx context.Context, name string) (*model.BaseUser, error)
	GetUserById(ctx context.Context, id uint) (*model.BaseUser, error)
}
