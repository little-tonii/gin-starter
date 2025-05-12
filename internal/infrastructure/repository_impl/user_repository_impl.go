package repository_impl

import (
	"health-care-system/internal/domain/entity"
	"health-care-system/internal/infrastructure/model"

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
		Preload("Role").
		Where("id = ?", id).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (repository *UserRepositoryImpl) Save(userEntity *entity.UserEntity) error {
	userModel := model.UserModel{
		Username:    userEntity.Username,
		Password:    userEntity.Password,
		PhoneNumber: userEntity.PhoneNumber,
		Email:       userEntity.Email,
		RoleId:      userEntity.Role.Id,
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
			"username":     userEntity.Username,
			"password":     userEntity.Password,
			"email":        userEntity.Email,
			"phone_number": userEntity.PhoneNumber,
			"role_id":      userEntity.Role.Id,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		Preload("Role").
		Where("email = ?", email).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (repository *UserRepositoryImpl) FindByUsername(username string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		Preload("Role").
		Where("username = ?", username).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}
