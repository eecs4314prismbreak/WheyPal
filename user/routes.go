package user

type UserService interface {
	Create(*User) (*User, error)
	Update(*User) (*User, error)
	GetAllUsers() ([]*User, error)
	Get(int) (*User, error)
}

func (s *userService) Create(user *User) (*User, error) {
	//user given random id lol
	// user.UserID = rand.Intn(100000) + 1000000

	resp, err := s.db.create(user)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *userService) Update(user *User) (*User, error) {
	resp, _ := s.db.update(user)
	return resp, nil
}

func (s *userService) GetAllUsers() ([]*User, error) {
	resp, _ := s.db.getAllUsers()
	return resp, nil
}

func (s *userService) Get(userID int) (*User, error) {
	resp, _ := s.db.get(userID)
	return resp, nil
}
