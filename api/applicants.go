package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "karyawan-app-be/db/sqlc"
	"karyawan-app-be/utils"
	"net/http"
)

type RegisterApplicantsRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterApplicantsResponse struct {
	UserId      string `json:"user_id"`
	ApplicantId string `json:"applicant_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}

func (server *Server) RegisterApplicant(ctx *gin.Context) {
	var req RegisterApplicantsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	applicantUUID := uuid.New()
	userUUID := uuid.New()
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	}

	applicantParams := db.CreateApplicantParams{
		ID:    applicantUUID,
		Email: req.Email,
	}
	userParams := db.CreateUserParams{
		ID:          userUUID,
		ApplicantID: uuid.NullUUID{UUID: applicantUUID, Valid: true},
		Username:    sql.NullString{String: req.Username, Valid: true},
		Password:    sql.NullString{String: hashedPassword, Valid: true},
		RolesID:     sql.NullInt64{Int64: 4, Valid: true},
	}

	applicant, err := server.store.CreateApplicant(ctx, applicantParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	}

	user, err := server.store.CreateUser(ctx, userParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	}

	var resp RegisterApplicantsResponse
	resp.UserId = user.ID.String()
	resp.ApplicantId = applicant.ID.String()
	resp.Username = user.Username.String
	resp.Email = applicant.Email

	ctx.JSON(http.StatusCreated, resp)
}

type ApplicantDataUpdateRequest struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	PhoneNumber          string `json:"phone_number"`
	Photo                string `json:"photo"`
	Pob                  string `json:"pob"`
	Dob                  string `json:"dob"`
	SexID                int    `json:"sex_id"`
	City                 string `json:"city"`
	Address              string `json:"address"`
	ReligionID           int    `json:"religion_id"`
	ApplicationLetter    string `json:"application_letter"`
	Cv                   string `json:"cv"`
	EducationCertificate string `json:"education_certificate"`
	KtpPhoto             string `json:"ktp_photo"`
	HealthDocument       string `json:"health_document"`
}

type ApplicantDataUpdateResponse struct {
	ID string `json:"id"`
}

func (server *Server) UpdateApplicantData(ctx *gin.Context) {
	// ASSIGN REQUEST
	var req ApplicantDataUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// GET APPLICANT BY ID
	applicant, err := server.store.GetApplicantById(ctx, uuid.MustParse(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	}

	// UPDATE PARAMS
	applicantParams := db.UpdateApplicantByIdParams{
		ID:                   uuid.MustParse(req.ID),
		Name:                 sql.NullString{String: req.Name, Valid: true},
		Email:                applicant.Email,
		PhoneNumber:          sql.NullString{String: req.PhoneNumber, Valid: true},
		Photo:                sql.NullString{String: req.Photo, Valid: true},
		Pob:                  sql.NullString{String: req.Pob, Valid: true},
		Dob:                  sql.NullTime{Time: utils.ParseDate(req.Dob), Valid: true},
		SexID:                sql.NullInt64{Int64: int64(req.SexID), Valid: true},
		City:                 sql.NullString{String: req.City, Valid: true},
		Address:              sql.NullString{String: req.Address, Valid: true},
		ReligionID:           sql.NullInt64{Int64: int64(req.ReligionID), Valid: true},
		ApplicationLetter:    sql.NullString{String: req.ApplicationLetter, Valid: true},
		Cv:                   sql.NullString{String: req.Cv, Valid: true},
		EducationCertificate: sql.NullString{String: req.EducationCertificate, Valid: true},
		KtpPhoto:             sql.NullString{String: req.KtpPhoto, Valid: true},
		HealthDocument:       sql.NullString{String: req.HealthDocument, Valid: true},
	}

	// UPDATE BY ID
	applicantNew, err := server.store.UpdateApplicantById(ctx, applicantParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	}
	// ASSIGN RESPONSE
	var resp ApplicantDataUpdateResponse
	resp.ID = applicantNew.ID.String()
	ctx.JSON(http.StatusOK, resp)
}

type ApplicantDataGetRequest struct {
	ID string `uri:"id" binding:"required"`
}

type ApplicantDataGetResponse struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	PhoneNumber          string `json:"phone_number"`
	Email                string `json:"email"`
	Photo                string `json:"photo"`
	Pob                  string `json:"pob"`
	Dob                  string `json:"dob"`
	SexID                int64  `json:"sex_id"`
	City                 string `json:"city"`
	Address              string `json:"address"`
	ReligionID           int64  `json:"religion_id"`
	ApplicationLetter    string `json:"application_letter"`
	Cv                   string `json:"cv"`
	EducationCertificate string `json:"education_certificate"`
	KtpPhoto             string `json:"ktp_photo"`
	HealthDocument       string `json:"health_document"`
}

func (server *Server) GetApplicantById(ctx *gin.Context) {
	var req ApplicantDataGetRequest
	var resp ApplicantDataGetResponse
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	}

	applicantById, err := server.store.GetApplicantById(ctx, uuid.MustParse(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	}

	resp.ID = applicantById.ID.String()
	resp.Name = applicantById.Name.String
	resp.PhoneNumber = applicantById.PhoneNumber.String
	resp.Email = applicantById.Email
	resp.Photo = applicantById.Photo.String
	resp.Pob = applicantById.Pob.String
	resp.Dob = applicantById.Dob.Time.Format("2006-01-02")
	resp.SexID = applicantById.SexID.Int64
	resp.City = applicantById.City.String
	resp.Address = applicantById.Address.String
	resp.ReligionID = applicantById.ReligionID.Int64
	resp.ApplicationLetter = applicantById.ApplicationLetter.String
	resp.Cv = applicantById.Cv.String
	resp.EducationCertificate = applicantById.EducationCertificate.String
	resp.KtpPhoto = applicantById.KtpPhoto.String
	resp.HealthDocument = applicantById.HealthDocument.String

	ctx.JSON(http.StatusOK, resp)
}
