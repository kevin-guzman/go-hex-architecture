package repository

import "golang-gingonic-hex-architecture/src/domain/user/model"

type RepositoryUser interface {
	ExistUserName(name string) (bool, error)
	Save(user model.User) error
}
