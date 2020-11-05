package auth

import (
	"errors"
)

type AuthRepo interface {
	getLogin(string) (*Login, error)
	update(*Login) (bool, error)
	create(*Login) (*Login, error)
	storeToken(int, int64, string) (string, error)
	retrieveToken(int) (*StoredToken, error)
}

type authRepo struct {
	LoginRepo map[int]*Login
	TokenRepo map[int]*StoredToken
}

func NewAuthRepo() AuthRepo {
	return &authRepo{
		LoginRepo: make(map[int]*Login),
		TokenRepo: make(map[int]*StoredToken),
	}
}

//database actions

func (r *authRepo) getLogin(email string) (*Login, error) {
	for userID, login := range r.LoginRepo {
		if login.Email == email {
			return r.LoginRepo[userID], nil
		}
	}

	return nil, errors.New("Could not find or retirieve user of given email")
}

func (r *authRepo) update(login *Login) (bool, error) {
	r.LoginRepo[login.UserID] = login
	return true, nil
}

func (r *authRepo) create(login *Login) (*Login, error) {
	// fmt.Println("Login being created | ", login.Email)
	r.LoginRepo[login.UserID] = login
	return login, nil
}

func (r *authRepo) storeToken(userID int, expiry int64, jwt string) (string, error) {
	r.TokenRepo[userID] = &StoredToken{
		Token:  jwt,
		Expiry: expiry,
	}

	return jwt, nil
}

func (r *authRepo) retrieveToken(userID int) (*StoredToken, error) {
	jwt, ok := r.TokenRepo[userID]
	if !ok {
		return nil, errors.New("Could not find or retireve token")
	}
	return jwt, nil
}
