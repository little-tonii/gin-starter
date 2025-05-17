package service

import (
	"context"
	"crypto/rand"
	"errors"
	"gin-starter/internal/application/request"
	"gin-starter/internal/application/response"
	"gin-starter/internal/application/routine"
	"gin-starter/internal/domain/entity"
	"gin-starter/internal/domain/repository"
	"gin-starter/internal/shared/constant"
	"gin-starter/internal/shared/utils"
	"net/http"
	"time"

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

func (service *UserService) RegisterUser(context context.Context, request *request.RegisterUserRequest) (*response.RegisterUserResponse, *response.ErrorResponse) {
	existedUser, err := service.userRepository.FindByEmail(context, request.Email)
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
	err = service.userRepository.Save(context, userEntity)
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

func (service *UserService) LoginUser(context context.Context, request *request.LoginUserRequest) (*response.LoginUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindByEmail(context, request.Username)
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

func (service *UserService) ProfileUser(context context.Context, claims *utils.Claims) (*response.ProfileUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindById(context, claims.UserId)
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

func (service *UserService) ChangePasswordUser(context context.Context, claims *utils.Claims, request *request.ChangePasswordUserRequest) (*response.ChanagePasswordUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindById(context, claims.UserId)
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
	if claims.TokenVersion < userEntity.TokenVersion {
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
	err = service.userRepository.Update(context, userEntity)
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

func (service *UserService) ForgotPasswordUser(context context.Context, request *request.ForgotPasswordUserRequest) (*response.ForgotPasswordUserResponse, *response.ErrorResponse) {
	userEntity, err := service.userRepository.FindByEmail(context, request.Email)
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
	buffer := make([]byte, 8)
	_, err = rand.Read(buffer)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	for i := range 8 {
		buffer[i] = (buffer[i] % 10) + '0'
	}
	otpCodeEntity := &entity.OtpCodeEntity{
		UserId:    userEntity.Id,
		Code:      string(buffer),
		ExpiredAt: time.Now().Add(5 * time.Minute),
	}
	err = service.otpCodeRepository.Save(context, otpCodeEntity)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	go routine.SendOtpCodeToEmail(
		"Gin Starter",
		userEntity.Email,
		"Quên mật khẩu",
		otpCodeEntity.Code,
	)
	return &response.ForgotPasswordUserResponse{Message: "OTP đã được gửi đến email của bạn"}, nil
}
