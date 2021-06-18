package controller

import (
	"books_api/entity"
	getall_category "books_api/service/get_all_category"
	"books_api/service/insert_category"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

var (
	ErrNoDataBody = errors.New("EOF")
)

type categoryControllerImpl struct {
	insertOneCategoryService insert_category.InsertOneCategory
	getAllCategoryService    getall_category.GetAllCategory
}

func NewCategoryRepository(
	insertOneCategoryService insert_category.InsertOneCategory,
	getAllCategoryService getall_category.GetAllCategory,
) *categoryControllerImpl {
	return &categoryControllerImpl{
		insertOneCategoryService,
		getAllCategoryService,
	}
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

	category_created, err := c.insertOneCategoryService.InsertOne(category)
	if err != nil {
		if err == insert_category.ErrRepositoryError {
			responseJsonErr(&w, err, http.StatusInternalServerError)
			return
		}
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	responseCreated(&w, category_created)
}

func (c *categoryControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := c.getAllCategoryService.GetAll()
	if err != nil {
		if err == getall_category.ErrRepositoryError {
			responseJsonErr(&w, err, http.StatusInternalServerError)
			return
		}
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	responseOk(&w, categories)
}
