package user

import (
	"time"
)

// Menggabungkan simple registrasi dengan detail informasi,
// dengan model satu informasi alamat user.
type InsertUser struct {
	Email           string `json:"email" binding:"email,max=100"`
	FirstName       string `json:"first_name" binding:"required,max=100"`
	LastName        string `json:"last_name" binding:"required,max=100"`
	Password        string `json:"password" binding:"required,max=128"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=Password"`
	Address         string `json:"address" binding:"required,max=200"`
	ShipperAreaID   uint64 `json:"shipper_area_id" binding:"required"`
	ShipperLat      string `json:"shipper_lat" binding:"required,max=50"`
	ShipperLng      string `json:"shipper_lng" binding:"required,max=50"`
}

func (i *InsertUser) toUser() User {
	return User{
		Email:         i.Email,
		FirstName:     i.FirstName,
		LastName:      i.LastName,
		Address:       i.Address,
		ShipperAreaID: i.ShipperAreaID,
		ShipperLat:    i.ShipperLat,
		ShipperLng:    i.ShipperLng,
	}
}

type User struct {
	ID            uint64
	Email         string
	FirstName     string
	LastName      string
	Password      string
	Address       string
	ShipperAreaID uint64
	ShipperLat    string
	ShipperLng    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type Token struct {
	Token string `json:"token"`
}
