package bcrypt

import "golang.org/x/crypto/bcrypt"

type Interface interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
}

type bcryptImp struct {
	cost int
}

func Init() Interface {
	return &bcryptImp{cost: bcrypt.DefaultCost}
}

func (b *bcryptImp) GenerateFromPassword(password string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", nil
	}
	return string(passwordHashed), nil
}

func (b *bcryptImp) CompareHashAndPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}
	return nil
}
