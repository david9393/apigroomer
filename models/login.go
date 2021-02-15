package models

import "github.com/dgrijalva/jwt-go"

// Login .
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claim .
type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
