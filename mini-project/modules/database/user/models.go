package user

import (
	"database/sql"
	"time"

	"github.com/dewidyabagus/go-hexagonal/business/user"
)

type User struct {
	ID            uint64       `gorm:"type:int unsigned;primaryKey;autoIncrement;not null"`
	Email         string       `gorm:"type:varchar(100);uniqueIndex:users_unique_index,priority:1;not null"`
	FirstName     string       `gorm:"type:varchar(100);not null"`
	LastName      string       `gorm:"type:varchar(100);not null"`
	Password      string       `gorm:"type:varchar(128);not null"`
	Address       string       `gorm:"type:varchar(200);not null"`
	ShipperAreaID uint64       `gorm:"column:shipper_area_id;type:int unsigned;not null"`
	ShipperLat    string       `gorm:"column:shipper_lat;type:varchar(50);not null"`
	ShipperLng    string       `gorm:"column:shipper_lng;type:varchar(50);not null"`
	ActiveStatus  uint8        `gorm:"column:active_status;type:tinyint unsigned;not null;default:1;uniqueIndex:users_unique_index,priority:2"`
	CreatedAt     time.Time    `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time    `gorm:"type:datetime;not null"`
	DeletedAt     sql.NullTime `gorm:"type:datetime"`
}

func (u *User) toBusinessUser() *user.User {
	return &user.User{
		ID:            u.ID,
		Email:         u.Email,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Password:      u.Password,
		Address:       u.Address,
		ShipperAreaID: u.ShipperAreaID,
		ShipperLat:    u.ShipperLat,
		ShipperLng:    u.ShipperLng,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
		DeletedAt:     u.DeletedAt.Time,
	}
}
