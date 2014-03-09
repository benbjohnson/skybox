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

type AccountTemplate struct {
	*Template
}

type FunnelTemplate struct {
	*Template
	Funnel *db.Funnel
}

type FunnelsTemplate struct {
	*Template
	Funnels []*db.Funnel
}
