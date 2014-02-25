package template

import (
	"github.com/benbjohnson/skybox/db"
)

type Template struct {
	User    *db.User
	Account *db.Account
}

func New(u *db.User, a *db.Account) *Template {
	return &Template{User: u, Account: a}
}

type ProjectTemplate struct {
	*Template
	Project *db.Project
}

type ProjectsTemplate struct {
	*Template
	Projects []*db.Project
}
