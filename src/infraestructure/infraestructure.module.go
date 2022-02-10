package infraestructure

import (
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider"
	"os"
	"sync"

	"github.com/gin-gonic/gin"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB
var once sync.Once

func InitInfraestructure(router *gin.RouterGroup) {

	once.Do(func() {
		DATABSE_STRING_CONNECTION := os.Getenv("DB")
		conn, err := gorm.Open(postgres.Open(DATABSE_STRING_CONNECTION), &gorm.Config{})
		conn = conn.Debug()

		if err != nil {
			log.Fatal("Error with db connection", err)
		}

		dbConnection = conn
		provider.UserProvider(dbConnection, router)
	})
}
