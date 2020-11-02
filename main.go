package main

import (
	auth "github.com/eecs4314prismbreak/WheyPal/auth"
	user "github.com/eecs4314prismbreak/WheyPal/user"

	"github.com/gin-gonic/gin"
)

var (
	userSrv *user.UserService
	authSrv *auth.AuthService
)

func main() {
	router := gin.Default()

	userSrv = user.NewService()
	authSrv = auth.NewService()

	router.GET("/", homeHandler)
	router.GET("/user", auth.CheckJWT(), getAllUsers)
	router.GET("/user/:id", auth.CheckJWT(), getUser)
	router.PUT("/user", auth.CheckJWT(), updateUser)
	router.POST("/user", createUser)
	router.POST("/login", login)
	router.PUT("/login", auth.CheckJWT(), updateLogin)
	router.POST("/auth", auth.CheckJWT(), validate)

	router.Run(":8080")
}
