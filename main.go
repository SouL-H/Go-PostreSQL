package main

import (
	"PsqlDb/models"
)

func main() {
	// product := models.Product{
	// 	Title:       "Header",
	// 	Description: "Test header",
	// 	Price:       249.99,
	// }
	// models.InsertProduct(product)//INSERT
	//models.GetProducts() //GET all
	//models.GetProductsByID(1)//Get ID
	data := models.Product{
		ID:    1,
		Title: "Uptate",
		Description: "Update",
		Price: 49.99,
	}
	models.UpdateProcut(data)//UPDATE
}
