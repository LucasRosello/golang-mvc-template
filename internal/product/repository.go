package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LucasRosello/golang-mvc-template/internal/domain"
)

var METODO_NO_IMPLEMENTADO = errors.New("Metodo no implementado")

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, exampleCode int) bool
	Save(ctx context.Context, w domain.Product) (int, error)
	Update(ctx context.Context, w domain.Product) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {

	return nil, METODO_NO_IMPLEMENTADO
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	return domain.Product{}, METODO_NO_IMPLEMENTADO
}

func (r *repository) Exists(ctx context.Context, exampleCode int) bool {
	return false
}

func (r *repository) Save(ctx context.Context, w domain.Product) (int, error) {
	return 0, METODO_NO_IMPLEMENTADO
}

func (r *repository) Update(ctx context.Context, w domain.Product) error {
	return METODO_NO_IMPLEMENTADO
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return METODO_NO_IMPLEMENTADO
}
