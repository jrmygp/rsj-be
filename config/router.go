package config

import (
	"server/controllers"
	"server/middleware"

	userServices "server/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creates a new router with routes and middleware set up
func NewRouter(userController *controllers.UserController, userService userServices.Service, customerController *controllers.CustomerController) *gin.Engine {
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
		trxWithoutPrefix.POST("/create-user", userController.CreateUser)
	}

	user := router.Group("/user")
	{
		user.Use(middleware.RequireAuth(userService)) // Protect /user routes with authentication
		user.GET("", userController.FindAllUsers)
		user.GET("/:id", userController.FindUserByID)
	}

	masterData := router.Group("/master-data")
	{
		masterData.Use(middleware.RequireAuth(userService))
		masterData.GET("/customer", customerController.FindAll)
		masterData.GET("/customer/:id", customerController.FindCustomerByID)
		masterData.POST("/customer", customerController.CreateCustomer)
		masterData.PATCH("/customer/:id", customerController.EditCustomer)
		masterData.DELETE("/customer/:id", customerController.DeleteCustomer)
	}

	return router
}
