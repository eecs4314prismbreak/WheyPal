package main

import (
	"os"
	"time"

	auth "github.com/eecs4314prismbreak/WheyPal/auth"
	rec "github.com/eecs4314prismbreak/WheyPal/recommendation"
	user "github.com/eecs4314prismbreak/WheyPal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	userSrv user.UserService
	authSrv auth.AuthService
	recSrv  rec.RecommendationService
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
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
	authSrv = auth.NewService()
	recSrv = rec.NewService()

	router.GET("/", homeHandler)
	router.GET("/users", auth.CheckJWT(), getAllUsers)
	router.GET("/user", auth.CheckJWT(), getUser)
	router.PUT("/user", auth.CheckJWT(), updateUser)
	router.POST("/user", createUser)
	router.POST("/login", login)
	router.PUT("/login", auth.CheckJWT(), updateLogin)
	router.POST("/auth", auth.CheckJWT(), validate)
	router.GET("/recommend", recommend)
	router.GET("/match", auth.CheckJWT(), getMatches)
	router.DELETE("/match/:id", auth.CheckJWT(), deleteMatch)

	if port == "443" {
		router.RunTLS(":"+port, "./config/private/wheypal.com_ssl_certificate.cer", "./config/private/wheypal.com_private_key.key")
	} else {
		router.Run(":" + port)
	}
}
