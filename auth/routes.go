package auth

import (
	"errors"
	"fmt"
	"time"
)

type AuthService interface {
	Login(*LoginRequest) (*AuthResponse, error)
	ValidateToken(int, string) (*Claims, error)
	Create(*Login) (*AuthResponse, error)
	Update(int, *LoginRequest) (bool, error)
	generateToken(int) (*StoredToken, error)
}

func (s *authService) Login(login *LoginRequest) (*AuthResponse, error) {
	//find user based on email
	retrievedLogin, err := s.Repo.getLogin(login.Email)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error retrieving login")
	}

	//validate snubmitted password is the retrievedLogin pw
	validLogin := checkPasswordHash(login.Password, retrievedLogin.Password)
	if !validLogin {
		return nil, errors.New("invalid login")
	}

	//retrieve existing token
	storedToken, err := s.Repo.retrieveToken(retrievedLogin.UserID)
	if err != nil { //token does not exist or error retrieving
		//handle by generating new token
		storedToken, err = s.generateToken(retrievedLogin.UserID)
		//return nil, fmt.Errorf("%v --> %s", err, "auth.routes.Login")
	}

	//check if token is expired
	//currently check is to take the generated token and retrieve claims for expiry
	//correct implementation will have jwt.expiry in redis so no call to get claims
	if err != nil {
		return nil, fmt.Errorf("%v --> Error getting claims --> %s", err, "auth.routes.Login")
	}
	if storedToken.Expiry < time.Now().Unix() { //if expired, generate new jwt
		storedToken, err = s.generateToken(retrievedLogin.UserID)
	}

	token := &AuthResponse{
		ID:          retrievedLogin.UserID,
		Email:       retrievedLogin.Email,
		StoredToken: storedToken,
	}

	return token, nil
}

func (s *authService) ValidateToken(userID int, token string) (*Claims, error) {
	//retrieve token from cache
	retrievedToken, err := s.Repo.retrieveToken(userID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error retrieving token")
	}

	//compare sentToken and retrievedToken
	if token != retrievedToken.Token {
		return nil, errors.New("Token not the same as previously stored; may be expired")
	}

	//get claims
	claims, err := claimsFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error retrieving claims")
	}

	if userID != claims.UserID {
		return nil, errors.New("UserID different than token claim")
	}

	return claims, nil
}

func (s *authService) Create(login *Login) (*AuthResponse, error) {
	login.Password, _ = hashPassword(login.Password)
	retrievedLogin, err := s.Repo.create(login)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error creating login")
	}

	tokenToStore, err := s.generateToken(login.UserID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error generating token")
	}

	token := &AuthResponse{
		ID:          retrievedLogin.UserID,
		Email:       retrievedLogin.Email,
		StoredToken: tokenToStore,
	}

	return token, nil
}

func (s *authService) Update(userID int, loginRequest *LoginRequest) (bool, error) {
	login := &Login{
		UserID:   userID,
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	updateSuccess, err := s.Repo.update(login)
	if err != nil {
		return false, fmt.Errorf("%v --> %s", err, "error creating login")
	}

	return updateSuccess, nil
}

func (s *authService) generateToken(userID int) (*StoredToken, error) {
	generatedToken, err := createJWT(userID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error generating token")
	}

	_, err = s.Repo.storeToken(userID, generatedToken)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error storing token")
	}

	return generatedToken, nil
}
