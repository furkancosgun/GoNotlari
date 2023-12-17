package service

import (
	"TodoAPI/controller/request"
	"TodoAPI/domain"
	"TodoAPI/persistence"
)

type ITodoService interface {
	CreateTodo(createTodoRequest request.CreateTodoRequst) error
	GetAllTodos() []domain.Todo
	GetTodosByTitle(title string) []domain.Todo
	GetTodosByStatus(status bool) []domain.Todo
	GetTodoById(id int) (domain.Todo, error)
	UpdateTodo(id int, updateTodoRequest request.UpdateTodoRequest) error
	DeleteTodoById(id int) error
}

type TodoService struct {
	todoRepository persistence.ITodoRepository
}

// CreateTodo implements ITodoService.
func (todoService *TodoService) CreateTodo(createTodoRequest request.CreateTodoRequst) error {
	return todoService.todoRepository.CreateTodo(createTodoRequest)
}

// DeleteTodoById implements ITodoService.
func (todoService *TodoService) DeleteTodoById(id int) error {
	return todoService.todoRepository.DeleteTodoById(id)
}

// GetAllTodos implements ITodoService.
func (todoService *TodoService) GetAllTodos() []domain.Todo {
	return todoService.todoRepository.GetAllTodos()
}

// GetTodoById implements ITodoService.
func (todoService *TodoService) GetTodoById(id int) (domain.Todo, error) {
	return todoService.todoRepository.GetTodoById(id)
}

// GetTodosByStatus implements ITodoService.
func (todoService *TodoService) GetTodosByStatus(status bool) []domain.Todo {
	return todoService.todoRepository.GetTodosByStatus(status)
}

// GetTodosByTitle implements ITodoService.
func (todoService *TodoService) GetTodosByTitle(title string) []domain.Todo {
	return todoService.todoRepository.GetTodosByTitle(title)
}

// UpdateTodo implements ITodoService.
func (todoService *TodoService) UpdateTodo(id int, updateTodoRequest request.UpdateTodoRequest) error {
	return todoService.todoRepository.UpdateTodo(id, updateTodoRequest)
}

func New(todoRepository persistence.ITodoRepository) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}
