package repository

import (
	"context"
	"gin-starter/internal/domain/entity"
)

type UserRepository interface {
	FindById(context context.Context, id int64) (*entity.UserEntity, error)
	FindByEmail(context context.Context, email string) (*entity.UserEntity, error)
	Save(context context.Context, userEntity *entity.UserEntity) error
	Update(context context.Context, userEntity *entity.UserEntity) error
}
