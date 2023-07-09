package user

//go:generate mockery --name Repository --filename mock_repository.go --output ./ --inpackage

// Repository provides methods to manage User objects.
type Repository interface {
	Add(u *User) error
}
