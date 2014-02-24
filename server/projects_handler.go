package server

import (
	"github.com/benbjohnson/skybox/db"
)

type projectsHandler struct {
	*handler
}

func newProjectsHandler(s *Server) {
	return &projectsHandler{handler: newHandler(s)}
}

func (h *projectsHandler) install() {
	s.HandleFunc("/projects", h.authorize(h.index)).Methods("GET")
}

func (h *projectsHandler) index(w http.ResponseWriter, r *http.Request) {
	// TODO(benbjohnson): Wrap in transaction (db.Transaction, db.RWTransaction).
	user, _ := h.user(r)
	account, _ := user.Account()
	projects, _ := account.Projects()
	t := templates.NewProjectTemplate(user, account, projects)
	t.Index(w)
}
