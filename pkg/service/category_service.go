package service

import (
	"github.com/khdiyz/book-store/model"
	"github.com/khdiyz/book-store/pkg/repository"
	"github.com/khdiyz/book-store/utils/logger"
)

type CategoryService struct {
	repo *repository.Repository
	log  *logger.Logger
}

func NewCategoryService(repo *repository.Repository, log *logger.Logger) *CategoryService {
	return &CategoryService{repo: repo, log: log}
}

func (s *CategoryService) Create(request model.CreateCategory) (id int, err error) {
	return s.repo.CategoryI.Create(request)
}

func (s *CategoryService) GetList(request model.CategoryListRequest) (categories []model.Category, err error) {
	return s.repo.CategoryI.GetList(request)
}

func (s *CategoryService) GetByID(request model.GetCategoryRequest) (category model.Category, err error) {
	return s.repo.CategoryI.GetByID(request)
}

func (s *CategoryService) Update(request model.UpdateCategory) error {
	return s.repo.CategoryI.Update(request)
}

func (s *CategoryService) Delete(request model.GetCategoryRequest) error {
	return s.repo.CategoryI.Delete(request)
}
