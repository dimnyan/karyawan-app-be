-- name: CreateExperience :one
INSERT INTO m_experiences
(
 id,
 applicant_id,
 job_title,
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
    $7
) RETURNING *;

-- name: GetExperienceByApplicantID :many
SELECT * FROM m_experiences
         WHERE applicant_id = $1;

-- name: UpdateExperienceByID :one
UPDATE m_experiences
SET
    applicant_id = $1,
    job_title = $2,
    description = $3,
    location = $4,
    "from" = $5,
    "to" = $6
WHERE id = $7
RETURNING *;

-- name: DeleteExperienceById :one
DELETE FROM m_experiences
WHERE id = $1
    RETURNING *;