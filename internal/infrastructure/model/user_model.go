package model

import "health-care-system/internal/domain/entity"

type UserModel struct {
	Id          int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Username    string     `gorm:"column:username;unique;not null"`
	Password    string     `gorm:"column:password;not null"`
	Email       string     `gorm:"column:email;unique;not null"`
	PhoneNumber string     `gorm:"column:phone_number;not null"`
	RoleId      int64      `gorm:"column:role_id;not null"`
	Role        *RoleModel `gorm:"foreignKey:RoleId;references:Id"`
}

func (UserModel) TableName() string {
	return "users"
}

func (model *UserModel) ToEntity() *entity.UserEntity {
	return &entity.UserEntity{
		Id:          model.Id,
		Username:    model.Username,
		Password:    model.Password,
		PhoneNumber: model.PhoneNumber,
		Email:       model.Email,
		Role:        model.Role.ToEntity(),
	}
}
