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

	user, err := c.service.GetUserWithID(ctx, claims.ID)
	if err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	res.JSONBasicData(ctx, http.StatusOK, "OK", toResponseProfile(user))
}

func (c *Controller) PutUser(ctx *gin.Context) {
	claims, err := c.mid.ExtractJWTUser(ctx)
	if err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	var editProfile user.EditProfile
	if err := ctx.ShouldBind(&editProfile); err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	if err := c.service.PutUserWithID(ctx, claims.ID, editProfile); err != nil {
		res.JSONErrorResponse(ctx, err)
		return
	}

	res.JSONBasicResponse(ctx, http.StatusOK, "sukses memperbarui profile")
}

func (c *Controller) ChangePassword(ctx *gin.Context) {}
