package repository_impl

import (
	"context"
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

func (repository *UserRepositoryImpl) FindById(ctx context.Context, id int64) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		WithContext(ctx).
		Where("id = ?", id).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, userEntity *entity.UserEntity) error {
	userModel := model.UserModel{
		Password: userEntity.Password,
		Email:    userEntity.Email,
	}
	result := repository.database.
		WithContext(ctx).
		Create(&userModel)
	if result.Error != nil {
		return result.Error
	}
	userEntity.Id = userModel.Id
	return nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, userEntity *entity.UserEntity) error {
	result := repository.database.
		WithContext(ctx).
		Model(&model.UserModel{}).
		Where("id = ?", userEntity.Id).
		Updates(map[string]any{
			"password":      userEntity.Password,
			"token_version": userEntity.TokenVersion,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		WithContext(ctx).
		Where("email = ?", email).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}
