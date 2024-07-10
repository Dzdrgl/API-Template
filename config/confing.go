package config

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"log"
)

func InitDatabaseConnections() (*sql.DB, *redis.Client) {
	db, err := sql.Open("postgres", "user=example password=example dbname=your_db sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return db, redisClient
}
