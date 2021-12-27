package product

import (
	"context"

	"github.com/LucasRosello/golang-mvc-template/internal/domain"
)

type Service interface {
	Get(ctx context.Context, id int) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, exampleText string) (domain.Product, error)
	Update(ctx context.Context, exampleText string) (domain.Product, error)
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

func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, exampleText string) (domain.Product, error) {

	product := domain.Product{
		ExampleText: exampleText,
	}

	_, err := s.repository.Save(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Update(ctx context.Context, exampleText string) (domain.Product, error) {

	product := domain.Product{
		ExampleText: exampleText,
	}

	err := s.repository.Update(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
