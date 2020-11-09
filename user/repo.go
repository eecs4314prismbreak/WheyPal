package user

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

type UserRepo interface {
	getAllUsers() ([]*User, error)
	get(int) (*User, error)
	create(*User) (*User, error)
	update(*User) (*User, error)
}

type userRepo struct {
	connector *sql.DB
}

func NewDatabase() UserRepo {
	return &userRepo{
		connector: LoadPGDB(),
	}
}

func (db *userRepo) getAllUsers() ([]*User, error) {
	var userList []*User

	sqlStatement := `SELECT * FROM users;`
	rows, err := db.connector.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	fmt.Printf("COLS: %s", strings.Join(cols, " "))
	// IWAACCT
	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.UserID, &u.Name, &u.Password, &u.Email); err != nil {
			panic(err)
		}
		userList = append(userList, u)
	}

	return userList, nil
}

func (db *userRepo) get(userID int) (*User, error) {
	//s.db.get

	sqlStatement := `SELECT * FROM users WHERE userid=$1;`
	row := db.connector.QueryRow(sqlStatement, userID)

	// IWAACCT
	u := &User{}
	if err := row.Scan(&u.UserID, &u.Name, &u.Password, &u.Email); err != nil {
		panic(err)
		// return nil, err // Can either return error or just panic here
	}

	return u, nil
}

// Why are we returning the same user we just created?
func (db *userRepo) create(user *User) (*User, error) {
	//s.db.get

	// IWAACCT
	sqlStatement := `INSERT INTO users(username, password, email)
	VALUES ($1, $2, $3);`
	_, err := db.connector.Exec(sqlStatement, user.Name, user.Password, user.Email)

	if err != nil {
		panic(err)
		// return nil, err // Can either return error or just panic here
	}

	return user, nil
}

func (db *userRepo) update(user *User) (*User, error) {
	// IWAACCT

	oldUser, _ := db.get(user.UserID)
	newUser := &User{}

	if user.Name == "" {
		newUser.Name = oldUser.Name
	} else {
		newUser.Name = user.Name
	}

	if user.Email == "" {
		newUser.Email = oldUser.Email
	} else {
		newUser.Email = user.Email
	}

	if user.Password == "" {
		newUser.Password = oldUser.Password
	} else {
		newUser.Password = user.Password
	}

	// Insert new user into database
	sqlStatement := `
	UPDATE users
	SET username = $1, email=$2, password = $3
	WHERE id = $4;`

	_, err := db.connector.Exec(sqlStatement, newUser.Name, newUser.Email, newUser.Password, user.UserID)
	if err != nil {
		panic(err)
	}
	return user, nil
}
