package server

import (
	"net/http"
	"strconv"

	"github.com/benbjohnson/skybox/db"
	"github.com/benbjohnson/skybox/server/template"
	"github.com/gorilla/mux"
)

type projectsHandler struct {
	handler
}

func newProjectsHandler(s *Server) *projectsHandler {
	return &projectsHandler{handler: handler{server: s}}
}

func (h *projectsHandler) install() {
	h.server.Handle("/projects", h.transact(h.authorize(http.HandlerFunc(h.index)))).Methods("GET")
	h.server.Handle("/projects/{id}", h.transact(h.authorize(http.HandlerFunc(h.edit)))).Methods("GET")
	h.server.Handle("/projects/{id}", h.rwtransact(h.authorize(http.HandlerFunc(h.save)))).Methods("POST")
}

func (h *projectsHandler) index(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	projects, _ := account.Projects()
	t := &template.ProjectsTemplate{template.New(h.session(r), user, account), projects}
	t.Index(w)
}

func (h *projectsHandler) edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, account := h.auth(r)

	var project *db.Project
	var err error
	if vars["id"] == "new" {
		project = &db.Project{}
	} else {
		id, _ := strconv.Atoi(vars["id"])
		if project, err = account.Project(id); err != nil {
			h.notFound(w, r)
			return
		}
	}

	t := &template.ProjectTemplate{template.New(h.session(r), user, account), project}
	t.Edit(w)
}

func (h *projectsHandler) save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, account := h.auth(r)
	session := h.session(r)
	tx := h.transaction(r)

	// Find project.
	var project = &db.Project{}
	if id, _ := strconv.Atoi(vars["id"]); id != 0 {
		var err error
		project, err = account.Project(id)
		if err != nil {
			h.notFound(w, r)
			return
		}
	}

	// Update values.
	project.Name = r.FormValue("name")

	// Save.
	var err error
	if project.ID() == 0 {
		err = account.CreateProject(project)
	} else {
		err = project.Save()
	}

	if err != nil {
		tx.Rollback()
		session.AddFlash(err.Error())
		session.Save(r, w)
		t := &template.ProjectTemplate{template.New(session, user, account), project}
		t.Edit(w)
		return
	}

	session.AddFlash("project successfully saved")
	session.Save(r, w)
	http.Redirect(w, r, "/projects", http.StatusFound)
}
