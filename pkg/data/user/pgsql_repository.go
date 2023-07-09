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

// Add adds a User to the storage.
func (r *PgSQLRepository) Add(u *User) error {
	if u == nil {
		return nil
	}

	if err := r.db.Create(u).Error; err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}

	return nil
}

// UserByLogin gets a User by their login.
func (r *PgSQLRepository) UserByLogin(login string) (*User, error) {
	var user User
	err := r.db.Table("users").Where("login = $1", login).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("db.First: %w", err)
	}

	return &user, nil
}
