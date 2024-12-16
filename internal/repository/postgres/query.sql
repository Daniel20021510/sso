-- name: CreateUser :one
INSERT INTO users(email, pass_hash)
VALUES(@email, @pass_hash)
RETURNING id;

-- name: GetUser :one
SELECT id, email, pass_hash
FROM users
WHERE email = @email;

-- name: GetApp :one
SELECT id, name, secret
FROM apps
WHERE id = @id;