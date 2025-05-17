package response

type RegisterUserResponse struct {
	Message string `json:"message"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
}

type ProfileUserResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type ChanagePasswordUserResponse struct {
	Message string `json:"message"`
}

type ForgotPasswordUserResponse struct {
	Message string `json:"message"`
}

type VerifyOtpResetPasswordRepsonse struct {
	ResetToken string `json:"reset_token"`
}
