package infraestructure

import (
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider"
	"sync"

	"github.com/gin-gonic/gin"

	"log"

	toy "github.com/bigpigeon/toyorm"
)

var dbConnection *toy.Toy
var once sync.Once

func InitInfraestructure(router *gin.Engine) {

	once.Do(func() {
		conn, err := toy.Open("postgres", "user=postgres dbname=toyorm_example sslmode=disable")
		if err != nil {
			log.Fatal("Error with db connection", err)
		}
		dbConnection = conn

		provider.UserProvider(*dbConnection, router)
	})
}
