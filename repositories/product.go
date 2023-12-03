package repositories

type CreateProductInput struct {
	Name  string
	Price int
}

type UpdateProductInput struct {
	Name  string
	Price int
}

type Product struct {
	Id    int
	Name  string
	Price int
}

type ProductRepository interface {
	Create(input CreateProductInput) (Product, error)
	Update(id int, input UpdateProductInput) (Product, error)
	Delete(id int) error
	List() ([]Product, error)
	FindById(id int) (Product, error)
}
