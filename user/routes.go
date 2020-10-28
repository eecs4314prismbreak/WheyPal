package user

import "math/rand"

func (s *UserService) Create(user *User) (*User, error) {
	//user given random id lol
	user.UserID = rand.Intn(100000) + 1000000

	resp, _ := s.db.create(user)
	return resp, nil
}

func (s *UserService) Update(user *User) (*User, error) {
	resp, _ := s.db.update(user)
	return resp, nil
}

func (s *UserService) GetAllUsers() ([]*User, error) {
	resp, _ := s.db.getAllUsers()
	return resp, nil
}

func (s *UserService) Get(userID int) (*User, error) {
	resp, _ := s.db.get(userID)
	return resp, nil
}
