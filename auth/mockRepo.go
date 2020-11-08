package auth

import (
	"errors"
)

type authMockRepo struct {
	LoginRepo map[int]*Login
	TokenRepo map[int]*StoredToken
}

func NewAuthMockRepo() AuthRepo {
	return &authMockRepo{
		LoginRepo: make(map[int]*Login),
		TokenRepo: make(map[int]*StoredToken),
	}
}

//database actions

func (r *authMockRepo) getLogin(email string) (*Login, error) {
	for userID, login := range r.LoginRepo {
		if login.Email == email {
			return r.LoginRepo[userID], nil
		}
	}

	return nil, errors.New("Could not find or retirieve user of given email")
}

func (r *authMockRepo) update(login *Login) (bool, error) {
	r.LoginRepo[login.UserID] = login
	return true, nil
}

func (r *authMockRepo) create(login *Login) (*Login, error) {
	// fmt.Println("Login being created | ", login.Email)
	r.LoginRepo[login.UserID] = login
	return login, nil
}

func (r *authMockRepo) storeToken(userID int, jwt *StoredToken) (*StoredToken, error) {
	r.TokenRepo[userID] = jwt

	return jwt, nil
}

func (r *authMockRepo) retrieveToken(userID int) (*StoredToken, error) {
	jwt, ok := r.TokenRepo[userID]
	if !ok {
		return nil, errors.New("Could not find or retireve token")
	}
	return jwt, nil
}
