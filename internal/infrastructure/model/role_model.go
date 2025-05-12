package model

import "health-care-system/internal/domain/entity"

type RoleModel struct {
	Id   int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;unique;not null"`
}

func (RoleModel) TableName() string {
	return "roles"
}

func (model *RoleModel) ToEntity() *entity.RoleEntity {
	return &entity.RoleEntity{
		Id:   model.Id,
		Name: model.Name,
	}
}
