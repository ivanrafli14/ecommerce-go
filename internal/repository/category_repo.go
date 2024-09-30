package repository

import (
	"database/sql"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/jmoiron/sqlx"
	"log"
)

type ICategoryRepository interface {
	GetCategoryNameByID(categoryID int) (string, error)
	GetAllCategories() ([]entity.Category, error)
}

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) ICategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategoryNameByID(categoryID int) (string, error) {
	var categoryName string
	query := `SELECT name FROM categories WHERE id = $1`
	err := r.db.QueryRowx(query, categoryID).Scan(&categoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", response.ErrCategoryIDProductNotFound
		}
		return "", response.ErrRepository
	}
	return categoryName, nil
}

func (r *CategoryRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category

	query := `SELECT id,name FROM categories`
	rows, err := r.db.Queryx(query)
	if err != nil {
		log.Println(err)
		return categories, response.ErrRepository
	}
	defer rows.Close()

	for rows.Next() {
		var category entity.Category
		err = rows.StructScan(&category)
		if err != nil {
			log.Println(err)
			return categories, response.ErrRepository
		}
		categories = append(categories, category)
	}
	return categories, nil
}
