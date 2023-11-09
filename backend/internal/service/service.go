package service

import (
	"todo/internal/entity"
)

type TodoService interface {
	Get() ([]entity.Todo, error)
	Create(title string) error
	UpdateStatus(id int) error
	UpdateTitle(id int, title string) error
	Delete(id int) error
}
