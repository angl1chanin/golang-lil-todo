package todo

import (
	"todo/internal/entity"
	"todo/internal/repository"
	"todo/internal/service"
)

type todoService struct {
	r repository.TodoRepository
}

var _ service.TodoService = (*todoService)(nil)

func NewTodoService(r repository.TodoRepository) service.TodoService {
	return &todoService{r: r}
}

func (s *todoService) Get() ([]entity.Todo, error) {
	return s.r.Get()
}

func (s *todoService) Create(title string) error {
	return s.r.Create(title)
}

func (s *todoService) UpdateStatus(id int) error {
	return nil
}

func (s *todoService) UpdateTitle(id int, title string) error {
	return nil
}

func (s *todoService) Delete(id int) error {
	return nil
}
