// internal/restaurant/repository/postgresql.go
package repository

import (
	"database/sql"
	"errors"
	"os"
	"restaurant/internal/restaurant/domain"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var ErrNotFound = errors.New("restaurant not found")

type postgresRepository struct {
	db *sql.DB
}

// TODO: Use sqlc to generate the queries
func NewPostgresRepository() (domain.Repository, error) {
	connStr := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	// Ensure the database is reachable
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db: db}, nil
}

func (r *postgresRepository) Create(restaurant domain.Restaurant) error {
	query := `INSERT INTO restaurants (id, name, address) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, restaurant.ID, restaurant.Name, restaurant.Address)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresRepository) Get(id string) (domain.Restaurant, error) {
	var restaurant domain.Restaurant
	query := `SELECT id, name, address FROM restaurants WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&restaurant.ID, &restaurant.Name, &restaurant.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return restaurant, ErrNotFound
		}
		return restaurant, err
	}

	return restaurant, nil
}

func (r *postgresRepository) Update(id string, restaurant domain.Restaurant) error {
	query := `UPDATE restaurants SET name = $1, address = $2 WHERE id = $3`
	_, err := r.db.Exec(query, restaurant.Name, restaurant.Address, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresRepository) Delete(id string) error {
	query := `DELETE FROM restaurants WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}
	return nil
}
