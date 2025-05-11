package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	postgres *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		postgres: pool,
	}
}

func (authRepository AuthRepository) SignUp(ctx context.Context, username string, displayName string, password string) (id int, err error) {
	const query = "INSERT INTO users (username, display_name, password) VALUES ($1, $2, $3) RETURNING id"

	row := authRepository.postgres.QueryRow(
		ctx,
		query,
		username,
		displayName,
		password,
	)

	if err = row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}
