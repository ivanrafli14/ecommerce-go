package database

import (
	"fmt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.DatabaseConfig) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Name,
	)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		panic(err)
	}
	return db
}
