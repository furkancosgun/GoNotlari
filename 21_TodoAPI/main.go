package main

import (
	"todo_api/api"
	"todo_api/repository"
	"todo_api/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Dependency Injection
	todoRepo := repository.NewInMemoryTodoRepository()
	todoService := service.NewTodoService(todoRepo)
	todoHandler := api.NewTodoHandler(*todoService)

	// Routes
	e.GET("/todos", todoHandler.GetAllTodos)
	e.GET("/todos/:id", todoHandler.GetTodoByID)
	e.POST("/todos", todoHandler.CreateTodo)
	e.PUT("/todos/:id", todoHandler.UpdateTodo)
	e.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// Start server
	e.Start(":8080")
}
