package repositories

import "fmt"

type CreateProductInput struct {
	Name  string
	Price int
}

type UpdateProductInput struct {
	Name  string
	Price int
}

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p Product) ToString() string {
	return fmt.Sprintf("Product {id:%d, name:%s, price:%d }", p.Id, p.Name, p.Price)
}

type ProductRepository interface {
	Create(input CreateProductInput) (*Product, error)
	Update(id int, input UpdateProductInput) (*Product, error)
	Delete(id int) error
	List() (*[]Product, error)
	FindById(id int) (*Product, error)
}
