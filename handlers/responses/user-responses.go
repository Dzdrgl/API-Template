package responses

import (
	"github.com/google/uuid"
)

type RegisterUserResponse struct {
	Message string `json:"message"`
}

type LoginUserResponse struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	ProfileURL string    `json:"profile_url"`
}
