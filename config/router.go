package config

import (
	"server/controllers"
	"server/middleware"

	userServices "server/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creates a new router with routes and middleware set up
func NewRouter(userController *controllers.UserController, userService userServices.Service) *gin.Engine {
	router := gin.Default()

	// Enable CORS for all routes
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Define routes
	trxWithoutPrefix := router.Group("")
	{
		trxWithoutPrefix.POST("/login", userController.LoginUser)
	}

	user := router.Group("/user")
	{
		user.Use(middleware.RequireAuth(userService)) // Protect /user routes with authentication
		user.GET("", userController.FindAllUsers)
		user.GET("/:id", userController.FindUserByID)
		user.POST("/create-user", userController.CreateUser)
	}

	return router
}
