package repository

import (
	"todo/internal/entity"
)

type TodoRepository interface {
	Get() ([]entity.Todo, error)
	Create(title string) error
	UpdateStatus(id int) error
	UpdateTitle(id int, title string) error
	Delete(id int) error
}
