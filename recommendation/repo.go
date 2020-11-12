package recommendation

import (
	"database/sql"

	"github.com/eecs4314prismbreak/WheyPal/user"
	_ "github.com/lib/pq"
)

type RecommendationRepo interface {
	getRecommendations(userID int) ([]*user.User, error)
	monoMatchHandle(userID, targetUserID int, resp int) error
	saveMatch(userID, targetUserID int) (*RecommendationResponse, error)
	deleteMonoMatch(userID, targetUserID int) error
	HandleRecommendationResponse(msg *RecommendationMessage) (bool, error)
}

type recommendationRepo struct {
	connector *sql.DB
}

func NewDatabase() RecommendationRepo {
	return &recommendationRepo{
		connector: LoadPGDB(),
	}
}

// Will return true if generating a match, false otherwise
func (r *recommendationRepo) HandleRecommendationResponse(msg *RecommendationMessage) (bool, error) {
	response := msg.RecommendationResponse
	sourceID := msg.UserID1
	targetID := msg.UserID2

	var err error = nil
	if response == PositiveResponse {
		err = r.monoMatchHandle(sourceID, targetID, PosResp)
	} else {
		err = r.monoMatchHandle(sourceID, targetID, NegResp)
	}

	if err != nil {
		return false, err
	}

	// Do a check here for potential matches
	mrListB, err := r.getUserMatchRequestRows(targetID)
	if err != nil {
		return false, err
	}

	// List empty => target user has not swiped on userA yet => do nothing
	if mrListB == nil {
		return false, nil
	}

	// Iterate through Match Request List
	for _, u := range mrListB {
		// From the perspective of every row in mrListB
		// The source user is our target user

		// If we find the requester's target user and it turns out to be our source user
		// we have a match
		if u.UserB == sourceID {
			err = r.createMatch(sourceID, targetID)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func (r *recommendationRepo) createMatch(userA int, userB int) error {
	// Update the userA userB column
	sqlStatement := ` UPDATE matchrequest
	SET status = $1 
	WHERE usera = $2 AND userb = $3;`
	_, err := r.connector.Exec(sqlStatement, STATUS_ACCEPT, userA, userB)
	if err != nil {
		return err
	}

	// Update the userB userA column
	sqlStatement = ` UPDATE matchrequest
	SET status = $1 
	WHERE usera = $2 AND userb = $3;`
	_, err = r.connector.Exec(sqlStatement, STATUS_ACCEPT, userB, userA)
	if err != nil {
		return err
	}

	return nil
}

func (r *recommendationRepo) getUserMatchRequestRows(userID int) ([]*MatchRequest, error) {
	sqlStatement := `SELECT matchrequestid, status, usera, userb FROM matchrequest WHERE usera=$1;`

	// Assume we won't find the user, then .Scan to write values from the query into the variables
	var mrList []*MatchRequest

	rows, err := r.connector.Query(sqlStatement, userID)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, nil
	}

	for rows.Next() {
		mr := &MatchRequest{}
		if err := rows.Scan(&mr.MatchRequestID, &mr.Status, &mr.UserA, &mr.UserB); err != nil {
			return nil, err
		}
		mrList = append(mrList, mr)
	}

	return mrList, nil
}

func (r *recommendationRepo) getRecommendations(userID int) ([]*user.User, error) {
	// Returns the list of Recommendations

	var userList []*user.User

	sqlStatement := `
		SELECT 
			u.userid, u.username, u.birthday, u.location, u.interest
		FROM
			users u
		WHERE
			u.interest IN
				(
					SELECT u.interest
					FROM users u
					WHERE u.userid = $1
				)
			AND
			u.userid != $1
		;
	`
	// rows, err := r.connector.Query(sqlStatement, "pendingUserB", userID)
	rows, err := r.connector.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := &user.User{}
		//11NOV @AMER make sure this here is correct
		if err := rows.Scan(&u.UserID, &u.Name, &u.Birthday, &u.Location, &u.Interest); err != nil {
			return nil, err
		}
		userList = append(userList, u)
	}

	return userList, nil
}

func (r *recommendationRepo) monoMatchHandle(userID, targetUserID int, resp int) error {
	sqlStatement := `INSERT INTO matchrequest( status, usera, userb ) VALUES ( $1, $2, $3 );`

	respString := ""

	if resp == PosResp {
		respString = STATUS_PENDING_B
	} else {
		respString = STATUS_DECLINED
	}
	_, err := r.connector.Exec(sqlStatement, respString, userID, targetUserID)

	if err != nil {
		return err
	}

	return nil
}

func (r *recommendationRepo) deleteMonoMatch(userID, targetUserID int) error {
	// TODO, may not even need it

	// sqlStatement := `
	// 	DELETE FROM
	// 		matchrequest
	// 	WEHRE
	// 		usera = $1 AND
	// 		userb = $2
	// 	;
	// `
	// _, err := db.connector.Exec(sqlStatement, userID, targetUserID)

	// if err != nil {panic(err)}

	// return err

	return nil
}

func (r *recommendationRepo) saveMatch(userID, targetUserID int) (*RecommendationResponse, error) {
	// TODO, may not even need it

	return nil, nil
}
