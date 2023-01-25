package main

import (
	"time"

	userusecase "github.com/dewidyabagus/go-hexagonal/business/user"
	userhandler "github.com/dewidyabagus/go-hexagonal/cmd/api/controller/user"
	"github.com/dewidyabagus/go-hexagonal/cmd/api/middlewares"
	usermodule "github.com/dewidyabagus/go-hexagonal/modules/database/user"

	"github.com/dewidyabagus/go-hexagonal/cmd/api/routes"
)

func prepareService(handler *routes.HandlerConfig, depends *dependencies) {
	// Persiapan proses authenticator
	jwtMid := middlewares.NewAuthenticator(handler.Config.JWTSecretKey)

	// Persiapan repository, business dan handler
	userUseCase := userusecase.NewService(usermodule.NewRepository(depends.db))
	userUseCase.SetJWTConfig(
		handler.Config.JWTSecretKey,
		time.Duration(handler.Config.JWTExpiredTime)*time.Minute,
	)
	handler.User = userhandler.NewController(userUseCase, jwtMid)
}
