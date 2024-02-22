package model

import (
	"time"
)

type Category struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type CreateCategory struct {
	Name string
}

type GetCategoryRequest struct {
	ID int
}

type CategoryListRequest struct {
	SearchKey string
}

type UpdateCategory struct {
	ID   int
	Name string
}
