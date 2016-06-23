package userAuth

type AuthResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func newAuthResponse(success bool, msg string) *AuthResponse {
	resp := AuthResponse{
		Success: success,
		Msg:     msg}
	return &resp
}
