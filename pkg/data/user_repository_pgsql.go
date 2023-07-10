package data

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/Darkren/getmark-home/pkg/data/schema"
)

// UserRepositoryPgSQL implements UserRepository using PostgreSQL as an internal storage.
type UserRepositoryPgSQL struct {
	db *gorm.DB
}

// NewUserRepositoryPgSQL constructs an instance of UserRepositoryPgSQL.
func NewUserRepositoryPgSQL(db *gorm.DB) UserRepository {
	return &UserRepositoryPgSQL{
		db: db,
	}
}

// Add adds a User to the storage.
func (r *UserRepositoryPgSQL) Add(u *schema.User) error {
	if u == nil {
		return nil
	}

	if err := r.db.Create(u).Error; err != nil {
		return fmt.Errorf("db.Create: %w", err)
	}

	return nil
}

// UserByLogin gets a User by their login.
func (r *UserRepositoryPgSQL) UserByLogin(login string) (*schema.User, error) {
	var user schema.User
	err := r.db.Table("users").Where("login = $1", login).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("db.First: %w", err)
	}

	return &user, nil
}
