package user

type UserService struct {
	db *UserDatabase
}

func NewService() *UserService {
	return &UserService{
		db: NewDatabase(),
	}
}

type User struct {
	UserID int    `json:"userID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}
