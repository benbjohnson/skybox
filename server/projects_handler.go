package server

import (
	"net/http"

	"github.com/benbjohnson/skybox/server/template"
)

type projectsHandler struct {
	handler
}

func newProjectsHandler(s *Server) *projectsHandler {
	return &projectsHandler{handler: handler{s}}
}

func (h *projectsHandler) install() {
	h.server.Handle("/projects", h.authorize(http.HandlerFunc(h.index))).Methods("GET")
}

func (h *projectsHandler) index(w http.ResponseWriter, r *http.Request) {
	// TODO(benbjohnson): Wrap in transaction (db.Transaction, db.RWTransaction).
	user, account := h.auth(r)
	projects, _ := account.Projects()
	t := &template.ProjectsTemplate{template.New(user, account), projects}
	t.Index(w)
}
