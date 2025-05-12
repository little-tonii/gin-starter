package entity

type UserEntity struct {
	Id          int64
	Username    string
	Password    string
	Email       string
	PhoneNumber string
	Role        *RoleEntity
}
