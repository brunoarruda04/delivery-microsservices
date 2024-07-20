package service

import (
	"errors"
	"restaurant/core/restaurant/domain"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("restaurant not found")

type Service interface {
	Create(domain.Restaurant) (string, error)
	Get(string) (domain.Restaurant, error)
	Update(string, domain.Restaurant) error
	Delete(string) error
}

type service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(restaurant domain.Restaurant) (string, error) {
	restaurant.ID = uuid.New().String()
	return restaurant.ID, s.repo.Create(restaurant)
}

func (s *service) Get(id string) (domain.Restaurant, error) {
	return s.repo.Get(id)
}

func (s *service) Update(id string, restaurant domain.Restaurant) error {
	return s.repo.Update(id, restaurant)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
