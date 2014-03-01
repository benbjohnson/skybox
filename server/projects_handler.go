package server

import (
	"net/http"

	"github.com/benbjohnson/skybox/server/template"
)

type projectsHandler struct {
	handler
}

func newProjectsHandler(s *Server) *projectsHandler {
	return &projectsHandler{handler: handler{server: s}}
}

func (h *projectsHandler) install() {
	h.server.Handle("/projects", h.transact(h.authorize(http.HandlerFunc(h.index)))).Methods("GET")
}

func (h *projectsHandler) index(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	projects, _ := account.Projects()
	t := &template.ProjectsTemplate{template.New(h.session(r), user, account), projects}
	t.Index(w)
}
