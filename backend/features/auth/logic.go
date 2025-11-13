package auth

import (
	"Perry_the_Platypus/backend/features/auth/models"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(ctx context.Context, db *pgx.Conn, input UserInput) (*models.User, error) {
	if input.Username == "" || input.Email == "" || input.Password == "" {
		return nil, errors.New("username, email, and password are required")
	}

	existingUser, _ := GetUserByEmail(ctx, db, input.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := CreateUser(ctx, db, input.Username, input.Email, string(hashedPassword))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Login(ctx context.Context, db *pgx.Conn, input LoginInput) (*models.User, string, error) {
	user, err := GetUserByEmail(ctx, db, input.Email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func generateJWT(userID int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(secret))
}
