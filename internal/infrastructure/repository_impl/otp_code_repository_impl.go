package repository_impl

import (
	"context"
	"gin-starter/internal/domain/entity"
	"gin-starter/internal/infrastructure/model"

	"gorm.io/gorm"
)

type OtpCodeRepositoryImpl struct {
	database *gorm.DB
}

func NewOtpCodeRepositoryImpl(database *gorm.DB) *OtpCodeRepositoryImpl {
	return &OtpCodeRepositoryImpl{
		database: database,
	}
}

func (repository *OtpCodeRepositoryImpl) FindByUserIdAndCode(context context.Context, userId int64, code string) (*entity.OtpCodeEntity, error) {
	var otpCodeModel model.OtpCodeModel
	result := repository.database.
		WithContext(context).
		Where("user_id = ? AND code = ?", userId, code).
		First(&otpCodeModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return otpCodeModel.ToEntity(), nil
}

func (repository *OtpCodeRepositoryImpl) DeleteByUserId(context context.Context, userId int64) error {
	result := repository.database.
		WithContext(context).
		Where("user_id = ?", userId).
		Delete(&model.OtpCodeModel{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *OtpCodeRepositoryImpl) Save(context context.Context, otpCodeEntity *entity.OtpCodeEntity) error {
	otpCodeModel := model.OtpCodeModel{
		Code:      otpCodeEntity.Code,
		ExpiredAt: otpCodeEntity.ExpiredAt,
		UserId:    otpCodeEntity.UserId,
	}
	result := repository.database.
		WithContext(context).
		Create(&otpCodeModel)
	if result.Error != nil {
		return result.Error
	}
	otpCodeEntity.Id = otpCodeModel.Id
	return nil
}

func (repository *OtpCodeRepositoryImpl) Update(context context.Context, otpCodeEntity *entity.OtpCodeEntity) error {
	result := repository.database.
		WithContext(context).
		Model(&model.OtpCodeModel{}).
		Where("id = ?", otpCodeEntity.Id).
		Updates(map[string]any{
			"reset_token": otpCodeEntity.ResetToken,
		})
	return result.Error
}
