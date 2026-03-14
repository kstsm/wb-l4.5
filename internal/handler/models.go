package handler

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Result any `json:"result"`
}
