package getall_category

import (
	"books_api/entity"
	"books_api/repository"
	"errors"
)

type GetAllCategory interface {
	GetAll() ([]*entity.Category, error)
}

type getAllCategoryImpl struct {
	categoryRepository repository.CategoryRepository
}

var (
	ErrRepositoryError = errors.New("service have a problem")
)

func NewGetAllCategory(categoryRepository repository.CategoryRepository) *getAllCategoryImpl {
	return &getAllCategoryImpl{categoryRepository}
}

func (cr *getAllCategoryImpl) GetAll() ([]*entity.Category, error) {
	categories, err := cr.categoryRepository.GetAll()
	if err != nil {
		return nil, ErrRepositoryError
	}
	return categories, err
}
