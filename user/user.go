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
	UserID   int    `json:"userID"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Location string `json:"location"`
	Interest string `json:"interest"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}
