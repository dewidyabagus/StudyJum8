package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dewidyabagus/go-hexagonal/business"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultSecretKey  = "default-secret-key"
	defaultJWTExpired = 24 * time.Hour
)

type service struct {
	repository   Repositorer
	jwtSecretKey string
	jwtExpired   time.Duration
}

func NewService(repo Repositorer) Servicer {
	return &service{
		repository:   repo,
		jwtSecretKey: defaultSecretKey,
		jwtExpired:   defaultJWTExpired,
	}
}

func (s *service) SetJWTConfig(secret string, expired time.Duration) {
	s.jwtSecretKey = secret
	s.jwtExpired = expired
}

func (s *service) PostUser(ctx context.Context, usr InsertUser) error {
	user := usr.toUser()

	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("generate from password failed: %w", err)
	}
	user.Password = string(hash)

	return s.repository.PostUser(ctx, user)
}

func (s *service) UserLogin(ctx context.Context, email, password string) (*Token, error) {
	user, err := s.repository.GetUserWithEmail(ctx, email)
	if err != nil {
		if errors.Is(err, &business.ErrNotFound{}) {
			return nil, business.NewErrUnauthorized("Incorrect email entered")
		}
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, business.NewErrUnauthorized("Incorrect password entered")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, fmt.Errorf("generate token failed: %w", err)
	}

	return &Token{Token: token}, nil
}

func (s *service) GetUserWithID(ctx context.Context, id uint64) (*User, error) {
	return s.repository.GetUserWithID(ctx, id)
}

func (s *service) PutUserWithID(ctx context.Context, id uint64, user EditProfile) error {
	return s.repository.PutUserWithID(ctx, id, user.toUser())
}

func (s *service) generateToken(user *User) (token string, err error) {
	eJWT := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  user.ID,
			"exp": time.Now().Add(s.jwtExpired).Unix(),
		},
	)

	return eJWT.SignedString([]byte(s.jwtSecretKey))
}
