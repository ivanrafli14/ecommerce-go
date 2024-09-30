package repository

import (
	"database/sql"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/jmoiron/sqlx"
	"log"
)

type IProductRepository interface {
	CreateProduct(product *entity.Product) error
	UpdateProduct(product entity.Product, productID int) error
	GetProductByProductID(productID int, merchantID int) (*entity.ProductDetailResponse, error)
	GetProductBySKU(sku string) (entity.DetailProductSKUResponse, error)
}

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(product *entity.Product) error {

	query := `INSERT INTO products (sku, name, description, price, stock, category_id, image_url,created_at, updated_at, merchant_id) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
			RETURNING id`

	err := r.db.QueryRowx(query, product.SKU, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ImageUrl, product.CreatedAt, product.UpdatedAt, product.MerchantID).Scan(&product.ID)
	if err != nil {
		log.Println(err)
		return response.ErrRepository
	}
	return nil
}

func (r *ProductRepository) UpdateProduct(product entity.Product, productID int) error {

	query := `UPDATE products SET name=$1, description=$2, price=$3, stock=$4, category_id=$5, image_url=$6,
        updated_at=$7 
        FROM merchants
        WHERE merchants.id = products.merchant_id
		AND products.merchant_id = $8
		AND products.id = $9`

	result, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ImageUrl, product.UpdatedAt, product.MerchantID, productID)
	if err != nil {
		log.Println(err)
		return response.ErrRepository
	}

	rowEffected, err := result.RowsAffected()
	if err != nil {
		return response.ErrRepository
	}
	if rowEffected == 0 {
		return response.ErrProductNotFound
	}
	return nil
}

func (r *ProductRepository) GetProductByProductID(productID int, merchantID int) (*entity.ProductDetailResponse, error) {
	var product entity.ProductDetailResponse
	query := `SELECT p.id, sku,p.name,description,price,stock,p.category_id, categories.name AS category,p.image_url,created_at,updated_at, merchant_id FROM products AS p
            	INNER JOIN categories ON categories.id = p.category_id
            	WHERE p.id=$1`

	var args []any
	args = append(args, productID)

	if merchantID > 0 {
		query += ` AND p.merchant_id=$2`
		args = append(args, merchantID)
	}
	err := r.db.QueryRowx(query, args...).StructScan(&product)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrProductNotFound
		}
		log.Println(err)
		return nil, response.ErrRepository
	}
	return &product, nil
}

func (r *ProductRepository) GetProductBySKU(sku string) (entity.DetailProductSKUResponse, error) {
	var product entity.DetailProductSKUResponse
	var merchantID int
	var merchantName, merchantCity string

	query := `SELECT products.id, sku, products.name, description, price, stock, categories.name AS category_name, category_id,
       			merchants.id AS merchant_id, merchants.name as merchant_name, merchants.city AS merchant_city, products.image_url, products.created_at, products.updated_at FROM products
				INNER JOIN merchants ON merchants.id = products.merchant_id
				INNER JOIN categories ON categories.id = products.category_id
				WHERE products.sku = $1`

	row := r.db.QueryRowx(query, sku)

	err := row.Scan(
		&product.ID,
		&product.SKU,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CategoryName,
		&product.CategoryID,
		&merchantID,
		&merchantName,
		&merchantCity,
		&product.ImageUrl,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return product, response.ErrProductNotFound
		}
		return product, response.ErrRepository
	}
	product.Merchant = entity.MerchantDetail{
		ID:   merchantID,
		Name: merchantName,
		City: merchantCity,
	}
	return product, nil
}
