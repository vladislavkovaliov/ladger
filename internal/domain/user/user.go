package user

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type UserResponse struct {
	ID        string
	Email     string
	CreatedAt time.Time
}
