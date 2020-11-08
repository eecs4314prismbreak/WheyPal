package user

import (
	"database/sql"
	database "github.com/eecs4314prismbreak/WheyPal/database"
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

func NewDatabase(dbname string) UserRepo {
	return &userRepo{
		connector: database.LoadPGDB(),
	}
}

func (db *userRepo) getAllUsers() ([]*User, error) {
	var userList []*User
	// for _, u := range db.users {
	// 	userList = append(userList, u)
	// }

	sqlStatement := `SELECT * FROM users;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	fmt.Printf("Cols: %s", cols)

	for rows.Next() {
		var u *User
		if err := rows.Scan(&u.userid, &u.username, &u.password, &u.email); err != nil {
			panic(err)
		}
		append(userList, u)
	}

	return userList, nil
}

func (db *userRepo) get(userID int) (*User, error) {
	//s.db.get
	var user *User
	user = db.users[userID]
	return user, nil
}

func (db *userRepo) create(user *User) (*User, error) {
	db.users[user.UserID] = user
	return user, nil
}

func (db *userRepo) update(user *User) (*User, error) {
	notUpdatedUser, _ := db.users[user.UserID]

	//if a user did not have a name or other detail specified to update, keep the old detail
	//@Amer, this should be able to be handled through pgx by not updating if it is nil/""
	if user.Name == "" {
		user.Name = notUpdatedUser.Name
	}

	db.users[user.UserID] = user
	return user, nil
}
