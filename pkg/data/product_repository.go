package data

import "github.com/Darkren/getmark-home/pkg/data/schema"

//go:generate mockery --name ProductRepository --filename product_repository_mock.go --output ./ --inpackage

// ProductRepository provides methods to manage Product objects.
type ProductRepository interface {
	Add(p *schema.Product) error
	Delete(barcode string, userID int64) error
	List(userID int64) ([]schema.Product, error)
	Product(barcode string, userID int64) (*schema.Product, error)
}
