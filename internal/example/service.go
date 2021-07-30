package example

import (
	"context"

	"github.com/LucasRosello/golang-mvc-template/internal/domain"
)

type Service interface {
	Get(ctx context.Context, id int) (domain.Example, error)
	GetAll(ctx context.Context) ([]domain.Example, error)
	Store(ctx context.Context, exampleText string) (domain.Example, error)
	Update(ctx context.Context, exampleText string) (domain.Example, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Get(ctx context.Context, id int) (domain.Example, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Example, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, exampleText string) (domain.Example, error) {

	example := domain.Example{
		ExampleText: exampleText,
	}

	_, err := s.repository.Save(ctx, example)
	if err != nil {
		return domain.Example{}, err
	}

	return example, nil
}

func (s *service) Update(ctx context.Context, exampleText string) (domain.Example, error) {

	example := domain.Example{
		ExampleText: exampleText,
	}

	err := s.repository.Update(ctx, example)
	if err != nil {
		return domain.Example{}, err
	}

	return example, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
