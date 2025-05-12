package repository

import (
	"health-care-system/internal/domain/entity"
)

type UserRepository interface {
	FindById(id int64) (*entity.UserEntity, error)
	FindByEmail(email string) (*entity.UserEntity, error)
	FindByUsername(username string) (*entity.UserEntity, error)
	Save(userEntity *entity.UserEntity) error
	Update(userEntity *entity.UserEntity) error
}
