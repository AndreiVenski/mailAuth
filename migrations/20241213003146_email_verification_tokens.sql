-- +goose Up
CREATE TABLE email_verification_tokens (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
)

CREATE INDEX idx_verification_token ON email_verification_tokens (token);

-- +goose Down
DROP TABLE IF EXISTS email_verification_tokens;