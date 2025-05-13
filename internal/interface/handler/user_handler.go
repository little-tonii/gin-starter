package handler

import (
	"errors"
	"gin-starter/internal/application/request"
	"gin-starter/internal/application/service"
	"gin-starter/internal/shared/constant"
	"net/http"
	"strings"

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
// @Failure		400 {object} godoc.MessagesResponse
// @Failure		409 {object} godoc.MessageResponse
// @Failure		500 {object} godoc.MessageResponse
// @Router		/user/register [post]
func (handler *UserHandler) HandleRegisterUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestRaw, exists := context.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			context.Error(errors.New("Không có dữ liệu request"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.RegisterUserRequest)
		request.Email = strings.ToLower(request.Email)
		if !ok {
			context.Error(errors.New("Không thể ép kiểu RegisterUserRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, err := handler.userService.RegisterUser(context.Request.Context(), request)
		if err != nil {
			context.Error(errors.New(err.Message))
			context.AbortWithStatus(err.StatusCode)
			return
		}
		context.JSON(http.StatusCreated, response)
	}
}

// Login 		godoc
// @Summary 	Đăng nhập
// @Produce 	application/json
// @Tags 		User
// @Param 		request body request.LoginUserRequest true "Request Body"
// @Success 	200 {object} response.LoginUserResponse
// @Failure		400 {object} godoc.MessagesResponse
// @Failure		401 {object} godoc.MessageResponse
// @Failure		500 {object} godoc.MessageResponse
// @Router		/user/login [post]
func (handler *UserHandler) HandleLoginUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestRaw, exists := context.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			context.Error(errors.New("Không có dữ liệu request"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.LoginUserRequest)
		if !ok {
			context.Error(errors.New("Không thể ép kiểu LoginUserRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, err := handler.userService.LoginUser(context.Request.Context(), request)
		if err != nil {
			context.Error(errors.New(err.Message))
			context.AbortWithStatus(err.StatusCode)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}
