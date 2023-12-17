package repository

import (
	"todo_api/model"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	GetByID(id int) (model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
	Update(id int, todo model.Todo) (model.Todo, error)
	Delete(id int) error
}

type InMemoryTodoRepository struct {
	todos  map[int]model.Todo
	autoID int
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos:  make(map[int]model.Todo),
		autoID: 1,
	}
}

func (r *InMemoryTodoRepository) GetAll() ([]model.Todo, error) {
	result := make([]model.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		result = append(result, todo)
	}
	return result, nil
}

func (r *InMemoryTodoRepository) GetByID(id int) (model.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return model.Todo{}, model.NOT_FOUND
	}
	return todo, nil
}

func (r *InMemoryTodoRepository) Create(todo model.Todo) (model.Todo, error) {
	todo.Id = r.autoID
	r.autoID++
	r.todos[todo.Id] = todo
	return todo, nil
}

func (r *InMemoryTodoRepository) Update(id int, todo model.Todo) (model.Todo, error) {
	_, ok := r.todos[id]
	if !ok {
		return model.Todo{}, model.NOT_FOUND
	}
	todo.Id = id
	r.todos[id] = todo
	return todo, nil
}

func (r *InMemoryTodoRepository) Delete(id int) error {
	_, ok := r.todos[id]
	if !ok {
		return model.NOT_FOUND
	}
	delete(r.todos, id)
	return nil
}
