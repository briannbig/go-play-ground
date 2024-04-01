package main

import (
	"fmt"
	"gorm-concept/service"
)

func main() {
	service := service.New()

	service.CreateProduct("Foot ball", 878)

	products := service.GetAllProducts()

	fmt.Println(products)

	service.UpdateProduct(1, "Chess board", 4028)

	fmt.Println("After updating...")
	products = service.GetAllProducts()
	fmt.Println(products)
	
}
