-- name: CreateTestQuestion :one
INSERT INTO m_test_questions ( id, job_id, question, answer)
VALUES ( $1, $2, $3, $4)
RETURNING *;

-- name: GetTestQuestionsByJobId :many
SELECT id, job_id, question FROM m_test_questions
WHERE job_id = $1;

-- name: GetTestAnswerByJobId :many
SELECT id, job_id, answer FROM m_test_questions
WHERE job_id = $1;

-- name: UpdateTestQuestion :one
UPDATE m_test_questions
SET
    job_id = $1,
    question = $2,
    answer = $3
WHERE id = $4
RETURNING *;

-- name: DeleteTestQuestionById :one
DELETE FROM m_test_questions
WHERE id = $1
RETURNING *;