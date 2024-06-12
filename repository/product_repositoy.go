package repository

import (
	"database/sql"
	"go-api/model"
	"log"
)

type ProductRepository struct {
	Connection *sql.DB
}

func NewProductRepository(Connection *sql.DB) ProductRepository {
	return ProductRepository{
		Connection: Connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.Connection.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var productList []model.Product

	for rows.Next() {
		var productObj model.Product
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		productList = append(productList, productObj)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return productList, nil
}
