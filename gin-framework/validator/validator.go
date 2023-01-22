package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// Const untuk formating tanggal
const (
	dateLayout     = "2006-01-02"
	dateTimeLayout = "2006-01-02 15:04:05"
)

// Membuat custom validator
var todayValFunc = func(fl validator.FieldLevel) bool {
	if val, ok := fl.Field().Interface().(string); ok {
		if date, err := time.Parse(dateTimeLayout, val); err != nil {
			return false

		} else if date.Format(dateLayout) != time.Now().Format(dateLayout) {
			return false
		}
	}

	return true
}

// Membuat proses untuk menerjemahkan error ketika data tidak valid
func translateError(err error) error {
	if err == nil {
		return nil
	}

	// Ketika error bukan dari proses validasi data akan langsung
	// dikembalikan tanpa di translate karena tidak bisa diekstraksi
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err

	} else if len(validationErrors) == 0 {
		return nil
	}

	var (
		english  = en.New()
		uni      = ut.New(english, english)
		trans, _ = uni.GetTranslator("en")
	)
	// Tanpa config dibawah ketika gagal validasi ID informasi yang diterima :
	// Key: 'Transaction.ID' Error:Field validation for 'ID' failed on the 'required' tag
	// Dengan menggunakan config dibawah ini untuk gagal validasi ID informasi yang diterima :
	// ID is a required field
	enTranslations.RegisterDefaultTranslations(validate, trans)

	// Tipe data validator.ValidationErrors adalah slice []FieldError
	// jadi kalau ingin mengumpulkan error dan mengembalikan semua daftar errornya dapat
	// melakukan iterasi untuk variabel dengan tipe data tsb.

	switch e := validationErrors[0]; e.Tag() {
	default:
		return errors.New(e.Translate(trans))

	case "required_if", "boolean":
		return fmt.Errorf("%s is a required field", e.Field())

	case "today":
		return fmt.Errorf("field %s nilai %s tanggal harus hari ini", e.Field(), e.Value())
	}
}
