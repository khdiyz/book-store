package service

import (
	"github.com/khdiyz/book-store/model"
	"github.com/khdiyz/book-store/pkg/repository"
	"github.com/khdiyz/book-store/utils/logger"
)

type Service struct {
	CategoryI
}

func NewService(repos *repository.Repository, log *logger.Logger) *Service {
	return &Service{
		CategoryI: NewCategoryService(repos, log),
	}
}

type CategoryI interface {
	Create(request model.CreateCategory) (id int, err error)
	GetList(request model.CategoryListRequest) (categories []model.Category, err error)
	GetByID(request model.GetCategoryRequest) (category model.Category, err error)
	Update(request model.UpdateCategory) error
	Delete(request model.GetCategoryRequest) error
}
