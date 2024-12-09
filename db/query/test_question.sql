-- name: CreateTestQuestion :one
INSERT INTO m_test_questions ( id, job_id, question, choice_a, choice_b, choice_c, choice_d, answer)
VALUES ( $1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetQuestionById :one
SELECT * FROM m_test_questions
WHERE id = $1;

-- name: GetTestQuestions :many
SELECT * FROM m_test_questions;

-- name: GetQuestionsByJobId :many
SELECT * FROM m_test_questions
WHERE job_id = $1;

-- name: UpdateTestQuestion :one
UPDATE m_test_questions
SET
    job_id = $1, question = $2, choice_a = $3,
    choice_b = $4, choice_c = $5, choice_d = $6,
    answer = $7
WHERE id = $8
RETURNING *;

-- name: DeleteTestQuestionById :one
DELETE FROM m_test_questions
WHERE id = $1
RETURNING *;