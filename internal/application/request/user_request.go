package request

type RegisterUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required,min=6"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,regexp=^(0|\\+84)[3|5|7|8|9][0-9]{8}$"`
}
