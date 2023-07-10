package data

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/Darkren/getmark-home/pkg/data/schema"
)

// ProductRepositoryPgSQL implements ProductRepository using PostgreSQL as an internal storage.
type ProductRepositoryPgSQL struct {
	db *gorm.DB
}

// NewProductRepositoryPgSQL constructs an instance of ProductRepositoryPgSQL.
func NewProductRepositoryPgSQL(db *gorm.DB) ProductRepository {
	return &ProductRepositoryPgSQL{
		db: db,
	}
}

// Add adds a Product to the storage.
func (r *ProductRepositoryPgSQL) Add(p *schema.Product) error {
	if p == nil {
		return nil
	}

	if err := r.db.Create(p).Error; err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}

	return nil
}

// Delete deletes a Product from the storage.
// Only deletes a Product with the specified barcode belonging to a User with
// the specified userID.
func (r *ProductRepositoryPgSQL) Delete(barcode string, userID int64) error {
	if err := r.db.Delete(&schema.Product{}, "barcode = $1 AND user_id = $2", barcode, userID).Error; err != nil {
		return fmt.Errorf("db.Delete: %w", err)
	}

	return nil
}

// List returns all Products belonging to a User with the specified userID.
func (r *ProductRepositoryPgSQL) List(userID int64) ([]schema.Product, error) {
	var products []schema.Product
	err := r.db.Table("products").Where("user_id = $1", userID).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("db.Find: %w", err)
	}

	return products, nil
}

// Product returns a product having specified barcode and belonging to a User with
// the specified userID.
func (r *ProductRepositoryPgSQL) Product(barcode string, userID int64) (*schema.Product, error) {
	var product schema.Product
	err := r.db.First(&product, "barcode = $1 AND user_id = $2", barcode, userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("db.First: %w", err)
	}

	return &product, nil
}
