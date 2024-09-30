package repository

import (
	"database/sql"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/jmoiron/sqlx"
)

type IAuthRepository interface {
	CreateAuth(auth *entity.Auth) error
	GetAuthByEmail(email string) (*entity.Auth, error)
	GetAuthByID(auth int) (*entity.Auth, error)
	UpdateRoleAuth(authID int) error
}

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateAuth(auth *entity.Auth) error {
	query := `INSERT INTO auth (email, password, role) VALUES (:email, :password, :role)`
	_, err := r.db.NamedExec(query, auth)
	if err != nil {
		return response.ErrRepository
	}
	return nil
}

func (r *AuthRepository) GetAuthByEmail(email string) (*entity.Auth, error) {
	query := `SELECT id, email, password, role FROM auth WHERE email = $1`
	var auth entity.Auth

	err := r.db.QueryRowx(query, email).StructScan(&auth)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrEmailNotFound
		}
		return nil, response.ErrRepository
	}

	return &auth, nil
}

func (r *AuthRepository) UpdateRoleAuth(authID int) error {
	query := `UPDATE auth SET role = $1 WHERE id = $2`
	_, err := r.db.Exec(query, "merchant", authID)
	if err != nil {
		return response.ErrRepository
	}
	return nil
}

func (r *AuthRepository) GetAuthByID(authID int) (*entity.Auth, error) {
	var auth entity.Auth
	query := `SELECT id, email, password, role FROM auth WHERE id = $1`
	err := r.db.QueryRowx(query, authID).StructScan(&auth)
	if err != nil {
		return nil, err
	}
	return &auth, nil

}
