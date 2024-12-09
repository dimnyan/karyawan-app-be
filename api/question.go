package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "karyawan-app-be/db/sqlc"
	"karyawan-app-be/utils"
	"net/http"
	"strings"
)

type CreateTestQuestionRequest struct {
	JobID    string `json:"job_id" required:"true"`
	Question string `json:"question" required:"true"`
	ChoiceA  string `json:"choice_a" required:"true"`
	ChoiceB  string `json:"choice_b" required:"true"`
	ChoiceC  string `json:"choice_c" required:"true"`
	ChoiceD  string `json:"choice_d" required:"true"`
	Answer   string `json:"answer" required:"true"`
}
type CreateTestQuestionResponse struct {
	ID string `json:"id"`
	CreateTestQuestionRequest
}

func (server *Server) CreateQuestion(ctx *gin.Context) {
	var req CreateTestQuestionRequest
	var resp CreateTestQuestionResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("Error JSON"))
	}

	args := db.CreateTestQuestionParams{
		ID:       uuid.New(),
		JobID:    uuid.MustParse(req.JobID),
		Question: req.Question,
		ChoiceA:  req.ChoiceA,
		ChoiceB:  req.ChoiceB,
		ChoiceC:  req.ChoiceC,
		ChoiceD:  req.ChoiceD,
		Answer:   req.Answer,
	}
	_, err := server.store.CreateTestQuestion(ctx, args)
	if err != nil {
		if strings.Contains(err.Error(), "m_test_questions_job_id_fkey") {
			ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("Job ID not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Cannot create question"))
		return
	}
	resp.ID = args.ID.String()
	resp.CreateTestQuestionRequest = req
	ctx.JSON(http.StatusCreated, resp)
}

func (server *Server) GetQuestionList(ctx *gin.Context) {
	questions, err := server.store.GetTestQuestions(ctx)
	if len(questions) == 0 {
		ctx.JSON(http.StatusNotFound, utils.ErrorMessage("No questions found"))
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Cannot get test questions"))
		return
	}
	ctx.JSON(http.StatusOK, questions)
}

type GetQuestionByIdRequest struct {
	ID string `uri:"id" required:"true"`
}

func (server *Server) GetQuestionById(ctx *gin.Context) {
	var req GetQuestionByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("URI ID Failed"))
		return
	}

	question, err := server.store.GetQuestionById(ctx, uuid.MustParse(req.ID))
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			ctx.JSON(http.StatusNotFound, utils.ErrorMessage("Question not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Cannot get test question"))
		return
	}
	ctx.JSON(http.StatusOK, question)
}

type GetQuestionByJobIdResponse struct {
	ID string `uri:"id"`
}

func (server *Server) GetQuestionByJobId(ctx *gin.Context) {
	var req GetQuestionByJobIdResponse
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("URI ID Failed"))
		return
	}
	if err := uuid.Validate(req.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("ID is not valid UUID"))
		return
	}
	questions, err := server.store.GetQuestionsByJobId(ctx, uuid.MustParse(req.ID))
	if len(questions) == 0 {
		ctx.JSON(http.StatusNotFound, utils.ErrorMessage("Question not found"))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Cannot get test questions"))
		return
	}
	ctx.JSON(http.StatusOK, questions)
}

type UriId struct {
	ID string `uri:"id" required:"true"`
}

type UpdateQuestionByIDRequest struct {
	UriId
	JobID    string `json:"job_id" required:"true"`
	Question string `json:"question" required:"true"`
	ChoiceA  string `json:"choice_a" required:"true"`
	ChoiceB  string `json:"choice_b" required:"true"`
	ChoiceC  string `json:"choice_c" required:"true"`
	ChoiceD  string `json:"choice_d" required:"true"`
	Answer   string `json:"answer" required:"true"`
}

func (server *Server) UpdateQuestionByID(ctx *gin.Context) {
	var reqId UriId
	var req UpdateQuestionByIDRequest
	if err := ctx.ShouldBindUri(&reqId); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("URI ID Failed"))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorMessage("JSON Failed"))
		return
	}
	// get question by id
	questionId, err := server.store.GetQuestionById(ctx, uuid.MustParse(reqId.ID))
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			ctx.JSON(http.StatusNotFound, utils.ErrorMessage("job id not found"))
			return
		}
	}
	args := db.UpdateTestQuestionParams{
		ID:       uuid.MustParse(reqId.ID),
		JobID:    questionId.JobID,
		Question: req.Question,
		ChoiceA:  req.ChoiceA,
		ChoiceB:  req.ChoiceB,
		ChoiceC:  req.ChoiceC,
		ChoiceD:  req.ChoiceD,
		Answer:   req.Answer,
	}
	question, err := server.store.UpdateTestQuestion(ctx, args)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			ctx.JSON(http.StatusNotFound, utils.ErrorMessage("Question not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Cannot update question"))
		return
	}
	ctx.JSON(http.StatusOK, question)
}
