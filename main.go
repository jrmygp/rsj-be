package main

import (
	"server/config"
	"server/controllers"
	"server/middleware"
	userRepo "server/repositories/user"
	userService "server/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DatabaseConnection()

	userRepository := userRepo.NewRepository(db)
	userService := userService.NewService(userRepository)
	userController := controllers.NewUserController(userService)

	router := gin.Default()
	router.Use(cors.Default())

	userRouter := router.Group("/user")
	userRouter.GET("", userController.FindAllUsers)
	userRouter.GET("/:id", userController.FindUserByID)
	userRouter.POST("/signup", middleware.RequireAuth(userRepository), userController.CreateUser)
	userRouter.POST("/login", userController.LoginUser)

	router.Run(":8080")
}
