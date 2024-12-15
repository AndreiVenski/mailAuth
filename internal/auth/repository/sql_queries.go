package repository

const (
	ifExistsUserQuery = `SELECT EXISTS (SELECT 1 FROM users WHERE nickname=$1 OR email=$2)`
	createUserQuery   = `INSERT INTO users (id, name, nickname, email, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING (id, name, nickname, email)`

	addEmailCode = `INSERT INTO email_verification_codes (user_id, email, code, expires_at) VALUES ($1, $2, $3, $4)`

	getIDAndUpdateUsedofEmailCode = `UPDATE email_verification_codes SET used = TRUE WHERE email = $1 AND code = $2 AND used = FALSE AND expires_at > CURRENT_TIMESTAMP RETURNING user_id`

	createRefreshTokenRecord = `INSERT INTO refresh_tokens (id, user_id, token, client_info, expires_at) VALUES ($1, $2, $3, $4, $5)`
)
