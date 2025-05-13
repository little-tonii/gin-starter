package response

type RegisterUserResponse struct {
	Message string `json:"message"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
}
