package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"karyawan-app-be/utils"
	"testing"
	"time"
)

func TestQueries_CreateJob(t *testing.T) {
	jobId := uuid.New()
	jobString := utils.RandomString(7)
	argsJob := CreateJobParams{
		ID:          jobId,
		Title:       jobString,
		Description: sql.NullString{String: "Desc", Valid: true},
		ClosedAt:    sql.NullTime{Time: time.Now().Add(5 * time.Hour), Valid: true},
	}

	job, err := testQueries.CreateJob(context.Background(), argsJob)

	require.NoError(t, err)
	require.NotEmpty(t, job)

	require.Equal(t, argsJob.ID, job.ID)
	require.Equal(t, argsJob.Title, job.Title)
	require.Equal(t, argsJob.Description.String, job.Description.String)
	require.NotZero(t, job.ClosedAt.Time)

	argsJobCriteria := CreateJobCriteriaParams{
		ID:         uuid.New(),
		JobID:      uuid.NullUUID{UUID: jobId, Valid: true},
		CriteriaID: sql.NullInt64{Int64: 1, Valid: true},
	}

	jobCriteria, err := testQueries.CreateJobCriteria(context.Background(), argsJobCriteria)
	require.NoError(t, err)
	require.NotEmpty(t, jobCriteria)
	require.Equal(t, argsJob.ID, jobCriteria.JobID.UUID)
	require.Equal(t, job.ID, jobCriteria.JobID.UUID)
	require.NotEmpty(t, jobCriteria.CriteriaID)
}
