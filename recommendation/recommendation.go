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

type MatchStatus string

const (
	StatusAccept   MatchStatus = "accepted"
	StatusDecline  MatchStatus = "declined"
	StatusPendingA MatchStatus = "pendingUserA"
	StatusPendingB MatchStatus = "pendingUserB"
	StatusNotFound MatchStatus = "notFound"
)
