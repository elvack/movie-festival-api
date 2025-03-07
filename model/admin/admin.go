package admin

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Admin struct {
		CreatedAt         time.Time
		Email             string
		EncryptedPassword string
		Id                uint32
		Token             string
		UpdatedAt         time.Time
	}

	AuthClaims struct {
		AdminId uint32
		*jwt.RegisteredClaims
	}
)
