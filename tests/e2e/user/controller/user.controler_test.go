package controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"golang-gingonic-hex-architecture/src/application/user/command"
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure"
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider"
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider/dao"

	"golang-gingonic-hex-architecture/src/infraestructure/user/provider/repository"

	"golang-gingonic-hex-architecture/tests/utils"
	"golang-gingonic-hex-architecture/tests/utils/mocks"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	interfaceDao "golang-gingonic-hex-architecture/src/domain/user/port/dao"
	interfaceRepository "golang-gingonic-hex-architecture/src/domain/user/port/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	repositoryUser *mocks.MockRepositoryUser
	daoUser        *mocks.MockDaoUser
	server         *gin.Engine
	context        *gin.Context
)

func init() {
	repositoryUser = &mocks.MockRepositoryUser{}
	daoUser = &mocks.MockDaoUser{}
	dao.GetDaoUser = func(conn *gorm.DB) *interfaceDao.DaoUser {
		ru := daoUser
		iru := interfaceDao.DaoUser(ru)
		_, _ = iru, ru
		return &iru
	}
	repository.GetRepositoryUser = func(conn *gorm.DB) *interfaceRepository.RepositoryUser {
		ru := repositoryUser
		iru := interfaceRepository.RepositoryUser(ru)
		_, _ = iru, ru
		return &iru
	}
	infraestructure.InitInfraestructure = func(router *gin.RouterGroup) {
		/* Mocking the databse connection to do it as a null or mocked connection */
		var db *sql.DB
		db, _, _ = sqlmock.New()
		dial := postgres.New(postgres.Config{
			DriverName: "postgres",
			Conn:       db,
		})
		conn, _ := gorm.Open(dial, &gorm.Config{
			SkipDefaultTransaction: true,
		})
		provider.UserProvider(conn, router)
	}

	gin.SetMode(gin.TestMode)
	context, server = gin.CreateTestContext(httptest.NewRecorder())
	CONTEXT_PATH := os.Getenv("CONTEXT_PATH")
	if CONTEXT_PATH == "" {
		CONTEXT_PATH = "api/v1"
	}
	path := server.Group(CONTEXT_PATH)
	infraestructure.InitInfraestructure(path)
}

func TestListUsers(t *testing.T) {
	assert := require.New(t)

	users := []*dto.UserDto{{Name: "name", Creation_date: time.Now()}}
	var expected, got gin.H
	utils.JSONParse(
		gin.H{
			"data":    users,
			"status":  200,
			"error":   "",
			"message": "",
		},
		&expected,
	)
	daoUser.On("List").Return(users)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/", nil)
	server.ServeHTTP(w, req)

	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		log.Fatal("Error on unmarsall reponse ", err)
	}

	assert.Equal(200, w.Code)
	assert.Equal(expected, got)
	context.Done()
}

func TestFailCreateUserForShortPassword(t *testing.T) {
	assert := require.New(t)

	var expected, got gin.H
	user := command.CommandRegisterUser{Name: "Juan", Password: "12", CreationDate: "2022-02-11T13:47:15.098Z"}
	userBytes, _ := json.Marshal(user)
	const message string = "The leng of the password is incorrect"
	utils.JSONParse(
		gin.H{
			"data":    nil,
			"status":  500,
			"error":   message,
			"message": "",
		},
		&expected,
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/", bytes.NewBuffer(userBytes))
	server.ServeHTTP(w, req)

	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		log.Fatal("Error on unmarsall reponse pas", err)
	}

	assert.Equal(w.Code, 500)
	assert.Equal(expected, got)
	assert.Equal(message, got["error"])
}

func TestFailCreateUserForExistentUser(t *testing.T) {
	assert := require.New(t)

	var expected, got gin.H
	user := command.CommandRegisterUser{Name: "Juan", Password: "12267gdgweg3", CreationDate: "2022-02-11T13:47:15.098Z"}
	userBytes, _ := json.Marshal(user)
	var message string = "The username " + user.Name + " already exist"
	utils.JSONParse(
		gin.H{
			"data":    nil,
			"status":  500,
			"error":   message,
			"message": "",
		},
		&expected,
	)
	repositoryUser.On("ExistUserName").Return(true, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/", bytes.NewBuffer(userBytes))
	server.ServeHTTP(w, req)

	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		log.Fatal("Error on unmarsall reponse ex", err)
	}

	assert.Equal(w.Code, 500)
	assert.Equal(expected, got)
	assert.Equal(message, got["error"])
}
