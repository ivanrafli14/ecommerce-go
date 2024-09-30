package entity

type Merchant struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Address     string `db:"address"`
	City        string `db:"city"`
	ImageUrl    string `db:"image_url"`
	AuthID      int    `db:"auth_id"`
}
type MerchantRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ImageUrl    string `json:"image_url"`
	AuthID      int    `json:"-"`
	Role        string `json:"-"`
}
