package repository

import (
	"gin-starter/internal/domain/entity"
)

type UserRepository interface {
	FindById(id int64) (*entity.UserEntity, error)
	FindByEmail(email string) (*entity.UserEntity, error)
	Save(userEntity *entity.UserEntity) error
	Update(userEntity *entity.UserEntity) error
}
