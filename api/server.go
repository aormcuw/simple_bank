package api

import (
	simplebank "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// server serves http request for our banking service
type Server struct {
	store  *simplebank.Store
	router *gin.Engine
}

// newserver create new http server and setup routing
func NewServer(store *simplebank.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// start run the http server on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
