package models

import (
	"time"
)

type RefreshToken struct {
	ID        int       `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"user_id" validate:"required,uuid"`
	TokenHash string    `db:"token_hash" json:"token_hash" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
