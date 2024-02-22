package repository

import (
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/khdiyz/book-store/model"
	"github.com/khdiyz/book-store/utils/logger"
)

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

type CategoryRepo struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewCategoryRepo(db *sqlx.DB, log *logger.Logger) *CategoryRepo {
	return &CategoryRepo{db: db, log: log}
}

func (r *CategoryRepo) Create(request model.CreateCategory) (id int, err error) {
	query := `
	INSERT INTO category (
		name
	)
	VALUES ($1) RETURNING id;`

	err = r.db.Get(&id, query, request.Name)
	if err != nil {
		r.log.Error(err)
		return id, err
	}

	return id, nil
}

func (r *CategoryRepo) GetList(request model.CategoryListRequest) (categories []model.Category, err error) {
	var (
		filterQuery = ""
		number      = 0
		filterArray []interface{}
	)

	query := `
	SELECT 
		id,
		name,
		created_at
	FROM category
	WHERE 
		is_deleted = 0 `

	if request.SearchKey != "" {
		number++
		filterQuery += ` AND name LIKE '%'||$` + strconv.Itoa(number) + `||'%' COLLATE NOCASE`
		filterArray = append(filterArray, request.SearchKey)
	}

	query += filterQuery

	err = r.db.Select(&categories, query, filterArray...)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepo) GetByID(request model.GetCategoryRequest) (category model.Category, err error) {
	query := `
	SELECT
		id,
		name,
		created_at
	FROM category
	WHERE 
		id = $1
		AND is_deleted = 0;`

	err = r.db.Get(&category, query, request.ID)
	if err != nil {
		r.log.Error(err)
		return category, err
	}

	return category, nil
}

func (r *CategoryRepo) Update(request model.UpdateCategory) error {
	query := `
	UPDATE category
	SET
		name = $2,
		updated_at = NOW()
	WHERE 
		id = $1
		AND is_deleted = 0;`

	row, err := r.db.Exec(query, request.ID, request.Name)
	if err != nil {
		r.log.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.log.Error(err)
		return err
	}
	if rowAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}

func (r *CategoryRepo) Delete(request model.GetCategoryRequest) error {
	query := `
	UPDATE category
	SET
		updated_at = NOW(),
		is_deleted = 1
	WHERE 
		id = $1
		AND is_deleted = 0;`

	row, err := r.db.Exec(query, request.ID)
	if err != nil {
		r.log.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.log.Error(err)
		return err
	}
	if rowAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}
