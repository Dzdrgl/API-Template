package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	ProfileURL string    `json:"profile_url"`
}
