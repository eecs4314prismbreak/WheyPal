package main

import (
	"os"
	"time"

	auth "github.com/eecs4314prismbreak/WheyPal/auth"
	user "github.com/eecs4314prismbreak/WheyPal/user"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	userSrv user.UserService
	authSrv auth.AuthService
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}

	redisAddr := os.Getenv("REDIS")
	if len(redisAddr) == 0 {
		redisAddr = "localhost:6379"
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://wheypal.com", "http://localhost:8080"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		// AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userSrv = user.NewService()
	authSrv = auth.NewService(redisAddr)

	// router.GET("/", homeHandler)
	// router.GET("/user", auth.CheckJWT(), getAllUsers)
	// router.GET("/user/:id", auth.CheckJWT(), getUser)
	// router.PUT("/user", auth.CheckJWT(), updateUser)
	// router.POST("/user", createUser)
	// router.POST("/login", login)
	// router.PUT("/login", auth.CheckJWT(), updateLogin)
	// router.POST("/auth", auth.CheckJWT(), validate)

	router.GET("/", homeHandler)
	router.GET("/user", getAllUsers)
	router.GET("/user/:id", getUser)
	router.PUT("/user", updateUser)
	router.POST("/user", createUser)
	router.POST("/login", login)
	router.PUT("/login", updateLogin)
	router.POST("/auth", validate)

	router.Run(":" + port)
}
