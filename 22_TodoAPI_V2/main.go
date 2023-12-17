package main

import (
	"TodoAPI/common/app"
	"TodoAPI/common/postgresql"
	"TodoAPI/controller"
	"TodoAPI/persistence"
	"TodoAPI/service"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()
	configurationManager := app.NewConfiguraitonManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgresSqlConfig)
	todoRepository := persistence.New(ctx, dbPool)
	todoService := service.New(todoRepository)
	todoControler := controller.New(todoService)

	todoControler.RegisterRoutes(e)

	e.Start("localhost:8080")
}
