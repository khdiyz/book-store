package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khdiyz/book-store/model"
	"github.com/khdiyz/book-store/utils/logger"
)

type Repository struct {
	CategoryI
}

func NewRepository(db *sqlx.DB, log *logger.Logger) *Repository {
	return &Repository{
		CategoryI: NewCategoryRepo(db, log),
	}
}

type CategoryI interface {
	Create(request model.CreateCategory) (id int, err error)
	GetList(request model.CategoryListRequest) (categories []model.Category, err error)
	GetByID(request model.GetCategoryRequest) (category model.Category, err error)
	Update(request model.UpdateCategory) error
	Delete(request model.GetCategoryRequest) error
}
