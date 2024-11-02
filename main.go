package main

import (
	"server/config"
	"server/controllers"
	userRepo "server/repositories/user"
	userService "server/services/user"
)

func main() {
	// Initialize database and repositories
	db := config.DatabaseConnection()

	// Initialize repository, service, and controller
	userRepository := userRepo.NewRepository(db)
	userService := userService.NewService(userRepository)
	userController := controllers.NewUserController(userService)

	// Set up the router
	router := config.NewRouter(userController, userService)

	// Start the server
	router.Run(":8080")
}
