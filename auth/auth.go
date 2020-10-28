package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	Repo *AuthRepo
}

func NewService() *AuthService {
	return &AuthService{
		Repo: NewAuthRepo(),
	}
}

type Claims struct {
	UserID int `json:"id"`
	jwt.StandardClaims
}

type StoredToken struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

type Login struct {
	UserID   int    `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"name"`
}

type AuthResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	*StoredToken
}

type AuthRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
