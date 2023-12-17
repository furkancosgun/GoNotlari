package controller

import (
	"TodoAPI/controller/request"
	"TodoAPI/controller/response"
	"TodoAPI/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	todoService service.ITodoService
}

func New(todoService service.ITodoService) TodoController {
	return TodoController{todoService: todoService}
}

func (todoController *TodoController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/todos", todoController.GetTodos)
	e.GET("/api/v1/todos/:id", todoController.GetTodoById)
	e.POST("/api/v1/todos", todoController.CreateTodo)
	e.PUT("/api/v1/todos/:id", todoController.UpdateTodo)
	e.DELETE("/api/v1/todos/:id", todoController.DeleteTodoById)
}
func (todoControler *TodoController) CreateTodo(e echo.Context) error {
	var createTodoRequest = request.CreateTodoRequst{}
	err := e.Bind(&createTodoRequest)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	err = todoControler.todoService.CreateTodo(createTodoRequest)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	return e.NoContent(http.StatusCreated)
}
func (todoControler *TodoController) UpdateTodo(e echo.Context) error {
	paramId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	var updateTodoRequest = request.UpdateTodoRequest{}
	err = e.Bind(&updateTodoRequest)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	err = todoControler.todoService.UpdateTodo(paramId, updateTodoRequest)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	return e.NoContent(http.StatusOK)
}
func (todoControler *TodoController) DeleteTodoById(e echo.Context) error {
	paramId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	err = todoControler.todoService.DeleteTodoById(paramId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	return e.NoContent(http.StatusOK)
}

func (todoControler *TodoController) GetTodos(e echo.Context) error {
	paramTitle := e.QueryParam("title")
	if paramTitle != "" {
		return e.JSON(http.StatusOK, todoControler.todoService.GetTodosByTitle(paramTitle))
	}
	paramStatus := e.QueryParam("status")
	if paramStatus != "" {
		paramStatus, _ := strconv.ParseBool(paramStatus)
		return e.JSON(http.StatusOK, todoControler.todoService.GetTodosByStatus(paramStatus))
	}
	return e.JSON(http.StatusOK, todoControler.todoService.GetAllTodos())
}

func (todoControler *TodoController) GetTodoById(e echo.Context) error {
	idParam, atoiErr := strconv.Atoi(e.Param("id"))
	if atoiErr != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorMessage: atoiErr.Error()})
	}
	product, err := todoControler.todoService.GetTodoById(idParam)
	if err != nil {
		return e.JSON(http.StatusNotFound, response.ErrorResponse{ErrorMessage: err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}
