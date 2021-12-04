package services

import (
	"database/sql"

	"github.com/samsul96maarif/rest-api-go/entities"
)

type ProductServiceInterface interface {
	Fetch() (*[]entities.Product, error)
}

type productService struct {
	db *sql.DB
}

func RegisterProductService(db *sql.DB) ProductServiceInterface {
	return &productService{db: db}
}

func (s productService) Fetch() (*[]entities.Product, error) {
	var products []entities.Product
	var product entities.Product

	res, err := s.db.Query(`select * from products`)
	if err != nil {
		panic(err.Error())
	}

	for res.Next() {
		err := res.Scan(&product.ID, &product.Name)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	return &products, nil
}
