package todo

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"todo/internal/entity"
	"todo/internal/repository"
)

type todoRepository struct {
	db *sql.DB
}

var _ repository.TodoRepository = (*todoRepository)(nil)

func NewTodoRepository(storagePath string) (*todoRepository, error) {
	const op = "todo.internal.repository.todo.NewTodoRepository"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS todo(
		    id INTEGER PRIMARY KEY,
		    title TEXT NOT NULL,
		    completed BOOLEAN
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &todoRepository{db: db}, nil
}

func (r *todoRepository) Create(title string) error {
	const op = "todo.internal.repository.todo.Create"

	stmt, err := r.db.Prepare("INSERT INTO todo(title, completed) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(title, 0)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *todoRepository) Get() ([]entity.Todo, error) {
	const op = "todo.internal.repository.todo.Get"

	rows, err := r.db.Query("SELECT * FROM todo ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var todos []entity.Todo
	var todo entity.Todo

	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		todos = append(todos, todo)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return todos, nil
}

func (r *todoRepository) UpdateStatus(id int) error {
	const op = "todo.internal.repository.todo.UpdateStatus"

	_, err := r.db.Exec("UPDATE todo SET completed = NOT completed WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *todoRepository) UpdateTitle(id int, title string) error {
	const op = "todo.internal.repository.todo.UpdateTitle"

	_, err := r.db.Exec("UPDATE todo SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *todoRepository) Delete(id int) error {
	const op = "todo.repositories.todo_repository.Delete"

	_, err := r.db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
