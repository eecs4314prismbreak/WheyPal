package auth

import (
	"errors"
)

type AuthRepo struct {
	LoginRepo map[int]*Login
	TokenRepo map[int]*StoredToken
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{
		LoginRepo: make(map[int]*Login),
		TokenRepo: make(map[int]*StoredToken),
	}
}

//database actions

func (r *AuthRepo) getLogin(email string) (*Login, error) {
	for userID, login := range r.LoginRepo {
		if login.Email == email {
			return r.LoginRepo[userID], nil
		}
	}

	return nil, errors.New("Could not find or retirieve user of given email")
}

func (r *AuthRepo) update(login *Login) (bool, error) {
	r.LoginRepo[login.UserID] = login
	return true, nil
}

func (r *AuthRepo) create(login *Login) (*Login, error) {
	// fmt.Println("Login being created | ", login.Email)
	r.LoginRepo[login.UserID] = login
	return login, nil
}

func (r *AuthRepo) storeToken(userID int, expiry int64, jwt string) (string, error) {
	r.TokenRepo[userID] = &StoredToken{
		Token:  jwt,
		Expiry: expiry,
	}

	return jwt, nil
}

func (r *AuthRepo) retrieveToken(userID int) (*StoredToken, error) {
	jwt, ok := r.TokenRepo[userID]
	if !ok {
		return nil, errors.New("Could not find or retireve token")
	}
	return jwt, nil
}
