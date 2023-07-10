package data

import "github.com/Darkren/getmark-home/pkg/data/schema"

//go:generate mockery --name UserRepository --filename user_repository_mock.go --output ./ --inpackage

// UserRepository provides methods to manage User objects.
type UserRepository interface {
	Add(u *schema.User) error
	UserByLogin(login string) (*schema.User, error)
}
