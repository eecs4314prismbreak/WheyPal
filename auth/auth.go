package auth

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
)

type authService struct {
	Repo AuthRepo
}

func NewService(redisAddr string) AuthService {
	return &authService{
		Repo: NewAuthRepo(redisAddr),
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

func (t *StoredToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}
