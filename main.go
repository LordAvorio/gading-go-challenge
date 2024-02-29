package main

import (
	"fmt"
	"gading-go-challenge/database"
	"gading-go-challenge/models"
	"gading-go-challenge/utils"
	"time"
)

func main() {
	utils.ReadConfig()
	database.StartDB()

	// createProduct()
	// updateProduct()
	// getProductById()
	// createVariant()
	// updateVariantById()
	// deleteVariantById()
	getProductWithVariant()
}

func createProduct(){
	db := database.GetDB()

	product := models.ProductResponse{}

	//INSERT INSIDE THIS STRUCT
	productRequest := models.Product{
		Name : "Iphone 13 Pro Max",
	}

	result, err := db.Exec(database.ProductStatements["CreateProduct"], productRequest.Name)

	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	err = db.QueryRow(database.ProductStatements["GetProductByID"], lastInsertID).Scan(&product.ID, &product.Name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Product Data: %+v\n", product)
}

func updateProduct(){
	db := database.GetDB()

	//INSERT INSIDE THIS STRUCT
	productUpdateRequest := models.Product{
		ID: 3,
		Name : "Iphone 14 Pro Max",
		UpdatedAt: time.Now(),
	}

	result, err := db.Exec(database.ProductStatements["UpdateProduct"], productUpdateRequest.Name, productUpdateRequest.UpdatedAt, productUpdateRequest.ID)	
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated product amount: ", count)
}

func getProductById(){
	db := database.GetDB()

	product := models.ProductResponse{}
	//INSERT INSIDE THIS STRUCT
	productRequest := models.Product{
		ID : 1,
	}

	err := db.QueryRow(database.ProductStatements["GetProductByID"], productRequest.ID).Scan(&product.ID, &product.Name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product Data: %+v\n", product)
}

func createVariant(){
	db := database.GetDB()

	variant := models.VariantResponse{}

	//INSERT INSIDE THIS STRUCT
	variantRequest := models.Variant{
		VariantName: "Jade",
		Quantity: 50,
		ProductID: 3,
	}

	result, err := db.Exec(database.VariantStatements["CreateVariant"], variantRequest.VariantName, variantRequest.Quantity, variantRequest.ProductID)

	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	err = db.QueryRow(database.VariantStatements["GetVariantByID"], lastInsertID).Scan(
		&variant.ID,
		&variant.VariantName,
		&variant.Quantity,
		&variant.ProductID, 
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Variant Data: %+v\n", variant)
}

func updateVariantById(){
	db := database.GetDB()

	//INSERT INSIDE THIS STRUCT
	variantUpdateRequest := models.Variant{
		ID: 1,
		VariantName : "Magenta",
		Quantity: 100,
		ProductID: 1,
		UpdatedAt: time.Now(),
	}

	result, err := db.Exec(database.VariantStatements["UpdateVariant"], variantUpdateRequest.VariantName, variantUpdateRequest.Quantity, variantUpdateRequest.ProductID, variantUpdateRequest.UpdatedAt, variantUpdateRequest.ID)	
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated variant amount: ", count)
}

func deleteVariantById(){
	db := database.GetDB()

	//INSERT INSIDE THIS STRUCT
	variantUpdateRequest := models.Variant{
		ID: 1,
		UpdatedAt: time.Now(),
	}

	result, err := db.Exec(database.VariantStatements["DeleteVariant"], variantUpdateRequest.UpdatedAt, variantUpdateRequest.ID)	
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted variant amount: ", count)
}

func getProductWithVariant(){
	db := database.GetDB()

	listOfResults := []models.ProductWithVariantResponse{}

	rows, err := db.Query(database.ProductStatements["GetProductWithVariant"])

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		result := models.ProductWithVariantResponse{}

		err = rows.Scan(&result.ProductName, &result.VariantName, &result.Quantity)
		if err != nil {
			panic(err)
		}

		listOfResults = append(listOfResults, result)

	}

	fmt.Println("Products with variant result : ", listOfResults)
}