package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// Token contains JWT token and provides methods for convenient access to its claims.
type Token struct {
	Token string

	jwt    *jwt.Token
	claims jwt.MapClaims
}

// NewToken constructs new token parsing incoming string as JWT.
func NewToken(tokenStr string) (*Token, error) {
	jwt, claims, err := parseToken(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("parseToken: %w", err)
	}

	return &Token{
		Token:  tokenStr,
		jwt:    jwt,
		claims: claims,
	}, nil
}

// Login gets login from token claims.
func (t *Token) Login() (string, error) {
	const loginClaimName = "login"

	if t.claims == nil {
		return "", fmt.Errorf("token is not properly initialized")
	}

	if _, ok := t.claims[loginClaimName]; !ok {
		return "", fmt.Errorf("token doesn't contain login")
	}

	login, ok := t.claims[loginClaimName].(string)
	if !ok {
		return "", fmt.Errorf("login is of wrong type")
	}

	return login, nil
}

// parseToken parses token as JWT.
func parseToken(token string) (*jwt.Token, jwt.MapClaims, error) {
	parser := jwt.NewParser()

	claims := jwt.MapClaims{}
	jwtToken, _, err := parser.ParseUnverified(token, claims)
	if err != nil {
		return nil, nil, fmt.Errorf("parser.ParseUnverified: %w", err)
	}

	return jwtToken, claims, nil
}
