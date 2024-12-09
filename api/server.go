package api

import (
	"github.com/gin-gonic/gin"
	db "karyawan-app-be/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20

	// Add routes to router
	router.POST("/api/v1/applicants/register", server.RegisterApplicant)
	router.PUT("/api/v1/applicants/data", server.UpdateApplicantData)
	router.GET("/api/v1/applicants/data/:id", server.GetApplicantById)
	router.POST("/api/v1/auth/login", server.Login)
	router.POST("/api/v1/auth/logout", server.Logout)
	router.POST("/api/v1/auth/checkToken", server.CheckToken)
	router.POST("/api/v1/jobs", server.CreateNewJob)
	router.GET("/api/v1/jobs/:id", server.GetJobByID)
	router.GET("/api/v1/jobs", server.GetJobList)
	router.DELETE("/api/v1/jobs/:id", server.DeleteJob)
	router.PUT("/api/v1/jobs/:id", server.UpdateJob)
	router.POST("/api/v1/jobs/criteria", server.AddJobCriteria)
	router.DELETE("/api/v1/jobs/criteria/:id", server.DeleteJobCriteria)

	//router.GET("/accounts/:id", server.getAccount)
	//router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
