package main

import (
	"github.com/gorilla/mux"
	conf "isteportal-api/config"
	"log"
	"net/http"
	"time"
)

func main() {
	db, redisClient := conf.InitDatabaseConnections()
	defer db.Close()
	defer redisClient.Close()

	// Initialize Repositories, Services and Handlers
	repository := conf.RegisterRepositories(db, redisClient)
	service := conf.RegisterServices(repository)
	handler := conf.RegisterHandlers(service)

	// Creating Routes and binding controllers.
	r := mux.NewRouter()
	r.Use(conf.AuthMiddleware)
	conf.RegisterRoutes(r, &conf.APIHandlers{
		UserHandler: &handler.UserHandler,
	})

	// Sunucuyu başlatma
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r, // mux router
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
