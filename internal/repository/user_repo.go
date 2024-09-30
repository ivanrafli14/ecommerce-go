package repository

import (
	"database/sql"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/jmoiron/sqlx"
	"log"
)

type IUserRepository interface {
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	FindByAuthID(authID int) (*entity.User, error)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) CreateUser(user *entity.User) error {
	query := `INSERT INTO users (name, date_of_birth, phone_number, gender,address,image_url,auth_id )
			VALUES  (:name, :date_of_birth, :phone_number, :gender, :address, :image_url, :auth_id)`

	_, err := r.db.NamedExec(query, user)
	if err != nil {
		log.Println(err)
		return response.ErrRepository
	}
	return nil
}

func (r *UserRepository) UpdateUser(user *entity.User) error {
	query := `UPDATE users SET name=:name, date_of_birth=:date_of_birth, phone_number=:phone_number, gender=:gender,
            address=:address, image_url=:image_url WHERE auth_id=:auth_id`

	_, err := r.db.NamedExec(query, user)
	if err != nil {
		log.Println(err.Error())
		return response.ErrRepository
	}
	return nil
}

func (r *UserRepository) FindByAuthID(authID int) (*entity.User, error) {
	var user entity.User
	query := `SELECT id,name, date_of_birth, phone_number, gender,address,image_url,auth_id FROM users 
            WHERE auth_id=$1`

	err := r.db.QueryRowx(query, authID).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrUserNotFound
		}
		return nil, response.ErrorRepository
	}
	return &user, nil

}
