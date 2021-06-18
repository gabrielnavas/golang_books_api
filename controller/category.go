package controller

import (
	"books_api/entity"
	servicecategory "books_api/service"
	"encoding/json"
	"errors"
	"fmt"
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
	categoryService servicecategory.InsertOneCategory
}

func NewCategoryRepository(categoryService servicecategory.InsertOneCategory) *categoryControllerImpl {
	return &categoryControllerImpl{categoryService}
}

func (c *categoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	category := entity.Category{}

	_, err := getIP(r)
	if err != nil {
		fmt.Print(err)
	}

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil && err == io.EOF {
		err = errors.New("no data on body found")
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	category_created, err := c.categoryService.InsertOne(category)
	if err != nil {
		if err == servicecategory.ErrRepositoryError {
			responseJsonErr(&w, err, http.StatusInternalServerError)
			return
		}
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	responseCreated(&w, category_created)
}
