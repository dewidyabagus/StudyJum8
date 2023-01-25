package validator

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type ValidationErrors = validator.ValidationErrors

var once sync.Once
var mx = new(sync.Mutex)
var validate *validator.Validate

func New() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})

	return validate
}

// Membuat proses untuk menerjemahkan error ketika data tidak valid
func TranslateError(err error) error {
	mx.Lock()
	defer mx.Unlock()

	if err == nil {
		return nil

	} else if validate == nil {
		New()
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

	case "required", "required_if", "boolean":
		return fmt.Errorf("%s is a required field", e.Field())

	case "email":
		return fmt.Errorf("%s field format must be %s", e.Field(), e.Tag())

	case "eqfield":
		return fmt.Errorf("field validation for '%s' must equal (%s) '%s'", e.Field(), e.Tag(), e.Param())

	case "today":
		return fmt.Errorf("field %s nilai %s tanggal harus hari ini", e.Field(), e.Value())
	}
}
