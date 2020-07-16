package data

import "time"

type Product struct {
	ID		int
	Name		string
	Description	string
	Price		float32
	SKU		string
	CreatedAt	string
	UpdatedAt	string
	DeletedAt	string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product {
	&Product{
		ID:				1,
		Name:			"Kenangan Mantan",
		Description:	"Kopi Kenangan Mantan",
		Price:			1.5,
		SKU:			"as81",
		CreatedAt:		time.Now().UTC().String(),
		UpdatedAt:		time.Now().UTC().String(),
		DeletedAt:		time.Now().UTC().String(),
	},
	&Product{
		ID:				2,
		Name:			"Americano",
		Description:	"Kopi hitam",
		Price:			1.2,
		SKU:			"sj1s",
		CreatedAt:		time.Now().UTC().String(),
		UpdatedAt:		time.Now().UTC().String(),
		DeletedAt:		time.Now().UTC().String(),
	},
}
