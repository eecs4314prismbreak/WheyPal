package auth

import (
	"database/sql"
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
	LoginRepo *sql.DB
	TokenRepo *redis.Client
}

func NewAuthRepo() AuthRepo {
	return &authRepo{
		LoginRepo: LoadPGDB(),
		TokenRepo: LoadRedis(),
	}
}

//database actions

func (r *authRepo) getLogin(email string) (*Login, error) {

	if email == "" {
		return nil, errors.New("EMAIL NULL")
	}
	sqlStatement := `SELECT * FROM logins WHERE email=$1;`
	row := r.LoginRepo.QueryRow(sqlStatement, email)

	// IWAACCT
	l := &Login{}
	if err := row.Scan(&l.UserID, &l.Email, &l.Password); err != nil {
		return nil, errors.New("Could not find or retrieve user of given email")
	}
	return l, nil
}

func (r *authRepo) getLoginFromID(userID int) (*Login, error) {
	sqlStatement := `SELECT * FROM logins WHERE userId=$1;`
	row := r.LoginRepo.QueryRow(sqlStatement, userID)

	// IWAACCT
	l := &Login{}
	if err := row.Scan(&l.UserID, &l.Email, &l.Password); err != nil {
		return nil, errors.New("Could not find or retirieve user of given userID")
	}
	return l, nil
}

func (r *authRepo) update(login *Login) (bool, error) {
	oldLogin, _ := r.getLoginFromID(login.UserID)
	newLogin := &Login{}

	if login.Email == "" {
		newLogin.Email = oldLogin.Email
	} else {
		newLogin.Email = login.Email
	}

	if login.Password == "" {
		newLogin.Password = oldLogin.Password
	} else {
		newLogin.Password = login.Password
	}

	// Insert new user into database
	sqlStatement := `
		UPDATE logins
		SET email=$1, hashedpass = $2
		WHERE userid = $3;`

	_, err := r.LoginRepo.Exec(sqlStatement, newLogin.Email, newLogin.Password, login.UserID)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *authRepo) create(login *Login) (*Login, error) {
	sqlStatement := `INSERT INTO logins (email, hashedPass) VALUES ($1, $2) RETURNING userid;`
	lastInsertID := 0
	row := r.LoginRepo.QueryRow(sqlStatement, login.Email, login.Password).Scan(&lastInsertID)

	if row == sql.ErrNoRows {
		return nil, errors.New("No login row created")
	}

	login.UserID = lastInsertID
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
