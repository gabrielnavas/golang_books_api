package repository

import (
	"books_api/entity"
	"database/sql"
)

type CategoryRepository interface {
	Insert(c entity.Category) (*entity.Category, error)
	GetByName(name string) (*entity.Category, error)
}

type categoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepositoryImpl {
	return &categoryRepositoryImpl{db}
}

func (cr *categoryRepositoryImpl) GetByName(name string) (*entity.Category, error) {
	sql := `
		SELECT id, name
		FROM book_api.category 
		WHERE name = $1;
	`
	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}
	if !row.Next() {
		return nil, nil
	}

	c := entity.Category{}
	err = row.Scan(&c.ID, &c.Name)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (cr *categoryRepositoryImpl) Insert(c entity.Category) (*entity.Category, error) {
	tx, err := cr.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO book_api.category(name) 
		VALUES ($1) 
		RETURNING id;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(c.Name).Scan(&c.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &c, nil
}
