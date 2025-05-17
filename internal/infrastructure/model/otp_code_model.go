package model

import (
	"gin-starter/internal/domain/entity"
	"time"
)

type OtpCodeModel struct {
	Id         int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Code       string     `gorm:"column:code;not null;index"`
	ExpiredAt  time.Time  `gorm:"column:expired_at;not null"`
	UserId     int64      `gorm:"column:user_id;not null;index"`
	ResetToken *string    `gorm:"column:reset_token;index"`
	User       *UserModel `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (OtpCodeModel) TableName() string {
	return "otp_codes"
}

func (model *OtpCodeModel) ToEntity() *entity.OtpCodeEntity {
	var userEntity *entity.UserEntity
	if model.User != nil {
		userEntity = model.User.ToEntity()
	}
	return &entity.OtpCodeEntity{
		Id:         model.Id,
		Code:       model.Code,
		ExpiredAt:  model.ExpiredAt,
		UserId:     model.UserId,
		User:       userEntity,
		ResetToken: model.ResetToken,
	}
}
