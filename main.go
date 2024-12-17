package main

import (
	"server/config"
	"server/controllers"
	costChargesRepo "server/repositories/costCharges"
	customerRepo "server/repositories/customer"
	invoiceRepo "server/repositories/invoice"
	portRepo "server/repositories/port"
	quotationRepo "server/repositories/quotation"
	shipperRepo "server/repositories/shipper"
	userRepo "server/repositories/user"
	costChargesService "server/services/costCharges"
	customerService "server/services/customer"
	invoiceService "server/services/invoice"
	portService "server/services/port"
	quotationService "server/services/quotation"
	shipperService "server/services/shipper"
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

	costChargesRepository := costChargesRepo.NewRepository(db)
	costChargesService := costChargesService.NewService(costChargesRepository)
	costChargesController := controllers.NewCostChargesController(costChargesService)

	quotationRepository := quotationRepo.NewRepository(db)
	quotationService := quotationService.NewService(quotationRepository, customerRepository, userRepository, portRepository)
	quotationController := controllers.NewQuotationController(quotationService)

	shipperRepository := shipperRepo.NewRepository(db)
	shipperService := shipperService.NewService(shipperRepository)
	shipperController := controllers.NewShipperController(shipperService)

	invoiceRepository := invoiceRepo.NewRepository(db)
	invoiceService := invoiceService.NewService(invoiceRepository, customerRepository, shipperRepository)
	invoiceController := controllers.NewInvoiceController(invoiceService)

	// Set up the router
	router := config.NewRouter(userController, userService, customerController, portController, costChargesController, quotationController, shipperController, invoiceController)

	// Start the server
	router.Run(":8080")
}
