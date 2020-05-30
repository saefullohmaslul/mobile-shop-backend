package response

// Response is template to client response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []Error     `json:"errors"`
}

// Error is template to error response
type Error struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}
