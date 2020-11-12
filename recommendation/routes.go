package recommendation

import (
	"log"

	"github.com/eecs4314prismbreak/WheyPal/user"
)

type RecommendationService interface {
	HandleRecommendationResponse(msg *RecommendationMessage) (RecommendationResponse, error)
	GetRecommendations(userID int) ([]*user.User, error)
}

func (s *recommendationService) GetRecommendations(userID int) ([]*user.User, error) {
	userList, err := s.db.getRecommendations(userID)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return userList, nil
}

func (s *recommendationService) HandleRecommendationResponse(msg *RecommendationMessage) (RecommendationResponse, error) {
	// TODO: Implement the following:
	// user will see targetuser and will do positive or negative response
	// Do either:
	// 1. submit to DB to do monoMatch ( not implemented ) (sqldb's matchrequest.status = "pendingUserB")
	// 2. submit to DB to decline the match ( not implemented ) (sqldb's matchrequest.status = "declined")
	//
	// If both user and targetuser have monomatches against each other,
	// 1. check if either of them declined another. Then just remove both matches
	// 2. if both sent positive response, (users' matchrequest.status against targetuser = "pendingUserB" &&
	//									   targetuser's matchrequest.status against user = "pendingUserB")
	//    then remove both response and insert a new one with usera & userb & status= "accepted"

	// userid := msg.UserID1
	// targetid := msg.UserID2

	// status = 'accepted' OR status = 'declined' OR status = 'pendingUserA' OR status = 'pendingUserB'
	resp, err := s.db.HandleRecommendationResponse(msg)
	return resp, err
}
