package config

import (
	api "isteportal-api/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(api.ContentType) != api.ApplicationJSON {
			api.JsonResponse(w, http.StatusUnsupportedMediaType, nil, "Content-Type header is required")
			return
		}
		authHeader := r.Header.Get("Apikey")
		if authHeader == "" {
			api.JsonResponse(w, http.StatusUnauthorized, nil, "Authorization header is required")
			return
		}
		apiKeyString := strings.TrimPrefix(authHeader, "Key ")

		if apiKeyString != "MYAPIKEY" {
			api.JsonResponse(w, http.StatusUnauthorized, nil, "Invalid API key")
			return
		}
		next.ServeHTTP(w, r)
	})
}
