package repository

import "health-care-system/internal/domain/entity"

type RoleRepository interface {
	FindAll() ([]*entity.RoleEntity, error)
}
