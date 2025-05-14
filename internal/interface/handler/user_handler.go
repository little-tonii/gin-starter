package handler

import (
	"errors"
	"fmt"
	"gin-starter/internal/application/request"
	"gin-starter/internal/application/response"
	"gin-starter/internal/application/service"
	"gin-starter/internal/shared/constant"
	"gin-starter/internal/shared/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UserHandler struct {
	userService *service.UserService
	redisClient *redis.Client
}

func NewUserHandler(redisClient *redis.Client, userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		redisClient: redisClient,
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

// Profile 		godoc
// @Summary 	Thông tin người dùng
// @Produce 	application/json
// @Tags 		User
// @Security	BearerAuth
// @Security	OAuth2Password
// @Success 	200 {object} response.ProfileUserResponse
// @Failure		400 {object} godoc.MessagesResponse
// @Failure		401 {object} godoc.MessageResponse
// @Failure		500 {object} godoc.MessageResponse
// @Router		/user/profile [get]
func (handler *UserHandler) HandleProfileUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		claimsRaw, exists := context.Get(constant.ContextKey.CLAIMS)
		if !exists {
			context.Error(errors.New("Không có thông tin xác thực"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		claims, ok := claimsRaw.(*utils.Claims)
		if !ok {
			context.Error(errors.New("Không thể ép kiểu Claims"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr, err := utils.GetOrCache[response.ProfileUserResponse, response.ErrorResponse](
			context.Request.Context(),
			handler.redisClient,
			fmt.Sprintf("user:profile:%v", claims.UserId),
			24*time.Hour,
			func() (*response.ProfileUserResponse, *response.ErrorResponse) {
				return handler.userService.ProfileUser(context.Request.Context(), claims)
			},
		)
		if err != nil {
			context.Error(err)
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if customErr != nil {
			context.Error(errors.New(customErr.Message))
			context.AbortWithStatus(customErr.StatusCode)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}
