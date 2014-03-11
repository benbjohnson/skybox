package template

import (
	"github.com/skybox/skybox/db"
)

type Template struct {
	Flashes []string
	User    *db.User
	Account *db.Account
}

func New(flashes []string, u *db.User, a *db.Account) *Template {
	return &Template{Flashes: flashes, User: u, Account: a}
}

type AccountTemplate struct {
	*Template
}

func NewAccountTemplate(flashes []string, u *db.User, a *db.Account) *AccountTemplate {
	return &AccountTemplate{New(flashes, u, a)}
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
