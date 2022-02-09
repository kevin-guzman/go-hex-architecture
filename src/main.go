package main

import (
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "golang-gingonic-hex-architecture/src/docs"

	"golang-gingonic-hex-architecture/src/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// r := server
	// docs.Sp

	infraestructure.InitInfraestructure(server)

	// server.GET("/", HealthCheck)

	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	if err := server.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
