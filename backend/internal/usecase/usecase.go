package usecase

import (
	"todo/internal/entity"
	"todo/internal/service"
)

type UseCase interface {
	Get() ([]entity.Todo, error)
	Create(title string) error
	UpdateStatus(id int) error
	UpdateTitle(id int, title string) error
	Delete(id int) error
}

type useCase struct {
	todoService service.TodoService
}

func NewUseCase(todoService service.TodoService) UseCase {
	return &useCase{
		todoService: todoService,
	}
}

func (s *useCase) Get() ([]entity.Todo, error) {
	return s.todoService.Get()
}

func (s *useCase) Create(title string) error {
	return s.todoService.Create(title)
}

func (s *useCase) UpdateStatus(id int) error {
	return s.todoService.UpdateStatus(id)
}

func (s *useCase) UpdateTitle(id int, title string) error {
	return nil
}

func (s *useCase) Delete(id int) error {
	return s.todoService.Delete(id)
}
