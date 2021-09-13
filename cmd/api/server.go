package main

import "github.com/gin-gonic/gin"

type Server struct {
}

func MustNewServer() Server {
	return Server{}
}

func (s Server) Listen() *gin.Engine {
	return gin.Default()
}
