package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreayExists = errors.New("email already exists")
	ErrWrongCredentials  = errors.New("wrong credentials")
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(email, password string) (*User, error) {
	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return nil, ErrEmailAlreayExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hashed password: %w", err)
	}

	user := &User{
		Email:    email,
		Password: string(hashedPass),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *Service) Login(email, password string) (*User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, ErrWrongCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrWrongCredentials
	}

	return user, nil
}
