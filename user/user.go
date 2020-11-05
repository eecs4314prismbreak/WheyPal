package user

type userService struct {
	db UserRepo
}

func NewService() UserService {
	return &userService{
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
