package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/reedkihaddi/REST-API/models"
)

type DB struct {
	*sql.DB
}

func New(db *sql.DB) (*DB, error) {
	// Configure any package-level settings
	return &DB{db}, nil
}

func (db *DB) CreateProduct(p *models.Product) error {
	err := db.QueryRow("INSERT INTO products(name,price) VALUES($1,$2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}
