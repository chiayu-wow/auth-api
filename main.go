package main

import (
	"github.com/chiayu/auth-api/handler"
	"github.com/chiayu/auth-api/repository"
	"github.com/chiayu/auth-api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	r.Run(":8080")
}
