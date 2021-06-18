package servicecategory

import (
	"books_api/entity"
	"books_api/repository"
	"errors"
	"fmt"
)

type InsertOneCategory interface {
	InsertOne(entity.Category) (*entity.Category, error)
}

var (
	ErrNameAlreadyExists = errors.New("category.name already exists")
	ErrNameIsSmall       = errors.New("category.name is small")
	ErrNameIsLarge       = errors.New("category.name is large")
	ErrRepositoryError   = errors.New("service have a problem")
)

type categoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryServiceImpl {
	return &categoryServiceImpl{categoryRepository}
}

func (c *categoryServiceImpl) InsertOne(category entity.Category) (*entity.Category, error) {
	fmt.Println(category)
	if category.Name == "" || len(category.Name) <= 2 {
		return nil, ErrNameIsSmall
	}
	if len(category.Name) >= 60 {
		return nil, ErrNameIsLarge
	}
	category_fetch, err := c.categoryRepository.GetByName(category.Name)
	if err != nil {
		fmt.Printf("categoryServiceImpl.InsertOne > repository.CategoryRepository.GetByName: %v", err)
		return nil, ErrRepositoryError
	}
	if category_fetch != nil {
		return nil, ErrNameAlreadyExists
	}
	category_created, err := c.categoryRepository.Insert(category)
	if err != nil {
		fmt.Printf("categoryServiceImpl.InsertOne > repository.CategoryRepository.Insert: %v", err)
		return nil, ErrRepositoryError
	}
	return category_created, nil
}
