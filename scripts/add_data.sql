INSERT INTO users (id, name, nickname, email, password_hash, created_at)
VALUES
    ('00000000-0000-0000-0000-000000000001', 'John Doe', 'johndoe', 'johndoe@example.com', 'hashedpassword1', NOW()),
    ('00000000-0000-0000-0000-000000000002', 'Jane Smith', 'janesmith', 'janesmith@example.com', 'hashedpassword2', NOW());


INSERT INTO email_verification_codes (user_id, email, code, created_at, used, expires_at)
VALUES
    ('00000000-0000-0000-0000-000000000001', 'johndoe@example.com', '123456', NOW(), FALSE, NOW() + INTERVAL '1 day'),
    ('00000000-0000-0000-0000-000000000002', 'janesmith@example.com', '654321', NOW(), FALSE, NOW() + INTERVAL '2 days');