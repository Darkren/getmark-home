package auth

// Service provides method to authorize and validate user tokens.
type Service interface {
	Auth(login string) (*Token, error)
	ValidateToken(token *Token) error
}
