package user

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
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

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.UserID, &u.Name, &u.Birthday, &u.Location, &u.Interest); err != nil {
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

	u := &User{}
	if err := row.Scan(&u.UserID, &u.Name, &u.Birthday, &u.Location, &u.Interest); err != nil {
		return nil, err
	}

	return u, nil
}

// Why are we returning the same user we just created?
func (db *userRepo) create(user *User) (*User, error) {

	// IWAACCT
	sqlStatement := `INSERT INTO users(userid, username, birthday, location, interest)
	VALUES ($1, $2, $3, $4, $5);`
	_, err := db.connector.Exec(sqlStatement, user.UserID, user.Name, user.Birthday, user.Location, user.Interest)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *userRepo) update(user *User) (*User, error) {
	// IWAACCT

	oldUser, err := db.get(user.UserID)
	if err != nil {
		return nil, err
	}
	newUser := &User{}

	// NAME
	if user.Name == "" {
		newUser.Name = oldUser.Name
	} else {
		newUser.Name = user.Name
	}

	// BIRTHDAY
	if user.Birthday == "" {
		newUser.Birthday = oldUser.Birthday
	} else {
		newUser.Birthday = user.Birthday
	}
	// LOCATION
	if user.Location == "" {
		newUser.Location = oldUser.Location
	} else {
		newUser.Location = user.Location
	}
	// INTEREST
	if user.Interest == "" {
		newUser.Interest = oldUser.Interest
	} else {
		newUser.Interest = user.Interest
	}

	// Insert new user into database
	sqlStatement := `
	UPDATE users
	SET username = $1, birthday=$2, location=$3, interest=$4
	WHERE id = $5;`

	_, err = db.connector.Exec(sqlStatement, newUser.Name, newUser.Birthday, newUser.Location, newUser.Interest, user.UserID)
	if err != nil {
		panic(err)
	}
	return user, nil
}
