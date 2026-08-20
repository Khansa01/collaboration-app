package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user User) (User, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, email
	`

	var result User
	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).
		Scan(&result.ID, &result.Name, &result.Email)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return result, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	query := `
		SELECT id, name, email, password
		FROM users
		WHERE email = $1
	`

	var user User
	err := r.db.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
