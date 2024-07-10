package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"isteportal-api/models"
)

type UserRepository interface {
	FetchUserIdByUsername(username string) (uuid.UUID, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetHashedPasswordByUsername(ctx context.Context, userID uuid.UUID) (string, error)
	IsUsernameUnique(ctx context.Context, username string) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GenerateToken(userID string) (string, error)
}

type userRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewUserRepository(db *sql.DB, redis *redis.Client) UserRepository {
	return &userRepository{db, redis}
}

func (r *userRepository) GetHashedPasswordByUsername(ctx context.Context, userID uuid.UUID) (string, error) {
	var hashedPass string
	query := `SELECT password FROM users WHERE id=$1`
	if err := r.db.QueryRowContext(ctx, query, userID).Scan(&hashedPass); err != nil {
		return "", err
	}

	return hashedPass, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, profile_url FROM users WHERE id=$1`
	if err := r.db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Username, &user.ProfileURL); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, username, password, profile_url) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Username, user.Password, user.ProfileURL)
	if err != nil {
		return err
	}

	return r.mapUserRedis(user.ID.String(), user.Username)
}

func (r *userRepository) IsUsernameUnique(ctx context.Context, username string) error {
	key := fmt.Sprintf("username:%s", username)
	exists, err := r.redis.Exists(key).Result()
	if err != nil {
		return err
	}
	if exists == 1 {
		return fmt.Errorf("user %s already exists", username)
	}

	return nil
}

func (r *userRepository) FetchUserIdByUsername(username string) (uuid.UUID, error) {
	usernameKey := fmt.Sprintf("username:%s", username)
	value, err := r.redis.Get(usernameKey).Result()
	if err != nil {
		if err == redis.Nil {
			return uuid.Nil, fmt.Errorf("username does not exist")
		}
		return uuid.Nil, err
	}
	return uuid.Parse(value)
}

func (r *userRepository) GenerateToken(userID string) (string, error) {
	token := uuid.New().String()
	tokenKey := fmt.Sprintf("token:%s", token)
	err := r.redis.Set(tokenKey, userID, 0).Err()
	return token, err
}

func (r *userRepository) mapUserRedis(id, username string) error {
	usernameKey := fmt.Sprintf("username:%s", username)
	return r.redis.Set(usernameKey, id, 0).Err()
}
