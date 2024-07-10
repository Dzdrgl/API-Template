package config

import (
	"isteportal-api/handlers"
)
import "github.com/gorilla/mux"

type APIHandlers struct {
	UserHandler *handlers.UserHandler
}

// RegisterRoutes ana router'ı yapılandırır ve tüm alt route'ları kaydeder
func RegisterRoutes(r *mux.Router, handlers *APIHandlers) {
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	RegisterUserRoutes(userRouter, handlers.UserHandler)
}

// User Routes
func RegisterUserRoutes(r *mux.Router, userHandler *handlers.UserHandler) {
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
}
