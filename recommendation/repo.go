package reccomendation

import "os/user"

type RecommendationRepo interface {
	GetRecommendations(userID int) ([]*user.User, error)
	MonoMatch(userID, targetUserID int) error
	SaveMatch(userID, targetUserID int) (RecomendationResponse, error)
	DeleteMonoMatch(userID, targetUserID int) error
}

type recommendationRepo struct {
}

func NewDatabase() RecommendationRepo {
	return &recommendationRepo{}
}

func (r *recommendationRepo) GetRecommendations(userID int) ([]*user.User, error) {
	return nil, nil
}

func (r *recommendationRepo) MonoMatch(userID, targetUserID int) error {
	return nil
}

func (r *recommendationRepo) DeleteMonoMatch(userID, targetUserID int) error {
	return nil
}

func (r *recommendationRepo) SaveMatch(userID, targetUserID int) (RecomendationResponse, error) {
	return 0, nil
}
