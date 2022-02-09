package provider

import (
	controller "golang-gingonic-hex-architecture/src/infraestructure/user/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/user/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/user/provider/repository"
	"time"

	// domainService "golang-gingonic-hex-architecture/src/domain/user/service"

	command "golang-gingonic-hex-architecture/src/application/user/command"
	query "golang-gingonic-hex-architecture/src/application/user/query"
	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/user/provider/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserProvider(conn *gorm.DB, router *gin.Engine) {
	repositoryUser := repository.GetRepositoryUser(conn)
	daoUser := dao.GetDaoUser(conn)

	serviceRegisterUser := infraestructureService.GetServiceRegisterUser(*repositoryUser)
	handleRegisterUser := command.NewHandlerRegisterUser(serviceRegisterUser)
	handleListUsers := query.NewHandlerListUsers(*daoUser)

	controller := controller.NewControllerUser(*handleRegisterUser, *handleListUsers)
	user := router.Group("/user")
	{
		user.POST("/", func(c *gin.Context) {
			command := command.CommandRegisterUser{Name: "wey", Password: "12345543", CreationDate: time.Now()}
			msg, err, status := controller.Create(command)
			c.JSON(status, gin.H{
				"message": msg,
				"error":   err.Error(),
			})
		})
		user.GET("/", func(c *gin.Context) {
			data := controller.List()
			c.JSON(200, gin.H{
				"data": data,
			})
		})
	}
}
