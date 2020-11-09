package reccomendation

type RecommendationService interface {
	HandleRecommenatdonResponse(*RecommendationMessage) (RecomendationResponse, error)
}

func (s *recommendationService) HandleRecommenatdonResponse(msg *RecommendationMessage) (RecomendationResponse, error) {
	//recommendationresponse is int = 1 or int = 2
	//1 means positive response (right swipe)
	//2 means negathive response (left swipe)
	return msg.RecomendationResponse, nil
}
