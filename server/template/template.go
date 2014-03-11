package template

import (
	"github.com/gorilla/sessions"
	"github.com/skybox/skybox/db"
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

func NewAccountTemplate(s *sessions.Session, u *db.User, a *db.Account) *AccountTemplate {
	return &AccountTemplate{New(s, u, a)}
}

type FunnelTemplate struct {
	*Template
	Funnel    *db.Funnel
	Result    *db.FunnelResult
	Resources []string
}

type FunnelsTemplate struct {
	*Template
	Funnels []*db.Funnel
}

// iif returns trueValue if condition is true. Otherwise returns falseValue.
func iif(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}
