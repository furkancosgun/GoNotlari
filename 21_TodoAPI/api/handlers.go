package api

import (
	"strconv"
	"todo_api/model"
	"todo_api/service"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{
		TodoService: todoService,
	}
}

func (h *TodoHandler) GetAllTodos(c echo.Context) error {
	todos, err := h.TodoService.GetAll()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}
	return c.JSON(200, todos)
}

func (h *TodoHandler) GetTodoByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoService.GetByID(id)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Todo not found"})
	}
	return c.JSON(200, todo)
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(400, map[string]string{"error": "Bad Request"})
	}
	createdTodo, err := h.TodoService.Create(todo)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}
	return c.JSON(201, createdTodo)
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(400, map[string]string{"error": "Bad Request"})
	}
	updatedTodo, err := h.TodoService.Update(id, todo)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Todo not found"})
	}
	return c.JSON(200, updatedTodo)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.TodoService.Delete(id)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Todo not found"})
	}
	return c.NoContent(204)
}
