-- name: CreateJob :one
INSERT INTO m_jobs (id, title, description, closed_at)
VALUES ( $1, $2, $3, $4)
RETURNING *;

-- name: GetJobs :many
SELECT * FROM m_jobs;

-- name: GetJob :one
SELECT * FROM m_jobs
WHERE id = $1;

-- name: UpdateJob :one
UPDATE m_jobs
SET id = $1, title = $2, description = $3
WHERE id = $4
    RETURNING *;

-- name: DeleteJob :one
DELETE FROM m_jobs
    WHERE id = $1
RETURNING *;

-- name: CreateJobCriteria :one
INSERT INTO m_job_criterias (id, job_id, criteria_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetJobCriteriasByJobId :many
SELECT * FROM m_job_criterias
WHERE job_id = $1;

-- name: DeleteJobCriteria :one
DELETE FROM m_job_criterias
WHERE job_id = $1
RETURNING *;

-- name: DeleteJobCriteriaByJobId :many
DELETE FROM m_job_criterias
WHERE job_id = $1
RETURNING *;