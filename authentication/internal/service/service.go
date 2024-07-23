// internal/auth/service/service.go
package service

import (
	"authentication/internal/domain"
	"authentication/internal/repository"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidRole     = errors.New("invalid role")
)

type AuthService interface {
	Register(username, password string, role domain.Role) (string, error)
	Login(username, password string) (string, error)
}

type authService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(username, password string, role domain.Role) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user, err := domain.New(
		uuid.New(),
		username,
		string(hashedPassword),
		role,
	)

	if err != nil {
		return "", err
	}

	return user.ID.String(), s.repo.Create(user)
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrUserNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidPassword
	}
	return user.ID.String(), nil
}
