package token

import "time"

// interface for managing tokens
type Maker interface {
	// method for singing a new token
	CreateToken(username string, duration time.Duration) (string, error)
	// method for verifying token
	VerifyToken(token string) (*Payload, error)
}
