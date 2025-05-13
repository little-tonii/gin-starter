package router

import (
	"gin-starter/internal/application/request"
	"gin-starter/internal/interface/middleware"
	"gin-starter/internal/shared/di"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(engine *gin.Engine, locator *di.Locator) {
	group := engine.Group("/user")
	{
		group.POST(
			"/register",
			middleware.BindingValidator[request.RegisterUserRequest](),
			locator.UserHandler.HandleRegisterUser(),
		)
		group.POST(
			"/login",
			middleware.BindingValidator[request.LoginUserRequest](),
			locator.UserHandler.HandleLoginUser(),
		)
	}
}
