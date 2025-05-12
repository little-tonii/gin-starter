package service

import (
	"errors"
	"gin-starter/internal/application/request"
	"gin-starter/internal/application/response"
	"gin-starter/internal/domain/entity"
	"gin-starter/internal/domain/repository"
	"gin-starter/internal/infrastructure/utils"
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

func (service *UserService) RegisterUser(request *request.RegisterUserRequest) (*response.RegisterUserResponse, *response.ErrorResponse) {
	existedUser, _ := service.userRepository.FindByEmail(request.Email)
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
	err = service.userRepository.Save(userEntity)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &response.ErrorResponse{
				Message:    "Email đã được sử dụng",
				StatusCode: http.StatusInternalServerError,
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
