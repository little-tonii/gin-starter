package service

import (
	"context"
	"errors"
	"gin-starter/internal/application/request"
	"gin-starter/internal/application/response"
	"gin-starter/internal/domain/entity"
	"gin-starter/internal/domain/repository"
	"gin-starter/internal/shared/constant"
	"gin-starter/internal/shared/utils"
	"net/http"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
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
		utils.Claims{UserId: userEntity.Id},
	)
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &response.LoginUserResponse{AccessToken: accessToken}, nil
}
