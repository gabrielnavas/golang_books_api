package entity

import "errors"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	ErrCategoryIdEmpty = errors.New("category.id is empty")

	ErrCategoryNameEmpty = errors.New("category.name is empty")
	ErrCategoryNameSmall = errors.New("category.name is small, min 3 character")
	ErrCategoryNameLarge = errors.New("category.name is large, max 80 character")
)

func (c Category) Validate() error {
	if c.ID <= "" {
		return ErrCategoryIdEmpty
	}
	if c.Name == "" {
		return ErrCategoryNameEmpty
	}
	if len(c.Name) < 3 {
		return ErrCategoryNameLarge
	}
	if len(c.Name) > 80 {
		return ErrCategoryNameLarge
	}
	return nil
}
