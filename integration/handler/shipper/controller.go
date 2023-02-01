package shipper

import (
	"net/http"

	"mini-project/integration/httpclient"

	"github.com/gin-gonic/gin"
)

type controller struct {
	client httpclient.Servicer
}

func NewController(c httpclient.Servicer) *controller {
	return &controller{client: c}
}

func (c *controller) GetCountries(ctx *gin.Context) {
	params := new(FormCountry)
	if err := ctx.BindQuery(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := c.client.GetCountries(ctx, params.ID, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *controller) GetProvinces(ctx *gin.Context) {
	params := new(FormProvince)
	if err := ctx.ShouldBindQuery(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := c.client.GetProvices(ctx, params.CountryID, params.ID, params.Limit, params.Page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
