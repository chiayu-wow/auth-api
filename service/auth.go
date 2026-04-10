package service

import (
	"fmt"

	"github.com/chiayu/auth-api/model"
	"github.com/chiayu/auth-api/repository"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req model.RegisterRequest) (model.AuthResponse, error) {
	if s.repo.ExistsByUsername(req.Username) {
		return model.AuthResponse{}, fmt.Errorf("username already taken: %s", req.Username)
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	saved, err := s.repo.Save(user)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		Token:    "fake-token-" + saved.Username,
		Username: saved.Username,
		Email:    saved.Email,
	}, nil
}

func (s *AuthService) Login(req model.LoginRequest) (model.AuthResponse, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("invalid username or password")
	}

	if user.Password != req.Password {
		return model.AuthResponse{}, fmt.Errorf("invalid username or password")
	}

	return model.AuthResponse{
		Token:    "fake-token-" + user.Username,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
