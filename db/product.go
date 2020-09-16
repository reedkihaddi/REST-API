package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/reedkihaddi/REST-API/models"
)

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) (*DB, error) {
	// Configure any package-level settings
	return &DB{db}, nil
}

func (db *DB) CreateProduct(p *models.Product) error {
	err := db.db.QueryRow("INSERT INTO products(name,price) VALUES($1,$2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetProduct(p *models.Product) error {
	err := db.db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateProduct(p *models.Product) error {
	_, err := db.db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
		p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteProduct(p *models.Product) error {
	_, err := db.db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetProducts(start, count int) ([]*models.Product, error) {

	rows, err := db.db.Query("SELECT id,name,price FROM products LIMIT $1 OFFSET $2",
		start, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*models.Product{}

	for rows.Next() {
		p := &models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}

		products = append(products, p)
	}
	return products, nil
}
