package main

import (
	"golang-gingonic-hex-architecture/src/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	infraestructure.InitInfraestructure(server)
	server.Run()
}
