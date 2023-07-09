package product

// Repository provides methods to manage Product objects.
type Repository interface {
	Add(p *Product) error
	Delete(id int64) error
	List(userID string) ([]Product, error)
	Product(id int64, userID string) (*Product, error)
}
