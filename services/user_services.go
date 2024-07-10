package services

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	request "isteportal-api/handlers/requests"
	"isteportal-api/handlers/responses"
	"isteportal-api/models"
	"isteportal-api/repositories"
)

type UserService interface {
	RegisterUser(ctx context.Context, newUser *models.User) error
	LoginUser(ctx context.Context, req *request.LoginUserReq) (*responses.LoginUserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) LoginUser(ctx context.Context, req *request.LoginUserReq) (*responses.LoginUserResponse, error) {
	userId, err := s.userRepo.FetchUserIdByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	hashedPass, err := s.userRepo.GetHashedPasswordByUsername(ctx, userId)
	if err != nil {
		return nil, err
	}

	if err := compareHashedPassword(hashedPass, req.Password); err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	response := &responses.LoginUserResponse{
		Id:         user.ID,
		Username:   user.Username,
		ProfileURL: user.ProfileURL,
	}
	return response, nil
}

func (s *userService) RegisterUser(ctx context.Context, newUser *models.User) error {
	fmt.Println("Register User Services")
	if err := s.userRepo.IsUsernameUnique(ctx, newUser.Username); err != nil {
		return err
	}
	hashedPass, err := hashPassword(newUser.Password)
	if err != nil {
		return fmt.Errorf("Hashing password failed: %v", err)
	}
	newUser.Password = hashedPass

	return s.userRepo.CreateUser(ctx, newUser)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func compareHashedPassword(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return fmt.Errorf("Incorrect password")
	}
	return nil
}
