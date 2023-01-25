package middlewares

type ErrResponse struct {
	Message     string `json:"message"`
	Status      uint16 `json:"status"`
	Description string `json:"description"`
}
