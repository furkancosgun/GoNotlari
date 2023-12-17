package main

import (
	"context"
	"product_api/common/app"
	"product_api/common/postgresql"
	"product_api/controller"
	"product_api/persistence"
	"product_api/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configurationManager := app.NewConfiguraitonManager()

	dbPool := postgresql.GetConnectionPool(context.Background(), configurationManager.PostgresSqlConfig)
	productRepository := persistence.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)

	e.Start("localhost:8080")
}
