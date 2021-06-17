package controller

import (
	"books_api/entity"
	"books_api/service"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

var (
	ErrNoDataBody = errors.New("EOF")
)

type categoryControllerImpl struct {
	categoryService service.CategoryService
}

func NewCategoryRepository(categoryService service.CategoryService) *categoryControllerImpl {
	return &categoryControllerImpl{categoryService}
}

func (c *categoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	category := entity.Category{}

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil && err == io.EOF {
		err = errors.New("no data on body found")
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	category_created, err := c.categoryService.InsertOne(category)
	if err != nil {
		if err == service.ErrRepositoryError {
			responseJsonErr(&w, err, http.StatusInternalServerError)
			return
		}
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	responseCreated(&w, category_created)
}
