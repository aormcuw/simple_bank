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

	//add routes to router
	server.router = router
	return server
}
