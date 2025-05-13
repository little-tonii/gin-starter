package main

import (
	"gin-starter/internal/infrastructure/config"
	"gin-starter/internal/interface/middleware"
	"gin-starter/internal/interface/router"
	"gin-starter/internal/shared/constant"
	"gin-starter/internal/shared/di"
	"io"
	"log"
	"os"

	_ "gin-starter/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Title
// @version 1.0
// @description description
// @host localhost:8080
// @BasePath /

// @securityDefinitions.oauth2.password OAuth2Password
// @tokenUrl http://localhost:8080/user/login
// @scope.read Grants read access
// @scope.write Grants write access

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	constant.InitContextKey()

	if errs := constant.LoadEnvironment(); errs != nil && len(errs) > 0 {
		log.Fatalf("Không thể tải biến môi trường: %v", errs)
	}

	if error := config.InitializeDatabase(); error != nil {
		log.Fatal(error)
	}
	defer config.CloseDatabase()

	locator := di.InitLocator()

	logFile, error := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if error != nil {
		log.Fatal(error)
	}
	defer logFile.Close()

	gin.ForceConsoleColor()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	engine.Use(cors.New(corsConfig))
	engine.Use(middleware.Recovery())
	engine.Use(middleware.ErrorHandler())
	engine.NoRoute(middleware.NoRoute())

	router.InitUserRouter(engine, locator.UserHandler)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if error := engine.Run(":8080"); error != nil {
		log.Fatalf("Khởi động thất bại: %v", error)
	}
}
