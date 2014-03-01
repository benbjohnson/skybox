package server

import (
	"net/http"

	"github.com/benbjohnson/skybox/server/template"
)

type homeHandler struct {
	handler
}

func (h *homeHandler) install() {
	h.server.Handle("/", http.HandlerFunc(h.index)).Methods("GET")
	h.server.Handle("/login", http.HandlerFunc(h.login)).Methods("GET")
	h.server.Handle("/logout", http.HandlerFunc(h.logout))
}

func (h *homeHandler) index(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	template.New(user, account).Index(w)
}

func (h *homeHandler) login(w http.ResponseWriter, r *http.Request) {
	template.New(nil, nil).Login(w)
}

func (h *homeHandler) logout(w http.ResponseWriter, r *http.Request) {
	// TODO(benbjohnson): Kill session.
}
