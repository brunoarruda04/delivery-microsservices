-- name: CreateUser :one
INSERT INTO users (id, username, password, role)
VALUES ($1, $2, $3, $4::user_role)
RETURNING id;

-- name: GetUserByUsername :one
SELECT id, username, password, role
FROM users
WHERE username = $1;
