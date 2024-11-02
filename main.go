package main

import (
	"server/config"
	"server/controllers"
	customerRepo "server/repositories/customer"
	userRepo "server/repositories/user"
	customerService "server/services/customer"
	userService "server/services/user"
)

func main() {
	// Initialize database and repositories
	db := config.DatabaseConnection()

	// Initialize repository, service, and controller
	userRepository := userRepo.NewRepository(db)
	userService := userService.NewService(userRepository)
	userController := controllers.NewUserController(userService)

	customerRepository := customerRepo.NewRepository(db)
	customerService := customerService.NewService(customerRepository)
	customerController := controllers.NewCustomerController(customerService)

	// Set up the router
	router := config.NewRouter(userController, userService, customerController)

	// Start the server
	router.Run(":8080")
}
