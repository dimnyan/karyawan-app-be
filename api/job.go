package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "karyawan-app-be/db/sqlc"
	"karyawan-app-be/utils"
	"net/http"
	"strconv"
)

type JobCriteriaParams struct {
	Criteria []string `json:"criteria"`
}

type CreateJobRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ClosedAt    string `json:"closed_at" binding:"required"`
	JobCriteriaParams
}

type CreateJobResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	JobCriteriaParams
	ClosedAt  string `json:"closed_at"`
	CreatedAt string `json:"created_at"`
}

func (server *Server) CreateNewJob(ctx *gin.Context) {
	var req CreateJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Error Bind JSON"))
		return
	}

	// cek criteria
	if len(req.Criteria) == 0 {
		ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("criteria is required"))
		return
	}

	// Job: Args
	jobArgs := db.CreateJobParams{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: true},
		ClosedAt:    int32(utils.ParseDateTime(req.ClosedAt)),
	}
	// Job: Create Job
	job, err := server.store.CreateJob(ctx, jobArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Job creation error"))
		return
	}
	// response
	var resp CreateJobResponse
	// Iterate through Request Criteria
	for _, criteria := range req.Criteria {
		criteriaId, err := strconv.Atoi(criteria)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Criteria is invalid"))
			return
		}
		// Job Criteria: Args
		criteriaArgs := db.CreateJobCriteriaParams{
			ID:         uuid.New(),
			JobID:      job.ID,
			CriteriaID: int64(criteriaId),
		}
		// Job Criteria: Save
		_, err = server.store.CreateJobCriteria(ctx, criteriaArgs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Job criteria creation error"))
			return
		}
		resp.Criteria = append(resp.Criteria, criteria)
	}
	resp.ID = job.ID.String()
	resp.Title = job.Title
	resp.Description = job.Description.String
	resp.ClosedAt = strconv.Itoa(int(job.ClosedAt))
	resp.CreatedAt = job.CreatedAt.String()

	ctx.JSON(http.StatusCreated, resp)
}

type GetJobIDRequest struct {
	ID string `uri:"id"`
}

type GetJobCriteriaFromJobIDResponse struct {
	ID       string `json:"id"`
	Criteria int64  `json:"criteria"`
}

type GetJobIDResponse struct {
	ID              string                            `json:"id"`
	Title           string                            `json:"title"`
	Description     string                            `json:"description"`
	JobCriteriaList []GetJobCriteriaFromJobIDResponse `json:"criteria"`
	ClosedAt        int                               `json:"closed_at"`
}

func (server *Server) GetJobByID(ctx *gin.Context) {
	var req GetJobIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Invalid URI"))
		return
	}
	if err := uuid.Validate(req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Invalid Job ID not UUID"))
		return
	}
	// Get Job
	job, err := server.store.GetJobById(ctx, uuid.MustParse(req.ID))
	if len(job) == 0 {
		ctx.JSON(http.StatusNotFound, utils.ErrorMessage("Job not found"))
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorMessage("Job not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Error Retrieving Job"))
		return
	}

	var resp GetJobIDResponse
	// Assign Job Criteria
	var criteriaList []GetJobCriteriaFromJobIDResponse
	if len(job) > 0 {
		for _, j := range job {
			criteriaList = append(criteriaList, GetJobCriteriaFromJobIDResponse{
				ID:       j.ID.String(),
				Criteria: j.CriteriaID,
			})
		}
		resp.ID = job[0].ID.String()
		resp.Title = job[0].Title
		resp.Description = job[0].Description.String
		resp.JobCriteriaList = criteriaList
		resp.ClosedAt = int(job[0].ClosedAt)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) GetJobList(ctx *gin.Context) {
	var list []GetJobIDResponse
	var response GetJobIDResponse
	//var criteriaList []GetJobCriteriaFromJobIDResponse
	jobCritMap := make(map[string]GetJobIDResponse)
	// Get List Jobs
	jobs, err := server.store.GetJobs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Error Retrieving Job List"))
	}
	if len(jobs) == 0 {
		ctx.JSON(http.StatusNotFound, utils.ErrorMessage("Job list is empty"))
		return
	}
	// iterate through jobs
	for _, job := range jobs {
		existedJob, isExist := jobCritMap[job.ID.String()]
		// if job not exist
		if !isExist {
			// input fields
			response.ID = job.ID.String()
			response.Title = job.Title
			response.Description = job.Description.String
			response.ClosedAt = int(job.ClosedAt)

			arg := GetJobCriteriaFromJobIDResponse{
				ID:       job.ID_2.String(),
				Criteria: job.CriteriaID,
			}
			slice := jobCritMap[job.ID_2.String()].JobCriteriaList
			slice = append(slice, arg)
			response.JobCriteriaList = slice
			// assign to map
			jobCritMap[job.ID.String()] = response
		}
		if isExist {
			// input existing fields
			response.ID = existedJob.ID
			response.Title = existedJob.Title
			response.Description = existedJob.Description
			response.ClosedAt = existedJob.ClosedAt
			// append slice
			arg := GetJobCriteriaFromJobIDResponse{
				ID:       job.ID_2.String(),
				Criteria: job.CriteriaID,
			}
			slice := existedJob.JobCriteriaList
			slice = append(slice, arg)
			response.JobCriteriaList = slice
			// assign to map
			jobCritMap[job.ID.String()] = response
		}
	}
	// all maps to output
	for _, innerMap := range jobCritMap {
		// Input innermap to list
		list = append(list, innerMap)
	}
	ctx.JSON(http.StatusOK, list)
}

type DeleteJobRequest struct {
	JobID string `uri:"id"`
}

type DeleteJobResponse struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

func (server *Server) DeleteJob(ctx *gin.Context) {
	var req DeleteJobRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Invalid URI"))
		return
	}
	job, err := server.store.DeleteJob(ctx, uuid.MustParse(req.JobID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorMessage("Job not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Error Deleting Job"))
		return
	}
	_, err = server.store.DeleteJobCriteriaByJobId(ctx, job.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Error Deleting Job Criteria"))
		return
	}
	// Job : Response
	var resp DeleteJobResponse
	resp.ID = job.ID.String()
	resp.Status = "success"
	ctx.JSON(http.StatusOK, resp)
}

type JobIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type UpdateJobRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ClosedAt    int    `json:"closed_at" binding:"required"`
}

type UpdateJobResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ClosedAt    int    `json:"closed_at" binding:"required"`
}

func (server *Server) UpdateJob(ctx *gin.Context) {
	var req UpdateJobRequest
	var idReq JobIDRequest
	var res UpdateJobResponse
	//var criteriaList []GetJobCriteriaFromJobIDResponse

	if err := ctx.BindUri(&idReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Invalid URI"))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(req.Title)
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Invalid JSON body"))
		return
	}
	// Update Job Params
	jobArgs := db.UpdateJobParams{
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: true},
		ClosedAt:    int32(req.ClosedAt),
		ID:          uuid.MustParse(idReq.ID),
	}
	// Update Job
	job, err := server.store.UpdateJob(ctx, jobArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Error Updating Job"))
		return
	}
	// Assign Response
	res.ID = job.ID.String()
	res.Title = job.Title
	res.Description = job.Description.String
	res.ClosedAt = int(job.ClosedAt)
	ctx.JSON(http.StatusOK, res)
}
