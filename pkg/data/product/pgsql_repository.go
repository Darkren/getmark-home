package product

import (
	"fmt"

	"gorm.io/gorm"
)

// PgSQLRepository implements Repository using PostgreSQL as an internal storage.
type PgSQLRepository struct {
	db *gorm.DB
}

// NewPgSQLRepository constructs an instance of PgSQLRepository.
func NewPgSQLRepository(db *gorm.DB) Repository {
	return &PgSQLRepository{
		db: db,
	}
}

// Add adds a Product to the storage.
func (r *PgSQLRepository) Add(p *Product) error {
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
func (r *PgSQLRepository) Delete(barcode string, userID int64) error {
	if err := r.db.Delete(&Product{}, "barcode = $1 AND user_id = $2", barcode, userID).Error; err != nil {
		return fmt.Errorf("db.Delete: %w", err)
	}

	return nil
}

// List returns all Products belonging to a User with the specified userID.
func (r *PgSQLRepository) List(userID int64) ([]Product, error) {
	var products []Product
	err := r.db.Table("products").Where("user_id = $1", userID).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("db.Find: %w", err)
	}

	return products, nil
}

// Product returns a product having specified barcode and belonging to a User with
// the specified userID.
func (r *PgSQLRepository) Product(barcode string, userID int64) (*Product, error) {
	var product Product
	err := r.db.First(&product, "barcode = $1 AND user_id = $2", barcode, userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("db.First: %w", err)
	}

	return &product, nil
}
