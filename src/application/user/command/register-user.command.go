package command

import "time"

type CommandRegisterUser struct {
	Name, Password string
	CreationDate   time.Time
}
