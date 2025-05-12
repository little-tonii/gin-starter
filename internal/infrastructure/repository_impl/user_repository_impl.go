package repository_impl

import (
	"gin-starter/internal/domain/entity"
	"gin-starter/internal/infrastructure/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepositoryImpl(database *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		database: database,
	}
}

func (repository *UserRepositoryImpl) FindById(id int64) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		Where("id = ?", id).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (repository *UserRepositoryImpl) Save(userEntity *entity.UserEntity) error {
	userModel := model.UserModel{
		Password: userEntity.Password,
		Email:    userEntity.Email,
	}
	result := repository.database.Create(&userModel)
	if result.Error != nil {
		return result.Error
	}
	userEntity.Id = userModel.Id
	return nil
}

func (repository *UserRepositoryImpl) Update(userEntity *entity.UserEntity) error {
	result := repository.database.
		Model(&model.UserModel{}).
		Where("id = ?", userEntity.Id).
		Updates(map[string]any{
			"password": userEntity.Password,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		Where("email = ?", email).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}
