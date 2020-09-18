package models

//Product is the model for product.
type Product struct{
	ID int `json:"id" example:"1"`
	Name string `json:"name" example:"Nvidia RTX3090"`
	Price float64 `json:"price" example:"1499.99"`
}


//HTTPOK is a sample struct for HTTP 200/201.
type HTTPOK struct {
	Result string `json:"result" example:"success"`
}

//HTTPError is a sample struct HTTP response.
type HTTPError struct {
	Error string `json:"error" example:"error message"`
}

