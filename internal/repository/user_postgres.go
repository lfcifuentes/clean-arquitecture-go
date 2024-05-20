package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lfcifuentes/clean-arquitecture/internal/domains"
)

type UserPostgresRepository struct {
	db *pgxpool.Pool
}

func NewUserPostgresRepository(db *pgxpool.Pool) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) GetAll() ([]domains.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, name, email, username, created_at FROM users ORDER BY id DESC")
	if err != nil {
		fmt.Println("Error getting users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []domains.User
	for rows.Next() {
		var user domains.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Username,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserPostgresRepository) Save(user domains.User) error {
	_, err := r.db.Exec(
		context.Background(),
		"INSERT INTO users (username, email, name, created_at) VALUES ($1, $2, $3, $4)",
		user.Username,
		user.Email,
		user.Name,
		time.Now(),
	)
	return err
}
