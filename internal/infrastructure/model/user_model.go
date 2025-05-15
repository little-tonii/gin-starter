package model

import "gin-starter/internal/domain/entity"

type UserModel struct {
	Id           int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Password     string `gorm:"column:password;not null"`
	Email        string `gorm:"column:email;unique;not null"`
	TokenVersion int64  `gorm:"column:token_version;not null;default:0"`
}

func (UserModel) TableName() string {
	return "users"
}

func (model *UserModel) ToEntity() *entity.UserEntity {
	return &entity.UserEntity{
		Id:           model.Id,
		Password:     model.Password,
		Email:        model.Email,
		TokenVersion: model.TokenVersion,
	}
}
