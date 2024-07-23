package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Role string

const (
	RestaurantAdmin Role = "RESTAURANT_ADMIN"
	RestaurantUser  Role = "RESTAURANT_USER"
	UserRole        Role = "USER"
	Delivery        Role = "DELIVERY"
)

var ErrInvalidRole = errors.New("invalid role")

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Role     Role
}

func New(id uuid.UUID, username, password string, role Role) (*User, error) {
	validRoles := map[Role]struct{}{
		RestaurantAdmin: {},
		RestaurantUser:  {},
		UserRole:        {},
		Delivery:        {},
	}

	if _, ok := validRoles[role]; !ok {
		return nil, ErrInvalidRole
	}

	user := &User{
		ID:       id,
		Username: username,
		Password: password,
		Role:     role,
	}

	return user, nil
}
