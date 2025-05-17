package repository

import (
	"context"
	"gin-starter/internal/domain/entity"
)

type OtpCodeRepository interface {
	Save(ctx context.Context, otpCodeEntity *entity.OtpCodeEntity) error
	FindByUserIdAndCode(ctx context.Context, userId int64, code string) (*entity.OtpCodeEntity, error)
	DeleteByUserId(ctx context.Context, userId int64) error
	Update(ctx context.Context, otpCodeEntity *entity.OtpCodeEntity) error
	FindByResetToken(ctx context.Context, resetToken string) (*entity.OtpCodeEntity, error)
}
