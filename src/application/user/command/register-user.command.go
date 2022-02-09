package command

type CommandRegisterUser struct {
	Name         string `json:"Name"`
	CreationDate string `json:"CreationDate"`
	Password     string `json:"Password" minLength:"6"`
}
