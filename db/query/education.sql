-- name: CreateEducation :one
INSERT INTO m_educations (
    id,
    applicant_id,
    major,
    grade,
    instance,
    description,
    location,
    "from",
    "to"
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
) returning *;

-- name: GetEducationByApplicantID :many
SELECT * FROM m_educations WHERE applicant_id = $1;

-- name: UpdateEducationById :one
UPDATE m_educations
SET
    applicant_id = $1,
    major = $2,
    grade = $3,
    instance = $4,
    description = $5,
    location = $6,
    "from" = $7,
    "to" = $8
WHERE id = $9
RETURNING *;

-- name: DeleteEducationById :one
DELETE FROM m_educations
WHERE id = $1
RETURNING *;