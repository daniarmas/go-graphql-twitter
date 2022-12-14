package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniarmas/gographqltwitter"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo gographqltwitter.UserRepo
}

func NewAuthService(ur gographqltwitter.UserRepo) *AuthService {
	return &AuthService{
		UserRepo: ur,
	}
}

func (as *AuthService) Register(ctx context.Context, input gographqltwitter.RegisterInput) (gographqltwitter.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return gographqltwitter.AuthResponse{}, err
	}

	// check if username is already taken.
	if _, err := as.UserRepo.GetByUsername(ctx, input.Username); !errors.Is(err, gographqltwitter.ErrNotFound) {
		return gographqltwitter.AuthResponse{}, gographqltwitter.ErrUsernameTaken
	}

	// check if email is already taken.
	if _, err := as.UserRepo.GetByEmail(ctx, input.Username); !errors.Is(err, gographqltwitter.ErrNotFound) {
		return gographqltwitter.AuthResponse{}, gographqltwitter.ErrEmailTaken
	}

	user := gographqltwitter.User{
		Email:    input.Email,
		Username: input.Username,
	}

	// hash the password.
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return gographqltwitter.AuthResponse{}, fmt.Errorf("error hashing password")
	}

	user.Password = string(hashPassword)

	// create the user.
	user, err = as.UserRepo.Create(ctx, user)
	if err != nil {
		return gographqltwitter.AuthResponse{}, fmt.Errorf("error creating user: %v", err)
	}

	// return the access token and user.
	return gographqltwitter.AuthResponse{
		AccessToken: "a token",
		User:        user,
	}, nil
}

func (as *AuthService) Lgoin(ctx context.Context, input gographqltwitter.LoginInput) (gographqltwitter.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return gographqltwitter.AuthResponse{}, err
	}

	user, err := as.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, gographqltwitter.ErrNotFound):
			return gographqltwitter.AuthResponse{}, gographqltwitter.ErrBadCredentials
		default:
			return gographqltwitter.AuthResponse{}, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return gographqltwitter.AuthResponse{}, gographqltwitter.ErrBadCredentials
	}

	return gographqltwitter.AuthResponse{
		AccessToken: "a token",
		User:        user,
	}, nil
}
