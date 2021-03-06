package repository

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/infraestructure/user/entity"
	"time"

	"gorm.io/gorm"
)

type RepositoryUserPostgreSql struct {
	userRepository *gorm.DB
}

func NewRepositoryUserPostgreSql(conn *gorm.DB) *RepositoryUserPostgreSql {
	return &RepositoryUserPostgreSql{
		userRepository: conn.Model(&entity.User{}),
	}
}

func (rup *RepositoryUserPostgreSql) Save(user model.User) error {
	entity := entity.User{Name: user.Name, Password: user.Password, Creation_date: time.Now()}
	result := rup.userRepository.Create(&entity)
	return result.Error
}

func (rup *RepositoryUserPostgreSql) ExistUserName(name string) (bool, error) {
	var user model.User
	var count int64 = 0
	result := rup.userRepository.Where("name = ?", name).Find(&user).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
