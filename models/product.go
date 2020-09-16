package models

//Product is the model for product.
type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}
