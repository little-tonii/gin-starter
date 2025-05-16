package request

type RegisterUserRequest struct {
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginUserRequest struct {
	Username string `json:"username" form:"username" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ChangePasswordUserRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
