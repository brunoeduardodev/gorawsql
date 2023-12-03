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
	Id    string
	Name  string
	Price int
}

type ProductRepository interface {
	Create(input CreateProductInput) Product
	Update(id string, input UpdateProductInput) Product
	Delete(id string) bool
	List() []Product
}
