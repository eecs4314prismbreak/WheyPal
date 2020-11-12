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
	RecommendationResponse int `json:"matchResponse"`
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

const PositiveResponse RecommendationResponse = 1
const NegativeResponse RecommendationResponse = 2

const PosResp = 1
const NegResp = 2

const STATUS_ACCEPT = "accepted"
const STATUS_DECLINED = "declined"
const STATUS_PENDING_A = "pendingUserA"
const STATUS_PENDING_B = "pendingUserB"
const STATUS_NOT_FOUND = "notFound"
