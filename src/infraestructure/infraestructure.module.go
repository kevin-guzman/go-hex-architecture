package infraestructure

import (
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider"
	"sync"

	"github.com/gin-gonic/gin"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB
var once sync.Once

func InitInfraestructure(router *gin.Engine) {

	once.Do(func() {
		dsn := "host=localhost user=go password=go dbname=golanghex port=5432 sslmode=disable"
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// defer conn.Cl()

		if err != nil {
			log.Fatal("Error with db connection", err)
		}

		dbConnection = conn
		provider.UserProvider(dbConnection, router)
	})
}
