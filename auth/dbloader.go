package auth

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	dbCredFile = "config/private/db_credentials"
)

func initConfig() {
	viper.SetConfigName(dbCredFile)
	viper.SetConfigType("ini")

	viper.AddConfigPath(".config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func LoadPGDB() *sql.DB {
	initConfig()

	// Construct Config String
	configStr := "user=" + viper.GetString("PGDB.user") +
		" password=" + viper.GetString("PGDB.password") +
		" host=" + viper.GetString("PGDB.host") +
		" port=" + viper.GetString("PGDB.port") +
		" dbname=" + viper.GetString("PGDB.dbname") +
		" sslmode=" + viper.GetString("PGDB.sslmode")

	// Attempt to open connection
	db, err := sql.Open("postgres", configStr)
	if err != nil {
		panic(err)
	}

	// Check if connection successful
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Printf("Auth Database Connection Successful\n")
	}
	return db
}

func LoadRedis() *redis.Client {
	addr := viper.GetString("Redis.address")
	log.Printf("Auth Redis Connection Successful @ %s\n", addr)
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
