package user

type UserDatabase struct {
	users map[int]*User
}

func NewDatabase() *UserDatabase {
	return &UserDatabase{
		users: make(map[int]*User),
	}
}

func (db *UserDatabase) getAllUsers() ([]*User, error) {
	//s.db.get
	var userList []*User
	for _, u := range db.users {
		userList = append(userList, u)
	}
	return userList, nil
}

func (db *UserDatabase) get(userID int) (*User, error) {
	//s.db.get
	var user *User
	user = db.users[userID]
	return user, nil
}

func (db *UserDatabase) create(user *User) (*User, error) {
	db.users[user.UserID] = user
	return user, nil
}

func (db *UserDatabase) update(user *User) (*User, error) {
	db.users[user.UserID] = user
	return user, nil
}
