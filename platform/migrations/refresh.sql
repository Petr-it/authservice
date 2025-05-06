CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    token_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);