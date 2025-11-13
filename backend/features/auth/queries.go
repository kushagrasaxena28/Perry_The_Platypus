package auth

import (
	"Perry_the_Platypus/backend/features/auth/models"
	"context"
	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, db *pgx.Conn, username, email, passwordHash string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow(ctx, `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, username, email, password_hash, privacy, verified, follower_count, following_count, created_at, updated_at`,
		username, email, passwordHash,
	).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Privacy,
		&user.Verified, &user.FollowerCount, &user.FollowingCount, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(ctx context.Context, db *pgx.Conn, email string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow(ctx, `SELECT id, username, email, password_hash, privacy, verified, follower_count, following_count, created_at, updated_at FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Privacy,
		&user.Verified, &user.FollowerCount, &user.FollowingCount, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(ctx context.Context, db *pgx.Conn, id int) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow(ctx, `SELECT id, username, email, password_hash, privacy, verified, follower_count, following_count, created_at, updated_at FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Privacy,
		&user.Verified, &user.FollowerCount, &user.FollowingCount, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return user, nil
}
