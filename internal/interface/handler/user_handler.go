package handler

import (
	"errors"
	"health-care-system/internal/application/request"
	"health-care-system/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register 	godoc
// @Summary 	Đăng ký
// @Produce 	application/json
// @Tags 		User
// @Param 		request body request.RegisterUserRequest true "Request Body"
// @Success 	201 {object} response.RegisterUserResponse
// @Failure		400 {object} godoc.ErrorsResponse
// @Failure		409 {object} godoc.ErrorResponse
// @Failure		500 {object} godoc.ErrorResponse
// @Router		/user/register [post]
func (handler *UserHandler) HandleRegisterUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestRaw, exists := context.Get("request_data")
		if !exists {
			context.Error(errors.New("Không có dữ liệu request"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.RegisterUserRequest)
		if !ok {
			context.Error(errors.New("Không thể ép kiểu RegisterUserRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, err := handler.userService.RegisterUser(request)
		if err != nil {
			context.Error(errors.New(err.Message))
			context.AbortWithStatus(err.StatusCode)
			return
		}
		context.JSON(http.StatusCreated, response)
	}
}
