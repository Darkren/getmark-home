package user

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

// Add adds a user to the storage.
func (r *PgSQLRepository) Add(u *User) error {
	if u == nil {
		return nil
	}

	if err := r.db.Create(u).Error; err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}

	return nil
}
