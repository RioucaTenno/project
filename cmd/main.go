// @title User Orders API
// @version 1.0
// @description Тестовое задание на практику
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "project_go/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"project_go/internal/handlers"
	"project_go/internal/middleware"
	"project_go/internal/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.ConnectDatabase()

	r := gin.Default()

	r.Use(middleware.ErrorHandler())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/auth/login", handlers.Login)
	r.POST("/users", handlers.CreateUser)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUserByID)
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)

		protected.POST("/users/:id/orders", handlers.CreateOrder)
		protected.GET("/users/:id/orders", handlers.GetOrders)
	}

	r.Run()
}
