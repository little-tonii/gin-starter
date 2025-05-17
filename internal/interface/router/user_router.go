package router

import (
	"gin-starter/internal/application/request"
	"gin-starter/internal/interface/handler"
	"gin-starter/internal/interface/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(engine *gin.Engine, handler *handler.UserHandler) {
	group := engine.Group("/user")
	{
		group.POST(
			"/register",
			middleware.BindingValidator[request.RegisterUserRequest](),
			handler.HandleRegisterUser(),
		)
		group.POST(
			"/login",
			middleware.BindingValidator[request.LoginUserRequest](),
			handler.HandleLoginUser(),
		)
		group.POST(
			"/forgot-password",
			middleware.BindingValidator[request.ForgotPasswordUserRequest](),
			handler.HandleForgotPasswordUser(),
		)
		group.POST(
			"/verify-otp-reset-password",
			middleware.BindingValidator[request.VerifyOtpResetPasswordUserRequest](),
			handler.HandleVerifyOtpResetPasswordUser(),
		)
		group.POST(
			"/reset-password",
			middleware.BindingValidator[request.ResetPasswordUserRequest](),
			handler.HandleResetPasswordUser(),
		)
	}

	authGroup := group.Group("", middleware.Authentication())
	{
		authGroup.GET("/profile", handler.HandleProfileUser())
		authGroup.POST(
			"/change-password",
			middleware.BindingValidator[request.ChangePasswordUserRequest](),
			handler.HandleChangePasswordUser(),
		)
	}
}
