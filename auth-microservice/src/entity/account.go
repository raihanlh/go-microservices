package entity

import "time"

type Account struct {
	Id                   int64     `json:"id"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Role                 int64     `json:"role"`
	Enable               bool      `json:"enable"`
	VerificationToken    string    `json:"verification_token"`
	VerificationTokenExp time.Time `json:"verification_token_exp"`
	Otp                  string    `json:"otp"`
	OtpExp               time.Time `json:"otp_exp"`
	Locked               bool      `json:"locked"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
}
