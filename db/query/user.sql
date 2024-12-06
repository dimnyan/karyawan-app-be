-- name: CreateUser :one
INSERT INTO m_users (
    id,
    applicant_id,
    username,
    password,
    roles_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;


-- name: GetUserByUsername :one
SELECT * FROM m_users
WHERE username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM m_users
WHERE id = $1 LIMIT 1;

-- name: UpdateUserById :one
UPDATE m_users
SET
    username = $1,
    password = $2
WHERE id = $3
RETURNING *;
