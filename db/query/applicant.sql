-- name: CreateApplicant :one
INSERT INTO "m_applicant_datas"
(
 id, name, email, phone_number, photo, pob,
 dob, sex_id, city, address, religion_id,
 application_letter, cv, education_certificate,
 ktp_photo, health_document
) VALUES (
          $1,$2,$3,$4,
          $5,$6,$7,$8,
          $9,$10,$11,
          $12,$13,
          $14,$15,
          $16
         ) returning *;

-- name: GetApplicantById :one
SELECT * FROM m_applicant_datas WHERE id = $1 LIMIT 1;

-- name: GetApplicants :many
SELECT * FROM m_applicant_datas;

-- name: GetApplicantsByJobId :many
SELECT *
FROM m_applicant_datas a
JOIN t_applicant_scores b
    ON a.id = b.id
WHERE b.job_id = $1;

-- name: UpdateApplicantById :one
UPDATE m_applicant_datas
SET
    name = $1,
    email = $2,
    phone_number = $3,
    photo = $4,
    pob = $5,
    dob = $6,
    sex_id = $7,
    address = $8,
    religion_id = $9,
    application_letter =$10,
    cv = $11,
    education_certificate = $12,
    ktp_photo = $13,
    health_document = $14,
    city = $15
WHERE id = $16
RETURNING *;