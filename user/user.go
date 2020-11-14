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
	UserID    int    `json:"userID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthdate string `json:"birthdate"`
	Location  string `json:"location"`
	Interest  string `json:"interest"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}

type MatchStatus string

const (
	StatusAccept   MatchStatus = "accepted"
	StatusDecline  MatchStatus = "declined"
	StatusPendingA MatchStatus = "pendingUserA"
	StatusNotFound MatchStatus = "notFound"
)
