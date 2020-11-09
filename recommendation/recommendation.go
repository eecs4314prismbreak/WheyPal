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
	UserID1               int `json:"userID1"`
	UserID2               int `json:"userID2"`
	RecomendationResponse int `json:"matchResponse"`
}

type Recommendations struct {
	Users []*user.User `json:"users"`
}

type RecommendationMessage struct {
	UserID1                int                  `json:"userID1"`
	UserID2                int                  `json:"userID2"`
	RecomendationResponse RecomendationResponse `json:"recomendationResponse"`
}

type RecomendationResponse int

const PositiveResponse RecomendationResponse = 1
const NegativeResponse RecomendationResponse = 2
