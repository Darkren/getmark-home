package product

// Repository provides methods to manage Product objects.
type Repository interface {
	Add(p *Product) error
	Delete(barcode string, userID int64) error
	List(userID int64) ([]Product, error)
	Product(barcode string, userID int64) (*Product, error)
}
