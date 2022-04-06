package controller

//go:generate mockgen -package controller -destination mocks/mock_TokenGenerator.go . TokenGenerator

import jwt "github.com/golang-jwt/jwt"

type TokenGenerator interface {
	Generate(method jwt.SigningMethod, claims jwt.Claims) (string, error)
	// NewWithClaims(method jwt.SigningMethod, claims jwt.Claims) Token
}

type tokenGenerator struct {
	secret []byte
}

func (t *tokenGenerator) Generate(method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(method, claims).SignedString(t.secret)
}

// func (t *tokenGenerator) NewWithClaims(method jwt.SigningMethod, claims jwt.Claims) Token {
// 	return jwt.NewWithClaims(method, claims)
// }

// type Token interface {
// 	SignedString(key interface{}) (string, error)
// }

// type token struct {
// 	jwt.Token
// }

// func (t *token) SignedString(key interface{}) (string, error) {
// 	return t.SignedString(key)
// }
