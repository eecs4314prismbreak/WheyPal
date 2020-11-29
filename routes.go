package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	auth "github.com/eecs4314prismbreak/WheyPal/auth"
	rec "github.com/eecs4314prismbreak/WheyPal/recommendation"
	user "github.com/eecs4314prismbreak/WheyPal/user"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	//fix for specific origin later
	CheckOrigin: func(r *http.Request) bool { return true },
}

// the '/' endpoint
func homeHandler(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "hello"},
	)
}

// GET /users
func getAllUsers(c *gin.Context) {
	resp, err := userSrv.GetAllUsers()

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

// GET /users/:id
func getUser(c *gin.Context) {
	// id, _ := strconv.Atoi(c.Param("id"))

	idFromToken := c.GetInt("userID")

	fmt.Printf("[UserService] [GetProfile] ID: %v ", idFromToken)
	// if idFromToken != id {
	// 	c.JSON(401, fmt.Sprintf("%v", errors.New("UserID does not match claim from token")))
	// 	return
	// }

	resp, err := userSrv.Get(idFromToken)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

//POST /users
func createUser(c *gin.Context) {
	type CreateMessage struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Birthdate string `json:"birthdate"`
		Location  string `json:"location"`
		Interest  string `json:"interest"`
	}

	var message *CreateMessage
	// fmt.Println("message", *message)

	err := c.ShouldBind(&message)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	fmt.Printf("[UserService] [CreateProfile] Name: %s %s ", message.FirstName, message.LastName)

	login := &auth.Login{
		Email:    message.Email,
		Password: message.Password,
	}

	authResponse, err := authSrv.Create(login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	user := &user.User{
		UserID:    authResponse.UserID,
		FirstName: message.FirstName,
		LastName:  message.LastName,
		Birthdate: message.Birthdate,
		Location:  message.Location,
		Interest:  message.Interest,
	}

	_, err = userSrv.Create(user)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &authResponse)
}

//PUT /user
func updateUser(c *gin.Context) {
	var user *user.User

	err := c.ShouldBind(&user)

	idFromToken := c.GetInt("userID")

	fmt.Printf("[UserService] [UpdateProfile] ID: %v ", idFromToken)
	// if idFromToken != user.UserID {
	// 	c.JSON(401, fmt.Sprintf("%v", errors.New("UserID adoes not match claims from token")))
	// 	return
	// }

	user.UserID = idFromToken

	updatedUser, err := userSrv.Update(user)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &updatedUser)
}

//PUT /login
func updateLogin(c *gin.Context) {
	var login *auth.LoginRequest
	err := c.ShouldBind(&login)

	idFromToken := c.GetInt("userID")

	fmt.Printf("[AuthService] [UpdateLogin] ID: %v ", idFromToken)
	// if login.UserID != idFromToken {
	// 	c.AbortWithStatusJSON(401, fmt.Sprintf("UserID from Token does not match UserID in request body"))
	// 	return
	// }

	resp, err := authSrv.Update(idFromToken, login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	//DEPRECIATED: USER PROFILE NO LONGER TRACKS EMAIL
	//update user profile email too
	// if login.Email != "" {
	// 	user := &user.User{
	// 		UserID: idFromToken,
	// 		Email:  login.Email,
	// 	}

	// 	_, err := userSrv.Update(user)

	// 	if err != nil {
	// 		c.JSON(500, fmt.Sprintf("%v", err))
	// 		return
	// 	}
	// }

	c.JSON(200, &resp)
}

func login(c *gin.Context) {
	type LoginResponse struct {
		*user.User
		Email string `json:"email"`
		*auth.StoredToken
	}

	var login *auth.LoginRequest

	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	fmt.Printf("[AuthServie] [Login] EMAIL: %s ", login.Email)

	authResponse, err := authSrv.Login(login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	userResponse, err := userSrv.Get(authResponse.UserID)
	// fmt.Printf("userID", userResponse.UserID)
	resp := &LoginResponse{
		User:        userResponse,
		Email:       authResponse.Email,
		StoredToken: authResponse.StoredToken,
	}
	// fmt.Println("resp", resp)
	c.JSON(200, &resp)
}

func validate(c *gin.Context) {
	var request *auth.AuthRequest
	c.ShouldBind(&request)

	idFromToken := c.GetInt("userID")

	resp, err := authSrv.ValidateToken(idFromToken, request.Token)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

func recommend(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	// var jwt string

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read token:", err)
		return
	}

	// err = json.Unmarshal(message, jwt)
	// if err != nil {
	// 	log.Println("unmarshall jwToken:", err)
	// 	return
	// }

	claims, err := auth.ClaimsFromToken(string(message))
	if err != nil {
		log.Printf("[RecService] [Error] %v", err)
		return
	}
	userID := claims.UserID

	var recommenaditonMessage *rec.RecommendationMessage
	var recommenaditonResponse rec.RecommendationResponse
	var recs []*user.User
	count := 0

	recs, err = recSrv.GetRecommendations(userID)
	// recs, err = recSrv.GetRecommendations(5)
	if err != nil {
		log.Printf("[RecService] [Error] UID %v | Error %v", userID, err)
		return
	}

	// testUser := &user.User{
	// 	UserID: 1,
	// 	Name:   "Allen",
	// }

	// recs = []*user.User{testUser}

	recMsg, _ := json.Marshal(recs)
	// log.Printf("SENDING RECS TO USER | UID %v | RECS %s", userID, recMsg)
	log.Printf("[RecService] [SendRecommendations] UID %v | # of RECS %v", userID, len(recs))

	err = conn.WriteMessage(1, recMsg) //1 = text, 2 = binacy
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		// log.Printf("recv: %s", message)

		// log.Printf("REVECIED REC RESPONSE | RAW %s | MASHALLED RESP %v", message, recommenaditonMessage)
		log.Printf("REVECIED REC RESPONSE | RAW %s", message)

		err = json.Unmarshal(message, &recommenaditonMessage)
		if err != nil {
			log.Printf("Could not unmashall recommendation response | MT %v | RESP: %s\nERR: %v", mt, message, err)
			break
		}
		// log.Printf("recv: %s", message)
		recommenaditonMessage.UserID1 = userID

		recommenaditonResponse, err = recSrv.HandleRecommendationResponse(recommenaditonMessage)
		if err != nil {
			log.Printf("Could not handle rec response | RESP: %v\nERR: %v", recommenaditonResponse, err)
			break
		}

		response, err := json.Marshal(recommenaditonResponse)
		if err != nil {
			log.Printf("Could not mashall outgoing rec response |RESP: %s\nERR: %v", response, err)
			break
		}

		log.Printf("SENDING REC RESPONSE | UID %v | RECS REMAINING %v | MASHALLED RESP %s", userID, count, response)
		err = conn.WriteMessage(mt, response)
		if err != nil {
			log.Println("write:", err)
			break
		}
		count++

		if count == len(recs) {
			count = 0
			recs, err = recSrv.GetRecommendations(userID)
			if err != nil {
				log.Printf("ERROR SENDING REC ON WEBSOCKET | %v", err)
				return
			}
			recMsg, _ := json.Marshal(recs)

			err = conn.WriteMessage(mt, recMsg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func getMatches(c *gin.Context) {
	idFromToken := c.GetInt("userID")

	resp, err := userSrv.GetMatches(idFromToken)
	if err != nil {
		log.Printf("ERROR GETTING MATCHES | %v", err)
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

func deleteMatch(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idFromToken := c.GetInt("userID")

	fmt.Printf("[UserService] [DeleteMatch] UID: %v | TARGET ID: %v ", id, idFromToken)

	resp, err := userSrv.DeleteMatch(idFromToken, id)

	if err != nil {
		log.Printf("ERROR GETTING MATCHES | %v", err)
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

func showLogs(c *gin.Context) {
	var password string
	c.ShouldBind(&password)
	if password != "showmethelogs" {
		c.Abort()
	}

	content, err := ioutil.ReadFile("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	c.String(200, text)
}

var requests = make(map[string][]time.Time)

func ping(c *gin.Context) {
	processPing(c.ClientIP())
	requests[c.ClientIP()] = append(requests[c.ClientIP()], time.Now())
	if len(requests[c.ClientIP()]) > 3 {
		log.Printf("[RateLimitter] [RequestRejected] IP: %s", c.ClientIP())
		c.JSON(502, "you are sending too many requests, please wait 1 minute before sending aother request")
	} else {
		log.Printf("[RateLimitter] [RequestRecieved] IP: %s", c.ClientIP())
		c.JSON(200, gin.H{"message": "pong"})
	}
}

func processPing(clientIP string) {
	i := 0
	for _, ping := range requests[clientIP] {
		tempPing := ping.Add(time.Minute)
		if tempPing.Before(time.Now()) {
			requests[clientIP] = append(requests[clientIP][:i], requests[clientIP][i+1:]...)
		} else {
			i++
		}
	}
}
