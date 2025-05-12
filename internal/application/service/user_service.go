package service

import (
	"errors"
	"health-care-system/internal/application/request"
	"health-care-system/internal/application/response"
	"health-care-system/internal/domain/entity"
	"health-care-system/internal/domain/repository"
	"health-care-system/internal/infrastructure/utils"
	"net/http"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository repository.UserRepository
	roleRepository repository.RoleRepository
}

func NewUserService(userRepository repository.UserRepository, roleRepository repository.RoleRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (service *UserService) RegisterUser(request *request.RegisterUserRequest) (*response.RegisterUserResponse, *response.ErrorResponse) {
	existedUser, _ := service.userRepository.FindByUsername(request.Username)
	if existedUser != nil {
		return nil, &response.ErrorResponse{
			Message:    "Username đã được sử dụng",
			StatusCode: http.StatusConflict,
		}
	}
	existedUser, _ = service.userRepository.FindByEmail(request.Email)
	if existedUser != nil {
		return nil, &response.ErrorResponse{
			Message:    "Email đã được sử dụng",
			StatusCode: http.StatusConflict,
		}
	}
	roleEntities, err := service.roleRepository.FindAll()
	if err != nil {
		return nil, &response.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	var rolePatient *entity.RoleEntity
	for _, role := range roleEntities {
		if role != nil && role.Name == "patient" {
			rolePatient = role
			break
		}
	}
	if rolePatient == nil {
		return nil, &response.ErrorResponse{
			Message:    "Bệnh nhân đăng ký thất bại",
			StatusCode: http.StatusInternalServerError,
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
		Username:    request.Username,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    hashedPassword,
		Role:        rolePatient,
	}
	err = service.userRepository.Save(userEntity)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &response.ErrorResponse{
				Message:    "Email hoặc username đã được sử dụng",
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
