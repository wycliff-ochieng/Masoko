package migrate

import "time"

type Product struct {
	ID          int
	Name        string
	Image       string
	Price       float64
	Description string
	Stock       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var Products = []*Product{
	{
		ID:          1,
		Name:        "Gaming Laptop",
		Image:       "/images/laptop.jpg",
		Price:       999.99,
		Description: "High-performance gaming laptop with RTX 3060",
		Stock:       25,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Mechanical Keyboard",
		Image:       "/images/keyboard.jpg",
		Price:       129.99,
		Description: "RGB mechanical keyboard with blue switches",
		Stock:       50,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

func GetProducts() []*Product {
	return Products
}
func CreateProduct(p *Product) {
	p.ID = getNextProduct()
	Products = append(Products, p)
}
func getNextProduct() int {
	lastProduct := Products[len(Products)-1]
	return lastProduct.ID + 1
}
