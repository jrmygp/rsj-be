package config

import (
	"server/controllers"
	"server/middleware"

	userServices "server/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creates a new router with routes and middleware set up
func NewRouter(userController *controllers.UserController, userService userServices.Service, customerController *controllers.CustomerController, portController *controllers.PortController, costChargesController *controllers.CostChargesController, quotationController *controllers.QuotationController) *gin.Engine {
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
		user.Use(middleware.RequireAuth(userService), middleware.RequireRole(1, 2))
		user.GET("", userController.FindAllUsers)
		user.GET("/:id", userController.FindUserByID)
	}

	masterData := router.Group("/master-data")
	{
		masterData.Use(middleware.RequireAuth(userService), middleware.RequireRole(1, 2))
		// Customer routes
		masterData.GET("/customer/no-pagination", customerController.FindAllCustomersWithoutPagination)
		masterData.GET("/customer", customerController.FindAll)
		masterData.GET("/customer/:id", customerController.FindCustomerByID)
		masterData.POST("/customer", customerController.CreateCustomer)
		masterData.PATCH("/customer/:id", customerController.EditCustomer)
		masterData.DELETE("/customer/:id", customerController.DeleteCustomer)

		// Port routes
		masterData.GET("/port/no-pagination", portController.FindAllPortsWithoutPagination)
		masterData.GET("/port", portController.FindAll)
		masterData.GET("/port/:id", portController.FindPortByID)
		masterData.POST("/port", portController.CreatePort)
		masterData.PATCH("/port/:id", portController.EditPort)
		masterData.DELETE("/port/:id", portController.DeletePort)

		// Cost Charges routes
		masterData.GET("/cost-charges/no-pagination", costChargesController.FindAllCostChargesWithoutPagination)
		masterData.GET("/cost-charges", costChargesController.FindAll)
		masterData.GET("/cost-charges/:id", costChargesController.FindCostChargeByID)
		masterData.POST("/cost-charges", costChargesController.CreateCostCharge)
		masterData.PATCH("/cost-charges/:id", costChargesController.EditCostCharge)
		masterData.DELETE("/cost-charges/:id", costChargesController.DeleteCostCharge)

	}

	quotation := router.Group("/quotation")
	{
		quotation.Use(middleware.RequireAuth(userService), middleware.RequireRole(1, 2))

		quotation.GET("/no-pagination", quotationController.FindAllQuotationsWithoutPagination)
		quotation.GET("", quotationController.FindAll)
		quotation.GET("/:id", quotationController.FindQuotationByID)
		quotation.POST("", quotationController.CreateQuotation)
		quotation.PATCH("/:id", quotationController.EditQuotation)
		quotation.DELETE("/:id", quotationController.DeleteQuotation)
		quotation.GET("/generate-pdf/:id", quotationController.GeneratePDF)
	}

	return router
}
