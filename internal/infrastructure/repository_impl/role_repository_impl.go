package repository_impl

import (
	"health-care-system/internal/domain/entity"
	"health-care-system/internal/infrastructure/model"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	database *gorm.DB
}

func NewRoleRepositoryImpl(datababse *gorm.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{
		database: datababse,
	}
}

func (repository *RoleRepositoryImpl) FindAll() ([]*entity.RoleEntity, error) {
	var roleModels []model.RoleModel
	result := repository.database.Find(&roleModels)
	if result.Error != nil {
		return nil, result.Error
	}
	roleEntities := make([]*entity.RoleEntity, 0, len(roleModels))
	for i, _ := range roleModels {
		roleEntities = append(roleEntities, roleModels[i].ToEntity())
	}
	return roleEntities, nil
}
