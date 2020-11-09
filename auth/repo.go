package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type AuthRepo interface {
	getLogin(string) (*Login, error)
	update(*Login) (bool, error)
	create(*Login) (*Login, error)
	storeToken(int, *StoredToken) (*StoredToken, error)
	retrieveToken(int) (*StoredToken, error)
}

type authRepo struct {
	LoginRepo map[int]*Login
	TokenRepo *redis.Client
}

func NewAuthRepo(redisAddr string) AuthRepo {
	return &authRepo{
		LoginRepo: make(map[int]*Login),
		TokenRepo: redis.NewClient(&redis.Options{
			Addr: redisAddr,
		}),
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

func (r *authRepo) storeToken(userID int, jwt *StoredToken) (*StoredToken, error) {
	expiry := time.Duration((jwt.Expiry - time.Now().Unix()) * 1000000)
	if _, err := r.TokenRepo.HSetNX(string(userID), "token", jwt).Result(); err != nil {
		return nil, fmt.Errorf("create: redis error: %w", err)
	}
	r.TokenRepo.Expire(string(userID), expiry)
	return jwt, nil
}

func (r *authRepo) retrieveToken(userID int) (*StoredToken, error) {
	jwt, err := r.TokenRepo.HGet(string(userID), "hoket").Result()
	if err == redis.Nil || err != nil {
		return nil, errors.New("Could not find or retireve token")

	}

	retrievedToken := &StoredToken{}

	if err := json.Unmarshal([]byte(jwt), &retrievedToken); err != nil {
		return nil, err
	}

	return retrievedToken, nil
}
