package user

import (
	"context"
	"time"

	svcUser "github.com/dewidyabagus/go-hexagonal/business/user"
	"github.com/dewidyabagus/go-hexagonal/modules/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) PostUser(ctx context.Context, user svcUser.User) error {
	return database.WrapError(
		r.db.Create(r.convertModelUser(user)).Error,
	)
}

func (r *repository) GetUserWithEmail(ctx context.Context, email string) (*svcUser.User, error) {
	var user = &User{}

	err := r.db.Select("id, password").
		First(user, "email = ? AND active_status = 1", email).Error
	if err != nil {
		return nil, database.WrapError(err)
	}

	return user.toBusinessUser(), nil
}

func (r *repository) GetUserWithID(ctx context.Context, id uint64) (*svcUser.User, error) {
	var user = &User{}

	// SELECT * FROM users WHERE id = ? AND active_status = 1 ORDER BY id LIMIT 1
	if err := r.db.First(user, "id = ? AND active_status = 1", id).Error; err != nil {
		return nil, err
	}

	return user.toBusinessUser(), nil
}

func (r *repository) PutUserWithID(ctx context.Context, id uint64, user svcUser.User) error {
	userUpdate := &User{
		Email:     user.Email, // Nilai field empty / zero value data type akan diabaikan
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address:   user.Address,
		UpdatedAt: user.UpdatedAt,
	}

	// UPDATE users SET email = ?, first_name = ?, last_name = ? .... WHERE id = ? AND active_status = 1;
	return r.db.Where("id = ? AND active_status = 1", id).Updates(userUpdate).Error
}

// func (r *repository) UserLogin(ctx context.Context, email, password string) (user svcUser.User, err error)
// func (r *repository) ChangePassword(ctx context.Context, old, new string) error

// Method untuk menkonversi struct
func (r *repository) convertModelUser(u svcUser.User) *User {
	return &User{
		Email:         u.Email,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Password:      u.Password,
		Address:       u.Address,
		ShipperAreaID: u.ShipperAreaID,
		ShipperLat:    u.ShipperLat,
		ShipperLng:    u.ShipperLng,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
