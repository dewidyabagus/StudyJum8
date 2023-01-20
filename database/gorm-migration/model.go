package main

import "time"

// Menentukan tipe data field:-> type
// Pemisah antar config menggunakan semicolon ;
// Untuk lebih lengkapnya: https://gorm.io/docs/models.html

type User struct {
	ID        uint64 `gorm:"column:id;type:int unsigned;primaryKey;not null"` // Eksplisit mapping name field column:id
	Email     string `gorm:"type:varchar(100);not null"`
	FirstName string `gorm:"type:varchar(100);not null"` // nama field first_name
	LastName  string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"type:varchar(128);not null"`
	Address   string `gorm:"type:varchar(200);not null"`
	// Handphone  string    `gorm:"type:varchar(10);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"` // default field name: create_at
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
	DeletedAt time.Time `gorm:"type:datetime"`
}
