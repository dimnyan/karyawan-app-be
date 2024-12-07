package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"karyawan-app-be/utils"
	"testing"
	"time"
)

func TestCreateNewUser(t *testing.T) {
	applicantUUID := uuid.New()
	name := utils.RandomString(7)
	dob := time.Now().AddDate(0, 0, -1)
	arg := CreateApplicantParams{
		ID:                   applicantUUID,
		Name:                 sql.NullString{String: name, Valid: true},
		Email:                fmt.Sprintf("%v@gmail.com", name),
		PhoneNumber:          sql.NullString{String: "0811111", Valid: true},
		Photo:                sql.NullString{String: "", Valid: true},
		Pob:                  sql.NullString{String: "Jakarta", Valid: true},
		Dob:                  sql.NullTime{Time: dob, Valid: true},
		SexID:                sql.NullInt64{Int64: 1, Valid: true},
		City:                 sql.NullString{String: "San Jose", Valid: true},
		Address:              sql.NullString{String: "", Valid: true},
		ReligionID:           sql.NullInt64{Int64: 1, Valid: true},
		ApplicationLetter:    sql.NullString{String: "", Valid: true},
		Cv:                   sql.NullString{String: "", Valid: true},
		EducationCertificate: sql.NullString{String: "", Valid: true},
		KtpPhoto:             sql.NullString{String: "", Valid: true},
		HealthDocument:       sql.NullString{String: "", Valid: true},
	}

	applicant, err := testQueries.CreateApplicant(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, applicant)
	require.Equal(t, arg.ID, applicant.ID)
	require.Equal(t, arg.Name, applicant.Name)
	require.Equal(t, arg.Email, applicant.Email)
	require.Equal(t, arg.PhoneNumber, applicant.PhoneNumber)
	require.Equal(t, arg.Photo, applicant.Photo)
	require.Equal(t, arg.Pob, applicant.Pob)

	aY, aM, aD := arg.Dob.Time.Date()
	rY, rM, rD := applicant.Dob.Time.Date()
	require.Equal(t, aY, rY)
	require.Equal(t, aM, rM)
	require.Equal(t, aD, rD)

	require.Equal(t, arg.SexID, applicant.SexID)
	require.Equal(t, arg.City, applicant.City)
	require.Equal(t, arg.Address, applicant.Address)
	require.Equal(t, arg.ReligionID, applicant.ReligionID)
	require.Equal(t, arg.ApplicationLetter, applicant.ApplicationLetter)
	require.Equal(t, arg.Cv, applicant.Cv)
	require.Equal(t, arg.EducationCertificate, applicant.EducationCertificate)
	require.Equal(t, arg.KtpPhoto, applicant.KtpPhoto)
	require.Equal(t, arg.HealthDocument, applicant.HealthDocument)
	require.NotEmpty(t, applicant.CreatedAt)

	password := utils.RandomString(8)

	userArg := CreateUserParams{
		ID:          uuid.New(),
		ApplicantID: uuid.NullUUID{UUID: applicantUUID, Valid: true},
		Username:    sql.NullString{String: utils.RandomString(8), Valid: true},
		Password:    sql.NullString{String: password, Valid: true},
		RolesID:     sql.NullInt64{Int64: 1, Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), userArg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, userArg.ID, user.ID)
	require.Equal(t, userArg.Username, user.Username)
	require.Equal(t, userArg.Password, user.Password)
	require.Equal(t, userArg.RolesID, user.RolesID)
	require.Equal(t, userArg.ApplicantID, user.ApplicantID)

	hashPassword, err := utils.HashPassword(password)
	isCorrect := utils.CheckPasswordHash(password, hashPassword)
	require.NoError(t, err)
	require.True(t, isCorrect)
}
