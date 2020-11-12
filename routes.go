package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	auth "github.com/eecs4314prismbreak/WheyPal/auth"
	rec "github.com/eecs4314prismbreak/WheyPal/recommendation"
	user "github.com/eecs4314prismbreak/WheyPal/user"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

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
	id, _ := strconv.Atoi(c.Param("id"))

	idFromToken := c.GetInt("userID")
	if idFromToken != id {
		c.JSON(401, fmt.Sprintf("%v", errors.New("UserID does not match claim from token")))
		return
	}

	resp, err := userSrv.Get(id)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

//POST /users
func createUser(c *gin.Context) {
	type CreateMessage struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Birthday string `json:"birthday"`
		Location string `json:"location"`
		Interest string `json:"interest"`
	}

	var message *CreateMessage
	// fmt.Println("message", *message)

	err := c.ShouldBind(&message)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	user := &user.User{
		Name:     message.Name,
		Birthday: message.Birthday,
		Location: message.Location,
		Interest: message.Interest,
	}

	userCreated, err := userSrv.Create(user)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}
	fmt.Printf("IN THE THING: %v\n", user)

	login := &auth.Login{
		UserID:   userCreated.UserID,
		Email:    message.Email,
		Password: message.Password,
	}

	resp, err := authSrv.Create(login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

//PUT /user
func updateUser(c *gin.Context) {
	var user *user.User

	err := c.ShouldBind(&user)

	idFromToken := c.GetInt("userID")
	if idFromToken != user.UserID {
		c.JSON(401, fmt.Sprintf("%v", errors.New("UserID adoes not match claims from token")))
		return
	}

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
		Name     string `json:"name"`
		Birthday string `json:"birthday"`
		Location string `json:"location"`
		Interest string `json:"interest"`
		*auth.AuthResponse
	}

	var login *auth.LoginRequest

	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	authResponse, err := authSrv.Login(login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	userResponse, err := userSrv.Get(authResponse.ID)

	resp := &LoginResponse{
		Name:         userResponse.Name,
		Birthday:     userResponse.Birthday,
		Location:     userResponse.Location,
		Interest:     userResponse.Interest,
		AuthResponse: authResponse,
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
	var recommenaditonMessage *rec.RecommendationMessage
	var recommenaditonResponse rec.RecommendationResponse
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// log.Printf("recv: %s", message)

		err = json.Unmarshal(message, recommenaditonMessage)
		if err != nil {
			log.Println("Could not unmashall recommendation:", err)
			break
		}
		// log.Printf("recv: %s", message)

		recommenaditonResponse, err = recSrv.HandleRecommendationResponse(recommenaditonMessage)

		response, err := json.Marshal(recommenaditonResponse)

		err = conn.WriteMessage(mt, response)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
