package entity

import "time"

type OtpCodeEntity struct {
	Id        int64
	Code      string
	ExpiredAt time.Time
	UserId    int64
	User      *UserEntity
}
