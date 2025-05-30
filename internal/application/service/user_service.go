package service

import (
	"context"
	"errors"
	"gin-starter/internal/application/request"
	"gin-starter/internal/application/response"
	"gin-starter/internal/application/routine"
	"gin-starter/internal/domain/entity"
	"gin-starter/internal/domain/repository"
	"gin-starter/internal/shared/constant"
	"gin-starter/internal/shared/utils"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository    repository.UserRepository
	otpCodeRepository repository.OtpCodeRepository
}

func NewUserService(userRepository repository.UserRepository, otpCodeRepository repository.OtpCodeRepository) *UserService {
	return &UserService{
		userRepository:    userRepository,
		otpCodeRepository: otpCodeRepository,
	}
}

func (service *UserService) RegisterUser(ctx context.Context, request *request.RegisterUserRequest) (*response.RegisterUserResponse, *response.ErrorResponse) {
	existedUser, err := service.userRepository.FindByEmail(ctx, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	if existedUser != nil {
		return nil, &response.ErrorResponse{
			Message:    "Email đã được sử dụng",
			StatusCode: http.StatusConflict,
		}
	}
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	var userEntity *entity.UserEntity = &entity.UserEntity{
		Email:    request.Email,
		Password: hashedPassword,
	}
	err = service.userRepository.Save(ctx, userEntity)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &response.ErrorResponse{
				Message:    "Email đã được sử dụng",
				StatusCode: http.StatusConflict,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	return &response.RegisterUserResponse{Message: "Đăng ký thành công"}, nil
}

func (service *UserService) LoginUser(ctx context.Context, request *request.LoginUserRequest) (*response.LoginUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindByEmail(ctx, request.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Tài khoản hoặc mật khẩu không chính xác",
				StatusCode: http.StatusUnauthorized,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	if !utils.CheckPasswordHash(request.Password, userEntity.Password) {
		return nil, &response.ErrorResponse{
			Message:    "Tài khoản hoặc mật khẩu không chính xác",
			StatusCode: http.StatusUnauthorized,
		}
	}
	accessToken, err := utils.GenerateAccessToken(
		constant.Environment.JWT_SECRET_KEY,
		utils.Claims{
			UserId:       userEntity.Id,
			TokenVersion: userEntity.TokenVersion,
		},
	)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &response.LoginUserResponse{AccessToken: accessToken}, nil
}

func (service *UserService) ProfileUser(ctx context.Context, claims *utils.Claims) (*response.ProfileUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindById(ctx, claims.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Người dùng không tồn tại",
				StatusCode: http.StatusUnauthorized,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	if claims.TokenVersion != userEntity.TokenVersion {
		return nil, &response.ErrorResponse{
			Message:    "Người dùng chưa đăng nhập",
			StatusCode: http.StatusUnauthorized,
		}
	}
	return &response.ProfileUserResponse{
		Id:    userEntity.Id,
		Email: userEntity.Email,
	}, nil
}

func (service *UserService) ChangePasswordUser(ctx context.Context, claims *utils.Claims, request *request.ChangePasswordUserRequest) (*response.ChanagePasswordUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindById(ctx, claims.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Người dùng không tồn tại",
				StatusCode: http.StatusUnauthorized,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	if claims.TokenVersion != userEntity.TokenVersion {
		return nil, &response.ErrorResponse{
			Message:    "Người dùng chưa đăng nhập",
			StatusCode: http.StatusUnauthorized,
		}
	}
	if !utils.CheckPasswordHash(request.OldPassword, userEntity.Password) {
		return nil, &response.ErrorResponse{
			Message:    "Tài khoản hoặc mật khẩu không chính xác",
			StatusCode: http.StatusUnauthorized,
		}
	}
	userEntity.Password, err = utils.HashPassword(request.NewPassword)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	userEntity.TokenVersion++
	err = service.userRepository.Update(ctx, userEntity)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &response.ChanagePasswordUserResponse{
		Message: "Đổi mật khẩu thành công",
	}, nil
}

func (service *UserService) ForgotPasswordUser(ctx context.Context, request *request.ForgotPasswordUserRequest) (*response.ForgotPasswordUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Người dùng không tồn tại",
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	code, err := utils.CreateOtpCode(8)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	otpCodeEntity := &entity.OtpCodeEntity{
		UserId:    userEntity.Id,
		Code:      code,
		ExpiredAt: time.Now().Add(5 * time.Minute),
	}
	err = service.otpCodeRepository.Save(ctx, otpCodeEntity)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	go func() {
		if err := routine.SendOtpCodeToEmail(
			"Gin Starter",
			userEntity.Email,
			"Quên mật khẩu",
			otpCodeEntity.Code,
		); err != nil {
			log.Printf("%v", err)
		}
	}()
	return &response.ForgotPasswordUserResponse{Message: "OTP đã được gửi đến email của bạn"}, nil
}

func (service *UserService) VerifyOtpResetPasswordUser(ctx context.Context, request *request.VerifyOtpResetPasswordUserRequest) (*response.VerifyOtpResetPasswordUserRepsonse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Mã OTP không hợp lệ",
				StatusCode: http.StatusBadRequest,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	otpCodeEntity, err := service.otpCodeRepository.FindByUserIdAndCode(ctx, userEntity.Id, request.OtpCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Mã OTP không hợp lệ",
				StatusCode: http.StatusBadRequest,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	if otpCodeEntity.ExpiredAt.Before(time.Now()) {
		return nil, &response.ErrorResponse{
			Message:    "Mã OTP không hợp lệ",
			StatusCode: http.StatusBadRequest,
		}
	}
	randomUUID := uuid.New().String()
	otpCodeEntity.ResetToken = &randomUUID
	err = service.otpCodeRepository.Update(ctx, otpCodeEntity)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &response.VerifyOtpResetPasswordUserRepsonse{
		ResetToken: *otpCodeEntity.ResetToken,
	}, nil
}

func (service *UserService) ResetPasswordUser(ctx context.Context, request *request.ResetPasswordUserRequest) (*response.ResetPasswordUserResponse, *response.ErrorResponse) {
	otpCodeEntity, err := service.otpCodeRepository.FindByResetToken(ctx, request.ResetToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Reset token không hợp lệ",
				StatusCode: http.StatusBadRequest,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	userEntity, err := service.userRepository.FindById(ctx, otpCodeEntity.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.ErrorResponse{
				Message:    "Reset token không hợp lệ",
				StatusCode: http.StatusBadRequest,
			}
		} else {
			return nil, &response.ErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}
	userEntity.Password, err = utils.HashPassword(request.NewPassword)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	err = service.userRepository.Update(ctx, userEntity)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := service.otpCodeRepository.DeleteByUserId(ctx, userEntity.Id); err != nil {
			log.Printf("%v", err)
		}
	}()
	return &response.ResetPasswordUserResponse{Message: "Đổi mật khẩu thành công, vui lòng đăng nhập lại"}, nil
}
