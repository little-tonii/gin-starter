package repository

import (
	"context"
	"gin-starter/internal/domain/entity"
)

type OtpCodeRepository interface {
	Save(context context.Context, otpCodeEntity *entity.OtpCodeEntity) error
	FindByUserIdAndCode(context context.Context, userId int64, code string) (*entity.OtpCodeEntity, error)
	DeleteByUserId(context context.Context, userId int64) error
	Update(context context.Context, otpCodeEntity *entity.OtpCodeEntity) error
}
