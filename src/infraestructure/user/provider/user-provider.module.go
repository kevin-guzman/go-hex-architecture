package provider

import (
	"golang-gingonic-hex-architecture/src/infraestructure/response"
	controller "golang-gingonic-hex-architecture/src/infraestructure/user/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/user/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/user/provider/repository"
	"net/http"
	"sync"

	command "golang-gingonic-hex-architecture/src/application/user/command"
	query "golang-gingonic-hex-architecture/src/application/user/query"
	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/user/provider/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var controllerInstance *controller.ControllerUser
var once sync.Once

func UserProvider(conn *gorm.DB, router *gin.RouterGroup) {
	once.Do(func() {
		repositoryUser := repository.GetRepositoryUser(conn)
		daoUser := dao.GetDaoUser(conn)

		serviceRegisterUser := infraestructureService.GetServiceRegisterUser(*repositoryUser)
		handleRegisterUser := command.NewHandlerRegisterUser(serviceRegisterUser)
		handleListUsers := query.NewHandlerListUsers(*daoUser)

		controllerInstance = controller.NewControllerUser(*handleRegisterUser, *handleListUsers)
		user := router.Group("/user")
		{
			user.POST("/", CreateUser)
			user.GET("/", ListUsers)
		}
	})
}

// Create a new user
// @Summary Create user
// @Schemes http https
// @Description Enpoint to create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body command.CommandRegisterUser true "create user"
// @Success 200 {object} response.ResponseModel
// @Failture 400 {object} response.ResponseModel
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user command.CommandRegisterUser
	if err := c.ShouldBindJSON(&user); err != nil {
		response.SendError(c, "Invalid data: "+err.Error(), "", http.StatusBadRequest)
		return
	}
	msg, err, status := controllerInstance.Create(user)
	if err != nil {
		response.SendError(c, err.Error(), msg, status)
		return
	}
	response.SendSucess(c, msg, status, nil)
}

// Get users
// @Summary Get users
// @Schemes http https
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failture 400 {object} response.ResponseModel
// @Router /user [get]
func ListUsers(c *gin.Context) {
	data := controllerInstance.List()
	response.SendSucess(c, "", 200, data)
}
