package user

import (
	"context"
	"time"
)

type Servicer interface {
	SetJWTConfig(secret string, expired time.Duration)
	PostUser(ctx context.Context, newUser InsertUser) error
	UserLogin(ctx context.Context, email, password string) (*Token, error)
}

type Repositorer interface {
	PostUser(ctx context.Context, user User) error

	GetUserWithEmail(ctx context.Context, email string) (*User, error)
}
