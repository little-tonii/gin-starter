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
	return func(c *gin.Context) {
		requestRaw, exists := c.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			c.Error(errors.New("Không có dữ liệu request"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.RegisterUserRequest)
		request.Email = strings.ToLower(request.Email)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu RegisterUserRequest"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr := handler.userService.RegisterUser(c.Request.Context(), request)
		if customErr != nil {
			c.Error(errors.New(customErr.Message))
			c.AbortWithStatus(customErr.StatusCode)
			return
		}
		c.JSON(http.StatusCreated, response)
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
	return func(c *gin.Context) {
		requestRaw, exists := c.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			c.Error(errors.New("Không có dữ liệu request"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.LoginUserRequest)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu LoginUserRequest"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr := handler.userService.LoginUser(c.Request.Context(), request)
		if customErr != nil {
			c.Error(errors.New(customErr.Message))
			c.AbortWithStatus(customErr.StatusCode)
			return
		}
		c.JSON(http.StatusOK, response)
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
	return func(c *gin.Context) {
		claimsRaw, exists := c.Get(constant.ContextKey.CLAIMS)
		if !exists {
			c.Error(errors.New("Không có thông tin xác thực"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		claims, ok := claimsRaw.(*utils.Claims)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu Claims"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr, err := utils.GetOrCache[response.ProfileUserResponse, response.ErrorResponse](
			c.Request.Context(),
			handler.redisClient,
			fmt.Sprintf("user:profile:%v", claims.UserId),
			24*time.Hour,
			func() (*response.ProfileUserResponse, *response.ErrorResponse) {
				return handler.userService.ProfileUser(c.Request.Context(), claims)
			},
		)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if customErr != nil {
			c.Error(errors.New(customErr.Message))
			c.AbortWithStatus(customErr.StatusCode)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

// ChangePassword 	godoc
// @Summary 		Đổi mật khẩu
// @Produce 		application/json
// @Tags 			User
// @Security		BearerAuth
// @Security		OAuth2Password
// @Param 			request body request.ChangePasswordUserRequest true "Request Body"
// @Success 		200 {object} response.ChanagePasswordUserResponse
// @Failure			400 {object} godoc.MessagesResponse
// @Failure			401 {object} godoc.MessageResponse
// @Failure			500 {object} godoc.MessageResponse
// @Router			/user/change-password [post]
func (handler *UserHandler) HandleChangePasswordUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsRaw, exists := c.Get(constant.ContextKey.CLAIMS)
		if !exists {
			c.Error(errors.New("Không có thông tin xác thực"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		claims, ok := claimsRaw.(*utils.Claims)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu Claims"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		requestRaw, exists := c.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			c.Error(errors.New("Không có dữ liệu request"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.ChangePasswordUserRequest)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu ChangePasswordUserRequest"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr := handler.userService.ChangePasswordUser(c.Request.Context(), claims, request)
		if customErr != nil {
			c.Error(errors.New(customErr.Message))
			c.AbortWithStatus(customErr.StatusCode)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

// ForgotPassword 	godoc
// @Summary 		Quên mật khẩu
// @Produce 		application/json
// @Tags 			User
// @Param 			request body request.ForgotPasswordUserRequest true "Request Body"
// @Success 		200 {object} response.ForgotPasswordUserResponse
// @Failure			400 {object} godoc.MessagesResponse
// @Failure			404 {object} godoc.MessageResponse
// @Failure			500 {object} godoc.MessageResponse
// @Router			/user/forgot-password [post]
func (handler *UserHandler) HandleForgotPasswordUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestRaw, exists := c.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			c.Error(errors.New("Không có dữ liệu request"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.ForgotPasswordUserRequest)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu ForgotPasswordUserRequest"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr := handler.userService.ForgotPasswordUser(c.Request.Context(), request)
		if customErr != nil {
			c.Error(errors.New(customErr.Message))
			c.AbortWithStatus(customErr.StatusCode)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

// VerifyOtpResetPassword 	godoc
// @Summary 				Xác thực OTP đặt lại mật khẩu
// @Produce 				application/json
// @Tags 					User
// @Param 					request body request.VerifyOtpResetPasswordUserRequest true "Request Body"
// @Success 				200 {object} response.VerifyOtpResetPasswordUserRepsonse
// @Failure					400 {object} godoc.MessagesResponse
// @Failure					500 {object} godoc.MessageResponse
// @Router					/user/verify-otp-reset-password [post]
func (handler *UserHandler) HandleVerifyOtpResetPasswordUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestRaw, exists := c.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			c.Error(errors.New("Không có dữ liệu request"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.VerifyOtpResetPasswordUserRequest)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu VerifyOtpResetPasswordRequest"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr := handler.userService.VerifyOtpResetPasswordUser(c.Request.Context(), request)
		if customErr != nil {
			c.Error(errors.New(customErr.Message))
			c.AbortWithStatus(customErr.StatusCode)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

// ResetPassword 	godoc
// @Summary 		Đặt lại mật khẩu
// @Produce 		application/json
// @Tags 			User
// @Param 			request body request.ResetPasswordUserRequest true "Request Body"
// @Success 		200 {object} response.ResetPasswordUserResponse
// @Failure			400 {object} godoc.MessagesResponse
// @Failure			500 {object} godoc.MessageResponse
// @Router			/user/reset-password [post]
func (handler *UserHandler) HandleResetPasswordUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestRaw, exists := c.Get(constant.ContextKey.REQUEST_DATA)
		if !exists {
			c.Error(errors.New("Không có dữ liệu request"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, ok := requestRaw.(*request.ResetPasswordUserRequest)
		if !ok {
			c.Error(errors.New("Không thể ép kiểu ResetPasswordUserRequest"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, customErr := handler.userService.ResetPasswordUser(c.Request.Context(), request)
		if customErr != nil {
			c.Error(errors.New(customErr.Message))
			c.AbortWithStatus(customErr.StatusCode)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
