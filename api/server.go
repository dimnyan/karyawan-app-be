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
	//router.GET("/accounts/:id", server.getAccount)
	//router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
