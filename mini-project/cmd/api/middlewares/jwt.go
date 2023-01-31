package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dewidyabagus/go-hexagonal/business"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type token struct {
	ID  uint64 `json:"id"`
	Exp int64  `json:"exp"`
}

type IAuthenticator interface {
	ExtractJWTUser(ctx *gin.Context) (*token, error)
}

type authenticator struct {
	secretKey string
}

func NewAuthenticator(secretKey string) IAuthenticator {
	return &authenticator{secretKey}
}

func (s *authenticator) ExtractJWTUser(ctx *gin.Context) (*token, error) {

	user, ok := ctx.Get("user") // Key user tidak langsung ada pada header request
	if !ok {
		return nil, business.NewErrUnauthorized("user token not found")
	}

	if _, ok := user.(*jwt.Token); !ok {
		return nil, business.NewErrUnauthorized("invalid token format")
	}

	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)

	res := new(token)
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(&claims)
	json.NewDecoder(buff).Decode(res)

	return res, nil
}

func JWTMiddlewareAuth(jwtSecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(
			c.GetHeader("Authorization"), "Bearer ", "", 1,
		)

		if token = strings.TrimSpace(token); token == "" {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				ErrResponse{
					Message:     "Bad Request",
					Status:      http.StatusBadRequest,
					Description: "Invalid Token",
				},
			)
			return
		}

		res, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
				return nil, errors.New("signing method invalid")
			}

			return []byte(jwtSecretKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				ErrResponse{
					Message:     "Request Unauthorized",
					Status:      http.StatusUnauthorized,
					Description: err.Error(),
				},
			)
			return
		}

		c.Set("user", res)
		c.Next()
	}
}
