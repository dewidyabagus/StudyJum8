package user

import (
	"net/http"

	"github.com/dewidyabagus/go-hexagonal/business/user"
	res "github.com/dewidyabagus/go-hexagonal/cmd/api/common/response"
	"github.com/dewidyabagus/go-hexagonal/cmd/api/middlewares"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service user.Servicer
	mid     middlewares.IAuthenticator
}

func NewController(svc user.Servicer, mid middlewares.IAuthenticator) *Controller {
	return &Controller{
		service: svc,
		mid:     mid,
	}
}

func (c *Controller) PostUser(ctx *gin.Context) {
	var user user.InsertUser
	if err := ctx.ShouldBind(&user); err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	if err := c.service.PostUser(ctx, user); err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	res.JSONBasicResponse(ctx, http.StatusCreated, "successfully created user")
}

func (c *Controller) UserLogin(ctx *gin.Context) {
	var login UserLogin
	if err := ctx.ShouldBind(&login); err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	token, err := c.service.UserLogin(ctx, login.Email, login.Password)
	if err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	res.JSONBasicData(ctx, http.StatusOK, "logged in successfully", token)
}

func (c *Controller) GetUser(ctx *gin.Context) {
	claims, err := c.mid.ExtractJWTUser(ctx)
	if err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK", "data": claims})
}

func (c *Controller) PutUser(ctx *gin.Context)        {}
func (c *Controller) ChangePassword(ctx *gin.Context) {}
