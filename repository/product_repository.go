package repository

import (
	"database/sql"
	"fmt"
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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.Connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close() // Use defer para garantir que a consulta seja fechada

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
