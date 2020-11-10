package recommendation

import (
	"database/sql"
	"os/user"
)

type RecommendationRepo interface {
	getRecommendations(userID int) ([]*user.User, error)
	monoMatch(userID, targetUserID int) error
	saveMatch(userID, targetUserID int) (RecomendationResponse, error)
	deleteMonoMatch(userID, targetUserID int) error
}

type recommendationRepo struct {
	connector *sql.DB
}

func NewDatabase() RecommendationRepo {
	return &recommendationRepo{
		connector: LoadPGDB(),
	}
}

func (r *recommendationRepo) getRecommendations(userID, targetUserID int) ([]*user.User, error) {
	// Returns the list of user

	var userList []*user.User

	sqlStatement := `
		SELECT 
			u.userid, u.username, u.password, u.email 
		FROM
			interest i 
				INNER JOIN userinterest ui ON i.interestid = ui.interestid
				INNER JOIN profile p ON p.userid = ui.userid
				INNER JOIN users u ON u.userid = ui.userid
		WHERE
			p.availability = true
		;
	`
	query, err := r.connector.Exec(sqlStatement, "pendingUserB", userID, targetUserID)
	if err != nil {
		panic(err)
	}

	defer query.Close()

	for query.Next() {
		u := &user.User{}
		if err := rows.Scan(&u.UserID, &u.Name, &u.Password, &u.Email); err != nil {
			panic(err)
		}
		userList := append(userList, u)
	}

	return userList, nil
}

func (r *recommendationRepo) monoMatch(userID, targetUserID int) error {
	// TODO, need to refactor what this does

	// sqlStatement := `
	// 	INSERT INTO
	// 		matchrequest( status, usera, userb )
	// 	VALUES
	// 		( $1, $2, $3 )
	// 	;
	// `
	// _, err := db.connector.Exec(sqlStatement, "pendingUserB", userId, targetUserID)

	// if err != nil {panic(err)}

	// return err
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

func (r *recommendationRepo) saveMatch(userID, targetUserID int) (RecomendationResponse, error) {
	// TODO, may not even need it

	return 0, nil
}
