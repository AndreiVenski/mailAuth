-- +goose Up
CREATE TABLE email_verification_codes (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL,
    code VARCHAR(6) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    used BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_verification_code_email ON email_verification_codes (email, code);
-- +goose Down
DROP TABLE IF EXISTS email_verification_tokens;