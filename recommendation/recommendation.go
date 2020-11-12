package recommendation

import "os/user"

type recommendationService struct {
	db RecommendationRepo
}

func NewService() RecommendationService {
	return &recommendationService{
		db: NewDatabase(),
	}
}

type Match struct {
	UserID1                int `json:"userID1"`
	UserID2                int `json:"userID2"`
	RecommendationResponse int `json:"recommendationResponse"`
}

type Recommendations struct {
	Users []*user.User `json:"users"`
}

type MatchRequest struct {
	MatchRequestID int    `json:"matchRequestID"`
	Status         string `json:"status"`
	UserA          int    `json:"userA"`
	UserB          int    `json:"userB"`
}

type RecommendationMessage struct {
	UserID1                int                    `json:"userID1"`
	UserID2                int                    `json:"userID2"`
	RecommendationResponse RecommendationResponse `json:"recommendationResponse"`
}

type RecommendationResponse int

// const PositiveResponse RecommendationResponse = 1
// const NegativeResponse RecommendationResponse = 2

const (
	NilResponse RecommendationResponse = iota
	PositiveResponse
	NegativeResponse
)

type MatchStatus int

const (
	StatusAccept   MatchStatus = iota //0
	StatusDecline                     //1
	StatusPendingA                    //2
	StatusPendingB
	StatusNotFound
)
