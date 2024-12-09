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
		ID:      uuid.New(),
		JobID:   uuid.MustParse(req.JobID),
		ChoiceA: req.ChoiceA,
		ChoiceB: req.ChoiceB,
		ChoiceC: req.ChoiceC,
		ChoiceD: req.ChoiceD,
		Answer:  req.Answer,
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
