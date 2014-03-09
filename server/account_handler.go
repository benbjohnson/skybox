package server

import (
	"net/http"

	"github.com/benbjohnson/skybox/server/template"
)

type accountHandler struct {
	handler
}

func newAccountHandler(s *Server) *accountHandler {
	return &accountHandler{handler: handler{server: s}}
}

func (h *accountHandler) install() {
	h.server.Handle("/account", h.transact(h.authorize(http.HandlerFunc(h.show)))).Methods("GET")
}

func (h *accountHandler) show(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	t := &template.AccountTemplate{template.New(h.session(r), user, account)}
	t.Show(w)
}
