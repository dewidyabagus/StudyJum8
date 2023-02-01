package httpclient

type Metadata struct {
	Path       string `json:"path"`
	StatusCode uint8  `json:"http_status_code"`
	Status     string `json:"http_status"`
	Timestamp  int64  `json:"timestamp"`
}

type Pagination struct {
	CurrentPage     uint64 `json:"current_page"`
	CurrentElements uint64 `json:"current_elements"`
	TotalPage       uint64 `json:"total_pages"`
	TotalElements   uint64 `json:"total_elements"`
}

type GetProvinces struct {
	Metadata   Metadata   `json:"metadata"`
	Data       []Province `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Province struct {
	ID      uint64  `json:"id"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Country Country `json:"country"`
}

type Country struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
