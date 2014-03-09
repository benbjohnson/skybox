package template

import (
	"github.com/benbjohnson/skybox/db"
	"github.com/gorilla/sessions"
)

type Template struct {
	Session *sessions.Session
	User    *db.User
	Account *db.Account
}

func New(s *sessions.Session, u *db.User, a *db.Account) *Template {
	return &Template{Session: s, User: u, Account: a}
}

type ProjectTemplate struct {
	*Template
	Project *db.Project
}

type ProjectsTemplate struct {
	*Template
	Projects []*db.Project
}

type FunnelTemplate struct {
	*Template
	Funnel *db.Funnel
}
