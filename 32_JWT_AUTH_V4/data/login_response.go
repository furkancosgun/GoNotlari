package data

import "time"

type LoginResponse struct {
	Username  string
	Token     string
	ExpiresAt time.Time
}
