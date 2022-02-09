package repository

// "golang-gingonic-hex-architecture/src/domain/user/port/repository"
import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/infraestructure/user/entity"

	toy "github.com/bigpigeon/toyorm"
)

type RepositoryUserPostgreSql struct {
	userRepository *toy.ToyBrick
}

func NewRepositoryUserPostgreSql(conn toy.Toy) *RepositoryUserPostgreSql {
	return &RepositoryUserPostgreSql{
		userRepository: conn.Model(&entity.User{}),
	}
}

func (rup *RepositoryUserPostgreSql) Save(user model.User) error {
	_, err := rup.userRepository.Insert(&user)
	return err
}

func (rup *RepositoryUserPostgreSql) ExistUserName(name string) (bool, error) {
	var user bool = false
	finded, err := rup.userRepository.Where(toy.ExprEqual, "name", name).Find(&user)
	if err != nil {
		return true, err
	}
	if resultErr := finded.Err(); resultErr != nil {
		return true, resultErr
	}
	return user == false, nil
}
