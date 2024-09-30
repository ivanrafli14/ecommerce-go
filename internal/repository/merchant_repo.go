package repository

import (
	"database/sql"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/jmoiron/sqlx"
	"log"
)

type IMerchantRepository interface {
	CreateMerchant(merchant *entity.Merchant) error
	GetMerchantByAuthId(authId int) (*entity.Merchant, error)
	UpdateMerchant(merchant *entity.Merchant) error
	GetMerchantByID(merchantID int) (entity.Merchant, error)
}

type MerchantRepository struct {
	db *sqlx.DB
}

func NewMerchantRepository(db *sqlx.DB) IMerchantRepository {
	return &MerchantRepository{
		db: db,
	}
}

func (r *MerchantRepository) CreateMerchant(merchant *entity.Merchant) error {
	query := `INSERT INTO merchants (name, phone_number, address,city,image_url, auth_id) 
    		 VALUES (:name, :phone_number, :address, :city, :image_url, :auth_id)`
	_, err := r.db.NamedExec(query, merchant)

	if err != nil {
		return response.ErrRepository
	}
	return nil
}

func (r *MerchantRepository) GetMerchantByAuthId(authId int) (*entity.Merchant, error) {
	var merchant entity.Merchant
	query := `SELECT id, name, phone_number, address,city,image_url,auth_id FROM merchants WHERE auth_id = $1`
	err := r.db.QueryRowx(query, authId).StructScan(&merchant)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrMerchantNotFound
		}
		return nil, err
	}

	return &merchant, nil
}

func (r *MerchantRepository) UpdateMerchant(merchant *entity.Merchant) error {
	query := `UPDATE merchants set name=:name, phone_number=:phone_number, address=:address, city=:city, image_url=:image_url where auth_id=:auth_id`
	_, err := r.db.NamedExec(query, merchant)

	if err != nil {
		log.Println(err)
		return response.ErrRepository
	}
	return nil
}

func (r *MerchantRepository) GetMerchantByID(merchantID int) (entity.Merchant, error) {
	var merchant entity.Merchant
	query := `SELECT id, name, phone_number, address,city,image_url,auth_id FROM merchants WHERE id=$1`
	err := r.db.QueryRowx(query, merchantID).StructScan(&merchant)
	if err != nil {
		if err == sql.ErrNoRows {
			return merchant, response.ErrMerchantNotFound
		}
		return merchant, err
	}
	return merchant, nil
}
