package api

import (
	db "github.com/cijin/go-start/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Serves HTTP requests to the service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router
	return server
}

// Start server on a specific address
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err}
}
