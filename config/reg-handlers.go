package config

import h "isteportal-api/handlers"

type Handlers struct {
	UserHandler h.UserHandler
}

func RegisterHandlers(svcs *Services) *Handlers {
	return &Handlers{
		UserHandler: *h.NewUserHandler(svcs.UserService),
	}
}
