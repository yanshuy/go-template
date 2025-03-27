-- name: CreateUser :one
INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = $1;