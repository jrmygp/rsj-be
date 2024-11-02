package main

import (
	"server/config"
	"server/controllers"
	customerRepo "server/repositories/customer"
	portRepo "server/repositories/port"
	userRepo "server/repositories/user"
	customerService "server/services/customer"
	portService "server/services/port"
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

	portRepository := portRepo.NewRepository(db)
	portService := portService.NewService(portRepository)
	portController := controllers.NewPortController(portService)

	// Set up the router
	router := config.NewRouter(userController, userService, customerController, portController)

	// Start the server
	router.Run(":8080")
}
