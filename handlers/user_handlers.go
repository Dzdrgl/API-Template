package handlers

import (
	"encoding/json"
	"fmt"
	request "isteportal-api/handlers/requests"
	response "isteportal-api/handlers/responses"
	"isteportal-api/models"
	"isteportal-api/services"
	api "isteportal-api/utils"
	"log"
	"net/http"
	"regexp"
)

const (
	UsernameEmpty   = "The username should not be empty"
	PasswordEmpty   = "The password should not be empty"
	InvalidRequest  = "Invalid request"
	InvalidUsername = "Username must be at least 8 max 20 characters long and contain only ASCII characters"
	InvalidPassword = "Password must be at least 8 max 24 characters long, contain at least one uppercase letter, one number, and one special character (_ - .)"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register User")
	var raw request.RegisterUserReq
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&raw); err != nil {
		log.Printf("RegisterUser - Invalid JSON format: %v", err)
		api.JsonResponse(w, http.StatusBadRequest, nil, InvalidRequest)
		return
	}
	var newUser models.User
	if err := parseAndValidateJSON(raw, &newUser); err != nil {
		api.JsonResponse(w, http.StatusBadRequest, nil, InvalidRequest)
		return
	}

	if err := checkUsernameAndPassword(newUser.Username, newUser.Password); err != nil {
		api.JsonResponse(w, http.StatusNotAcceptable, nil, err.Error())
		return
	}

	if err := h.userService.RegisterUser(r.Context(), &newUser); err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, nil, err.Error())
		return
	}
	data := response.RegisterUserResponse{Message: "User registered successfully"}

	api.JsonResponse(w, http.StatusCreated, data, "")
}

func parseAndValidateJSON(raw request.RegisterUserReq, body *models.User) error {
	if raw.Username == nil || raw.Password == nil || raw.ProfileURL == nil {
		return fmt.Errorf("missing required fields")
	}
	if err := json.Unmarshal(raw.Username, &body.Username); err != nil {
		return err
	}
	if err := json.Unmarshal(raw.Password, &body.Password); err != nil {
		return err
	}
	if err := json.Unmarshal(raw.ProfileURL, &body.ProfileURL); err != nil {
		return err
	}

	return nil
}

// !!!!!!!!!!!!Functions
func checkUsernameAndPassword(username string, password string) error {
	if username == "" {
		return fmt.Errorf(UsernameEmpty)
	} else if len(username) < 8 || !isASCII(username) || len(password) > 36 {
		return fmt.Errorf(InvalidUsername)
	}
	if password == "" {
		return fmt.Errorf(PasswordEmpty)
	} else if len(password) < 8 || !isASCII(password) || len(password) > 36 {
		return fmt.Errorf(InvalidPassword)
	} else if !isValidPassword(password) {
		return fmt.Errorf(InvalidPassword)
	}
	return nil
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

func isValidPassword(password string) bool {
	regex := regexp.MustCompile(`[A-Z]`)
	return regex.MatchString(password)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req request.LoginUserReq

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		log.Printf("RegisterUser - Invalid JSON format: %v", err)
		api.JsonResponse(w, http.StatusBadRequest, nil, InvalidRequest)
		return
	}

	if err := checkUsernameAndPassword(req.Username, req.Password); err != nil {
		api.JsonResponse(w, http.StatusNotAcceptable, nil, err.Error())
		return
	}

	userResponse, err := h.userService.LoginUser(r.Context(), &req)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, false, "Error logging in")
		return
	}

	api.JsonResponse(w, http.StatusOK, userResponse, "")
}
