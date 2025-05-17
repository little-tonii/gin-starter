package request

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserRequest struct {
	Username string `json:"username" form:"username" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ChangePasswordUserRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type ForgotPasswordUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOtpResetPasswordRequest struct {
	Email   string `json:"email" binding:"required,email"`
	OtpCode string `json:"otp_code" binding:"required"`
}
