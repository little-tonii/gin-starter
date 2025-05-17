package repository

import (
	"context"
	"gin-starter/internal/domain/entity"
)

type UserRepository interface {
	FindById(ctx context.Context, id int64) (*entity.UserEntity, error)
	FindByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	Save(ctx context.Context, userEntity *entity.UserEntity) error
	Update(ctx context.Context, userEntity *entity.UserEntity) error
}
