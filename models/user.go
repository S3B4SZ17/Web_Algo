package models

import "time"

type User struct {
	Email        string    `bson:"email,omitempty"`
	Token        string    `bson:"token,omitempty"`
	Expiry       time.Time `bson:"expriy,omitempty"`
	RefreshToken string    `bson:"refreshToken`
}
