package topic

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExists = errors.New("topic already exists")
	ErrTopicNotFound = errors.New("topic not found")
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name string) (*Topic, error) {
	_, err := s.repo.GetByName(name)
	if err == nil {
		return nil, ErrAlreadyExists
	}
	topic := &Topic{Name: name}
	if err := s.repo.Create(topic); err != nil {
		return nil, fmt.Errorf("failed to create topic: %w", err)
	}

	return topic, nil
}

func (s *Service) GetByID(id uint) (*Topic, error) {
	topic, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrTopicNotFound
	}
	return topic, err
}

func (s *Service) GetAll() ([]Topic, error) {
	return s.repo.GetAll()
}

func (s *Service) Delete(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return ErrTopicNotFound
	}

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete topic: %w", err)
	}
	return nil
}
