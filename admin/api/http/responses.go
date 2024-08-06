package http

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}


type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code"`
}
