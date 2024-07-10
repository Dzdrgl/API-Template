package config

import (
	"database/sql"
	"github.com/go-redis/redis"
	r "isteportal-api/repositories"
)

type Repositories struct {
	UserRepository r.UserRepository
}

func RegisterRepositories(db *sql.DB, redisClient *redis.Client) *Repositories {
	return &Repositories{
		UserRepository: r.NewUserRepository(db, redisClient),
	}
}
