package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgProductRepository struct {
	conn *pgx.Conn
}

func (repository PgProductRepository) FindById(targetId int) (*Product, error) {
	var id int
	var name string
	var price int

	err := repository.conn.QueryRow(context.Background(), "SELECT id, name, price FROM PRODUCTS WHERE ID = $1", targetId).Scan(&id, &name, &price)
	if err != nil {
		return nil, err
	}

	return &Product{Id: id, Name: name, Price: price}, nil
}

func (repository PgProductRepository) Create(input CreateProductInput) (*Product, error) {
	var id int
	var name string
	var price int

	err := repository.conn.QueryRow(context.Background(), "INSERT INTO PRODUCTS (name, price) values ($1, $2) returning id, name, price", input.Name, input.Price).Scan(&id, &name, &price)
	if err != nil {
		return nil, err
	}

	return &Product{Id: id, Name: name, Price: price}, nil
}

func (repository PgProductRepository) List() (*[]Product, error) {
	rows, err := repository.conn.Query(context.Background(), "SELECT id, name, price from PRODUCTS")
	if err != nil {
		return nil, err
	}

	products := []Product{}

	for rows.Next() {
		product := Product{}
		rows.Scan(&product.Id, &product.Name, &product.Price)
		products = append(products, product)
	}

	return &products, nil
}

func (repository PgProductRepository) Update(id int, input UpdateProductInput) (*Product, error) {
	product := Product{}
	err := repository.conn.QueryRow(context.Background(), "UPDATE PRODUCTS SET NAME=$2, price=$3 WHERE ID = $1 returning id, name, price", id, input.Name, input.Price).Scan(&product.Id, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func MakePgProductRepository(conn *pgx.Conn) PgProductRepository {
	return PgProductRepository{conn: conn}
}
