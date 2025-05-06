package queries

import (
	"authservice/app/models"

	"github.com/jmoiron/sqlx"
)

type TokenQueries struct {
	*sqlx.DB
}

func (q *TokenQueries) CreateToken(t *models.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (user_id, token_hash, created_at) VALUES ($1, $2, $3)`

	_, err := q.Exec(query, t.UserID, t.TokenHash, t.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
