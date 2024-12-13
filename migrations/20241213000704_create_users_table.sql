-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    name VARCHAR(255),
    nickname VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL   UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
)

CREATE UNIQUE INDEX idx_users_email ON users (email);
CREATE UNIQUE INDEX idx_users_nickname ON users (nickname);


-- +goose Down
DROP TABLE IF EXISTS users;