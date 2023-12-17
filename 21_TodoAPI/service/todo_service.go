package service

import (
	"todo_api/model"
	"todo_api/repository"
)

type TodoService struct {
	TodoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) *TodoService {
	return &TodoService{
		TodoRepo: todoRepo,
	}
}

func (s *TodoService) GetAll() ([]model.Todo, error) {
	return s.TodoRepo.GetAll()
}

func (s *TodoService) GetByID(id int) (model.Todo, error) {
	return s.TodoRepo.GetByID(id)
}

func (s *TodoService) Create(todo model.Todo) (model.Todo, error) {
	return s.TodoRepo.Create(todo)
}

func (s *TodoService) Update(id int, todo model.Todo) (model.Todo, error) {
	return s.TodoRepo.Update(id, todo)
}

func (s *TodoService) Delete(id int) error {
	return s.TodoRepo.Delete(id)
}
