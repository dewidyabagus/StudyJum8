package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dewidyabagus/go-hexagonal/business/user"
	"github.com/dewidyabagus/go-hexagonal/business/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostUser(t *testing.T) {
	var (
		ctx         = context.Background()
		userRepo    = &mocks.Repositorer{}
		userUseCase = user.NewService(userRepo)
		userInsert  = user.InsertUser{
			Password: "AAAA",
		}
	)

	// Menjalankan langsung tanpa subtest
	t.Log("Menjalankan langsung tanpa subtest")
	userRepo.On(
		"PostUser",                                                                 // Nama fungsi pada repository
		mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("user.User"), // Tipe parameter fungsi
	).Return(
		errors.New("ekspektasi nilai kembali"),
	).Once() // Menyatakan ini hanya digunakan sekali
	err := userUseCase.PostUser(ctx, userInsert)             // Pemanggilan use case
	assert.NotNil(t, err)                                    // Menyatakan nilai yang diharapkan bukan nil
	assert.Equal(t, "ekspektasi nilai kembali", err.Error()) // Menyatakan bahwa nilai harus sesuai dengan ekspektasi

	// Menjalankan dengan model subtest
	t.Run("Gagal post data", func(t *testing.T) {
		userRepo.On("PostUser", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("user.User")).Return(errors.New("gagal insert")).Once()

		err := userUseCase.PostUser(ctx, userInsert)
		assert.NotNil(t, err)
		assert.Equal(t, "gagal insert", err.Error())
	})

	userRepo.On("PostUser", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("user.User")).Return(nil)

	t.Run("Sukses post data", func(t *testing.T) {
		err := userUseCase.PostUser(ctx, userInsert)
		assert.Nil(t, err)
	})
}
