package repository

import (
	"context"
	"database/sql"

	"authentication/internal/domain"
	db "authentication/sqlc/models"
)

type Repository interface {
	Create(user *domain.User) error
	GetByUsername(username string) (*domain.User, error)
}

type sqlcRepository struct {
	queries *db.Queries
}

func NewSQLCRepository(sql *sql.DB) Repository {
	return &sqlcRepository{
		queries: db.New(sql),
	}
}

func (r *sqlcRepository) Create(user *domain.User) error {
	_, err := r.queries.CreateUser(context.TODO(), db.CreateUserParams{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Column4:  db.UserRole(user.Role),
	})
	return err
}

func (r *sqlcRepository) GetByUsername(username string) (*domain.User, error) {
	user, err := r.queries.GetUserByUsername(context.TODO(), username)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Role:     domain.Role(user.Role),
	}, nil
}
