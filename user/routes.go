package user

type UserService interface {
	Create(*User) (*User, error)
	Update(*User) (*User, error)
	GetAllUsers() ([]*User, error)
	Get(int) (*User, error)
	GetMatches(int) ([]*User, error)
	DeleteMatch(int, int) (bool, error)
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
	resp, err := s.db.update(user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) GetAllUsers() ([]*User, error) {
	resp, err := s.db.getAllUsers()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) Get(userID int) (*User, error) {
	resp, err := s.db.get(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) GetMatches(userID int) ([]*User, error) {
	resp, err := s.db.getMatches(userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) DeleteMatch(userID, targetID int) (bool, error) {
	resp, err := s.db.deleteMatch(userID, targetID)
	if err != nil {
		return false, err
	}
	return resp, nil
}
