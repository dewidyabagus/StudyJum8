package routes

import (
	"github.com/dewidyabagus/go-hexagonal/cmd/api/controller/user"
	"github.com/dewidyabagus/go-hexagonal/cmd/api/middlewares"
	"github.com/dewidyabagus/go-hexagonal/config"

	"github.com/gin-gonic/gin"
)

type HandlerConfig struct {
	Config      *config.Api
	HealthCheck healthCheck
	User        *user.Controller
}

func (h *HandlerConfig) CreateRouting(r *gin.Engine) {
	// Custom middleware untuk optional headers
	r.Use(middlewares.Headers())
	// Menggunakan costom logger
	r.Use(middlewares.CustomLogger())
	// Menggunakan custom recovery handler
	r.Use(gin.CustomRecovery(middlewares.ErrorHandler))

	// Routing
	r.GET("/health/ping", h.HealthCheck.Ping)

	// Group routing /v1
	v1 := r.Group("/v1")
	v1.POST("/users", h.User.PostUser)
	v1.POST("/login", h.User.UserLogin)

	// Group routing /v1 dengan auth JWT
	v1JWTAuth := v1.Use(middlewares.JWTMiddlewareAuth(h.Config.JWTSecretKey))
	v1JWTAuth.GET("/profile", h.User.GetUser)
	v1JWTAuth.PUT("/profile", h.User.PutUser)
}
