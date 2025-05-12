package router

import (
	"health-care-system/internal/application/request"
	"health-care-system/internal/interface/middleware"
	"health-care-system/internal/shared/di"

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
	}
}
