package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client = resty.Client

type aggregator struct {
	*resty.Client
}

type Servicer interface {
	GetCountries(ctx context.Context, id, limit, page int) (map[string]interface{}, error)
	GetProvices(ctx context.Context, countryID, id, limit, page int) (*GetProvinces, error)
}

func NewShipperAggrClient(baseUrl, token string) Servicer {
	c := resty.New()
	c.SetBaseURL(baseUrl)
	c.SetHeader("X-API-Key", token)
	c.SetHeader("Content-Type", "application/json")
	c.SetTimeout(180 * time.Second)

	return &aggregator{c}
}

func (c *aggregator) GetCountries(ctx context.Context, id, limit, page int) (map[string]interface{}, error) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	params := map[string]string{}
	if id != 0 {
		params["country_id"] = strconv.Itoa(id)
	}
	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if page != 0 {
		params["page"] = strconv.Itoa(page)
	}

	r := c.NewRequest() // c.R()
	r.SetContext(ctxWT)
	r.SetQueryParams(params)

	// time.Sleep(time.Second)

	res, err := r.Get(c.BaseURL + "/v3/location/countries")
	if err != nil {
		return nil, fmt.Errorf("getting failed: %w", err)

	} else if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	result := map[string]interface{}{}
	if err := json.Unmarshal(res.Body(), &result); err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %w", err)
	}

	return result, nil
}

func (c *aggregator) GetProvices(ctx context.Context, countryID, id, limit, page int) (*GetProvinces, error) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	params := map[string]string{}
	if id != 0 {
		params["province_id"] = strconv.Itoa(id)
	}
	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if page != 0 {
		params["page"] = strconv.Itoa(page)
	}

	r := c.NewRequest() // c.R()
	r.SetContext(ctxWT)
	r.SetQueryParams(params)
	r.SetPathParam("country_id", strconv.Itoa(countryID))

	res, err := r.Get(c.BaseURL + "/v3/location/country/{country_id}/provinces")
	if err != nil {
		return nil, fmt.Errorf("getting failed: %w", err)

	} else if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	result := new(GetProvinces)
	if err := json.Unmarshal(res.Body(), result); err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %w", err)
	}

	return result, nil
}
