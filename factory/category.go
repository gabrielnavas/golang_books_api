package factorycontroller

import (
	"books_api/controller"
	"books_api/repository"
	getall_category "books_api/service/get_all_category"
	"books_api/service/insert_category"
	"database/sql"
)

func MakeCategoryController(db *sql.DB) controller.CategoryController {
	category_repository := repository.NewCategoryRepository(db)
	insert_one_category_service := insert_category.NewInsertOneCategory(category_repository)
	get_all_category_service := getall_category.NewGetAllCategory(category_repository)
	category_controller := controller.NewCategoryRepository(
		insert_one_category_service,
		get_all_category_service,
	)
	return category_controller
}
