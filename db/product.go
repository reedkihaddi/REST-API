package database

import (
	"database/sql"
	"errors"

	//pq is the PostgreSQL driver.
	"github.com/lib/pq"
	"github.com/reedkihaddi/REST-API/logging"
	"github.com/reedkihaddi/REST-API/models"
)

//DB struct for the sql connection.
type DB struct {
	db *sql.DB
}

//New passes the sql connection to the DB struct.
func New(db *sql.DB) (*DB, error) {
	// Configure any package-level settings
	return &DB{db}, nil
}

//CreateProduct inserts a product into the database.
func (db *DB) CreateProduct(p *models.Product) error {
	//var s string
	_, err := db.db.Exec("INSERT INTO products(id,name,price) VALUES($1,$2,$3)",
		p.ID, p.Name, p.Price)
	if err, ok := err.(*pq.Error); ok {
		switch {
		case err.Code.Name() == "unique_violation":
			return errors.New("already found a record in database with associated ID")
		default:
			return err
		}
	}
	return nil
}

//GetProduct fetches the product from the database.
func (db *DB) GetProduct(p *models.Product) error {
	err := db.db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

//UpdateProduct updates the product in the database.
func (db *DB) UpdateProduct(p *models.Product) error {
	_, err := db.db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
		p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteProduct deletes the product from the database.
func (db *DB) DeleteProduct(p *models.Product) error {
	_, err := db.db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	if err != nil {
		return err
	}
	return nil
}

//ListProducts lists all the products from the database.
func (db *DB) ListProducts() ([]*models.Product, error) {

	rows, err := db.db.Query("SELECT id,name,price FROM products")
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

//Check if DB connection is alive.
func (db *DB) Check() bool {
	err := db.db.Ping()
	if err != nil {
		logging.Log.Error("Database connection not alive.")
		return false

	}
	logging.Log.Info("Database connection still alive.")
	return true
}
