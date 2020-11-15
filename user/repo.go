package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepo interface {
	getAllUsers() ([]*User, error)
	get(int) (*User, error)
	create(*User) (*User, error)
	update(*User) (*User, error)
	getMatches(int) ([]*User, error)
	deleteMatch(int, int) (bool, error)
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

	// cols, _ := rows.Columns()
	// fmt.Printf("COLS: %s", strings.Join(cols, " "))

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Birthdate, &u.Location, &u.Interest); err != nil {
			return nil, err
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
	if err := row.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Birthdate, &u.Location, &u.Interest); err != nil {
		return nil, err
	}
	return u, nil
}

// Why are we returning the same user we just created?
func (db *userRepo) create(user *User) (*User, error) {

	// IWAACCT
	sqlStatement := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := db.connector.Exec(sqlStatement, user.UserID, user.FirstName, user.LastName, user.Birthdate, user.Location, user.Interest)

	// Put ID into the user

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
	if user.FirstName == "" {
		newUser.FirstName = oldUser.FirstName
	} else {
		newUser.FirstName = user.FirstName
	}

	if user.LastName == "" {
		newUser.LastName = oldUser.LastName
	} else {
		newUser.LastName = user.LastName
	}

	// BIRTHDAY
	if user.Birthdate == "" {
		newUser.Birthdate = oldUser.Birthdate
	} else {
		newUser.Birthdate = user.Birthdate
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
	SET firstName = $1, lastName = $2, birthdate=$3, location=$4, interest=$5
	WHERE userID = $6;`

	_, err = db.connector.Exec(sqlStatement, newUser.FirstName, newUser.LastName, newUser.Birthdate, newUser.Location, newUser.Interest, user.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *userRepo) getMatches(userID int) ([]*User, error) {
	var userList []*User

	sqlStatement := `SELECT
						userid, firstName, lastName, birthdate, location, interest 
					FROM
						matchrequest m 
						RIGHT JOIN
						users u 
						ON m.userb = u.userid 
					WHERE
						m.usera = $1 
						and status = $2;`
	rows, err := db.connector.Query(sqlStatement, userID, StatusAccept)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// cols, _ := rows.Columns()
	// fmt.Printf("COLS: %s\n", strings.Join(cols, " "))

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Birthdate, &u.Location, &u.Interest); err != nil {
			return nil, err
		}
		userList = append(userList, u)
	}

	return userList, nil
}

func (db *userRepo) deleteMatch(userID, targetID int) (bool, error) {
	sqlStatement := `DELETE FROM matchrequest WHERE (usera=$1 AND userb=$2) OR (usera=$2 and userb=$1);`
	_, err := db.connector.Exec(sqlStatement, userID, targetID)
	if err != nil {
		return false, err
	}

	return true, nil
}
