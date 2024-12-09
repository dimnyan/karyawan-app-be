CREATE TYPE "roles" AS ENUM (
    'superadmin',
    'president',
    'human_resource',
    'applicant'
    );

CREATE TYPE "religion" AS ENUM (
    'islam',
    'kristen',
    'katolik',
    'hindu',
    'konghuchu'
    );

CREATE TYPE "sex" AS ENUM (
    'male',
    'female'
    );

CREATE TYPE "criteria" AS ENUM (
    'experience',
    'education',
    'test',
    'health',
    'age',
    'address'
    );

CREATE TYPE "applied_status" AS ENUM (
    'applied',
    'approved_by_hrd',
    'pass',
    'fail'
    );

CREATE TABLE "m_users" (
                           "id" uuid PRIMARY KEY,
                           "applicant_id" uuid,
                           "username" varchar UNIQUE,
                           "password" varchar,
                           "roles_id" bigint
);

CREATE TABLE "m_roles" (
                           "id" serial PRIMARY KEY,
                           "roles" roles
);

CREATE TABLE "m_applicant_datas" (
                                     "id" uuid PRIMARY KEY,
                                     "name" varchar,
                                     "email" varchar NOT NULL,
                                     "phone_number" varchar,
                                     "photo" varchar,
                                     "pob" varchar,
                                     "dob" date,
                                     "sex_id" bigint,
                                     "city" varchar,
                                     "address" varchar,
                                     "religion_id" bigint,
                                     "application_letter" varchar,
                                     "cv" varchar,
                                     "education_certificate" varchar,
                                     "ktp_photo" varchar,
                                     "health_document" varchar,
                                     "updated_at" timestamptz,
                                     "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "m_sex" (
                         "id" serial PRIMARY KEY,
                         "sex" sex
);

CREATE TABLE "m_religion" (
                              "id" serial PRIMARY KEY,
                              "religion" religion
);

CREATE TABLE "m_experiences" (
                                 "id" uuid PRIMARY KEY,
                                 "applicant_id" uuid NOT NULL,
                                 "job_title" varchar NOT NULL,
                                 "description" varchar,
                                 "location" varchar,
                                 "from" date,
                                 "to" date
);

CREATE TABLE "m_educations" (
                                "id" uuid PRIMARY KEY,
                                "applicant_id" uuid NOT NULL,
                                "major" varchar NOT NULL,
                                "grade" float NOT NULL,
                                "instance" varchar NOT NULL,
                                "description" varchar,
                                "location" varchar,
                                "from" date,
                                "to" date
);

CREATE TABLE "t_applicant_scores" (
                                      "id" uuid UNIQUE PRIMARY KEY,
                                      "job_id" uuid,
                                      "applicant_id" uuid NOT NULL,
                                      "experience" integer,
                                      "education" integer,
                                      "test" integer,
                                      "health" integer,
                                      "age" integer,
                                      "address" integer,
                                      "final_score" float,
                                      "status_id" bigint,
                                      "applied_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "m_applied_status" (
                                    "id" serial PRIMARY KEY,
                                    "applied_status" applied_status UNIQUE
);

CREATE TABLE "m_jobs" (
                          "id" uuid PRIMARY KEY,
                          "title" varchar NOT NULL,
                          "description" varchar,
                          "closed_at" integer NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "m_job_criterias" (
                                   "id" uuid PRIMARY KEY,
                                   "job_id" uuid NOT NULL,
                                   "criteria_id" bigint NOT NULL
);

CREATE TABLE "m_criteria" (
                              "id" serial PRIMARY KEY,
                              "criteria" criteria
);

CREATE TABLE "m_test_questions" (
                                    "id" uuid PRIMARY KEY,
                                    "job_id" uuid,
                                    "question" varchar NOT NULL,
                                    "answer" varchar NOT NULL
);

CREATE TABLE "t_test_results" (
                                  "id" uuid PRIMARY KEY,
                                  "applicant_id" uuid NOT NULL,
                                  "question_id" uuid NOT NULL,
                                  "applicant_answer" varchar NOT NULL
);

ALTER TABLE "m_users" ADD FOREIGN KEY ("applicant_id") REFERENCES "m_applicant_datas" ("id");

ALTER TABLE "m_users" ADD FOREIGN KEY ("roles_id") REFERENCES "m_roles" ("id");

ALTER TABLE "m_applicant_datas" ADD FOREIGN KEY ("sex_id") REFERENCES "m_sex" ("id");

ALTER TABLE "m_applicant_datas" ADD FOREIGN KEY ("religion_id") REFERENCES "m_religion" ("id");

ALTER TABLE "m_experiences" ADD FOREIGN KEY ("applicant_id") REFERENCES "m_applicant_datas" ("id");

ALTER TABLE "m_educations" ADD FOREIGN KEY ("applicant_id") REFERENCES "m_applicant_datas" ("id");

ALTER TABLE "t_applicant_scores" ADD FOREIGN KEY ("job_id") REFERENCES "m_jobs" ("id");

ALTER TABLE "t_applicant_scores" ADD FOREIGN KEY ("applicant_id") REFERENCES "m_applicant_datas" ("id");

ALTER TABLE "t_applicant_scores" ADD FOREIGN KEY ("status_id") REFERENCES "m_applied_status" ("id");

ALTER TABLE "m_job_criterias" ADD FOREIGN KEY ("criteria_id") REFERENCES "m_criteria" ("id");

ALTER TABLE "m_test_questions" ADD FOREIGN KEY ("job_id") REFERENCES "m_jobs" ("id");

ALTER TABLE "t_test_results" ADD FOREIGN KEY ("applicant_id") REFERENCES "m_applicant_datas" ("id");

ALTER TABLE "t_test_results" ADD FOREIGN KEY ("question_id") REFERENCES "m_test_questions" ("id");

ALTER TABLE "m_job_criterias" ADD FOREIGN KEY ("job_id") REFERENCES "m_jobs" ("id") ON DELETE CASCADE;

-- default values
INSERT INTO "m_roles" (roles) VALUES ('superadmin'),('president'),('human_resource'),('applicant');
INSERT INTO "m_religion" (religion) VALUES ('islam'), ('kristen'), ('katolik'), ('hindu'), ('konghuchu');
INSERT INTO "m_sex" (sex) VALUES ('male'), ('female');
INSERT INTO "m_criteria" (criteria) VALUES ('experience'), ('education'), ('test'), ('health'), ('age'), ('address');
INSERT INTO "m_applied_status" (applied_status) VALUES ('applied'), ('approved_by_hrd'), ('pass'), ('fail');