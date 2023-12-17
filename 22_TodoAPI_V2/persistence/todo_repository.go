package persistence

import (
	"TodoAPI/controller/request"
	"TodoAPI/domain"
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ITodoRepository interface {
	CreateTodo(createTodoRequest request.CreateTodoRequst) error
	GetAllTodos() []domain.Todo
	GetTodosByTitle(title string) []domain.Todo
	GetTodosByStatus(status bool) []domain.Todo
	GetTodoById(id int) (domain.Todo, error)
	UpdateTodo(id int, updateTodoRequest request.UpdateTodoRequest) error
	DeleteTodoById(id int) error
}

type TodoRepository struct {
	dbPool *pgxpool.Pool
	ctx    context.Context
}

func New(ctx context.Context, dbpool *pgxpool.Pool) ITodoRepository {
	return &TodoRepository{ctx: ctx, dbPool: dbpool}
}

// CreateTodo implements ITodoRepository.
func (todoRepository *TodoRepository) CreateTodo(createTodoRequest request.CreateTodoRequst) error {
	_, err := todoRepository.dbPool.Exec(todoRepository.ctx,
		"INSERT INTO todos (title,description,status) values ($1,$2,$3)",
		createTodoRequest.Title,
		createTodoRequest.Description,
		createTodoRequest.Status,
	)
	return err
}

// DeleteTodoById implements ITodoRepository.
func (todoRepository *TodoRepository) DeleteTodoById(id int) error {
	_, err := todoRepository.dbPool.Exec(todoRepository.ctx,
		"DELETE FROM todos WHERE id = $1",
		id,
	)
	return err
}

// GetAllTodos implements ITodoRepository.
func (todoRepository *TodoRepository) GetAllTodos() []domain.Todo {
	rows, err := todoRepository.dbPool.Query(todoRepository.ctx, "SELECT * FROM todos")
	if err != nil {
		return []domain.Todo{}
	}
	return extractTodosFromRows(rows)
}

// GetTodoById implements ITodoRepository.
func (todoRepository *TodoRepository) GetTodoById(id int) (domain.Todo, error) {
	row := todoRepository.dbPool.QueryRow(todoRepository.ctx,
		"SELECT * FROM todos where id = $1",
		id,
	)
	return extractTodoFromRow(row)
}

// GetTodosByStatus implements ITodoRepository.
func (todoRepository *TodoRepository) GetTodosByStatus(status bool) []domain.Todo {
	rows, err := todoRepository.dbPool.Query(todoRepository.ctx, "SELECT * FROM todos WHERE status = $1", status)
	if err != nil {
		return []domain.Todo{}
	}
	return extractTodosFromRows(rows)
}

// GetTodosByTitle implements ITodoRepository.
func (todoRepository *TodoRepository) GetTodosByTitle(title string) []domain.Todo {
	rows, err := todoRepository.dbPool.Query(todoRepository.ctx, "SELECT * FROM todos WHERE title LIKE $1", title)
	if err != nil {
		return []domain.Todo{}
	}
	return extractTodosFromRows(rows)
}

// UpdateTodo implements ITodoRepository.
func (todoRepository *TodoRepository) UpdateTodo(id int, updateTodoRequest request.UpdateTodoRequest) error {
	_, err := todoRepository.dbPool.Exec(todoRepository.ctx,
		"UPDATE todos SET title = $1 ,description = $2, status = $3 WHERE id = $4",
		updateTodoRequest.Title,
		updateTodoRequest.Description,
		updateTodoRequest.Status,
		id,
	)
	return err
}

func extractTodosFromRows(rows pgx.Rows) []domain.Todo {
	var list = []domain.Todo{}
	for rows.Next() {
		model, err := extractTodoFromRow(rows)
		if err != nil {
			break
		}
		list = append(list, model)
	}
	return list
}

func extractTodoFromRow(row pgx.Row) (domain.Todo, error) {
	var model = domain.Todo{}
	err := row.Scan(&model.Id, &model.Title, &model.Description, &model.Status)
	if err != nil {
		return model, err
	}
	return model, nil
}
