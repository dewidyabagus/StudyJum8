package user

import (
	"github.com/dewidyabagus/go-hexagonal/business/user"
	"github.com/dewidyabagus/go-hexagonal/constant/format"
)

// Tag validation: https://pkg.go.dev/github.com/go-playground/validator/v10

type UserLogin struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type Profile struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toResponseProfile(u *user.User) Profile {
	return Profile{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Address:   u.Address,
		CreatedAt: u.CreatedAt.Format(format.DateTimeLayout),
		UpdatedAt: u.UpdatedAt.Format(format.DateTimeLayout),
	}
}
