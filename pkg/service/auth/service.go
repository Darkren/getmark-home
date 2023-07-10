package auth

//go:generate mockery --name Service --filename mock_service.go --output ./ --inpackage

// Service provides method to authorize and validate user tokens.
type Service interface {
	Auth(login string) (*Token, error)
	ValidateToken(token *Token) error
}
