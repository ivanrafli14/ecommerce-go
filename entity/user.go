package entity

import "time"

type User struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	PhoneNumber string    `db:"phone_number"`
	Gender      string    `db:"gender"`
	Address     string    `db:"address"`
	ImageUrl    string    `db:"image_url"`
	AuthID      int       `db:"auth_id"`
}

type UserRequest struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
	AuthID      int    `json:"-"`
	Role        string `json:"-"`
}

type UserResponse struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}

type UserBuyer struct {
	ID          int
	Name        string
	PhoneNumber string
	Email       string
}
