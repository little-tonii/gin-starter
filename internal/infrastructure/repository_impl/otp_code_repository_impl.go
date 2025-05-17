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

func (repository *OtpCodeRepositoryImpl) FindByCode(context context.Context, code string) ([]*entity.OtpCodeEntity, error) {
	otpCodeModels := make([]model.OtpCodeModel, 0)
	err := repository.database.
		WithContext(context).
		Preload("User").
		Where("code = ?", code).
		Find(&otpCodeModels).Error
	if err != nil {
		return nil, err
	}
	otpCodeEntities := make([]*entity.OtpCodeEntity, 0, len(otpCodeModels))
	for _, otpCodeModel := range otpCodeModels {
		otpCodeEntities = append(otpCodeEntities, otpCodeModel.ToEntity())
	}
	return otpCodeEntities, nil
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
