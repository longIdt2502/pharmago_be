package api

import (
	"github.com/gin-gonic/gin"
	"github.com/longIdt2502/pharmago_be/api/login"
	"github.com/longIdt2502/pharmago_be/api/register"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	// Public router
	publicRouter := router.Group("/")
	login.NewLoginRouter(publicRouter, server.store, server.tokenMaker)
	register.NewRegisterRouter(publicRouter, server.store, server.tokenMaker)

	server.router = router
}
