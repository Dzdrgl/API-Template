package requests

import "encoding/json"

type RegisterUserReq struct {
	Username   json.RawMessage `json:"username" `
	Password   json.RawMessage `json:"password"`
	ProfileURL json.RawMessage `json:"profile_url"`
}

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
