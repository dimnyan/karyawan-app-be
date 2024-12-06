-- name: CreateTestResult :one
INSERT INTO t_test_results
(id, applicant_id, question_id, applicant_answer)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateTestResult :one
UPDATE t_test_results
SET
    id = $1, applicant_id = $2, question_id = $3,
    applicant_answer = $4
RETURNING *;


-- name: CreateTestScore :one
INSERT INTO t_applicant_scores
(id, job_id, applicant_id, status_id)
VALUES
    ($1, $2, $3, "1")
RETURNING *;

-- name: InsertApplicantScore :one
UPDATE t_applicant_scores
SET
    experience= $1, education= $2, test= $3,
    health= $4, age= $5, address= $6
WHERE id = $7
RETURNING *;

-- name: InsertApplicantFinalScore :one
UPDATE t_applicant_scores
SET
    final_score= $1
WHERE id = $2
RETURNING *;

-- name: UpdateApplicantStatus :one
UPDATE t_applicant_scores
SET
    status_id = $1
WHERE id = $2
RETURNING *;