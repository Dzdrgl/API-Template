package config

import s "isteportal-api/services"

type Services struct {
	UserService s.UserService
}

func RegisterServices(repos *Repositories) *Services {
	return &Services{
		UserService: s.NewUserService(repos.UserRepository),
	}
}
