package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/skybox/skybox/db"
	"github.com/skybox/skybox/server/template"
)

type funnelsHandler struct {
	handler
}

func newFunnelsHandler(s *Server) *funnelsHandler {
	return &funnelsHandler{handler: handler{server: s}}
}

func (h *funnelsHandler) install() {
	h.server.Handle("/funnels", h.transact(h.authorize(http.HandlerFunc(h.index)))).Methods("GET")
	h.server.Handle("/funnels/{id}", h.transact(h.authorize(http.HandlerFunc(h.show)))).Methods("GET")
	h.server.Handle("/funnels/{id}/edit", h.transact(h.authorize(http.HandlerFunc(h.edit)))).Methods("GET")
	h.server.Handle("/funnels/{id}", h.rwtransact(h.authorize(http.HandlerFunc(h.save)))).Methods("POST")
}

func (h *funnelsHandler) index(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	funnels, _ := account.Funnels()
	t := &template.FunnelsTemplate{template.New(h.session(r), user, account), funnels}
	t.Index(w)
}

func (h *funnelsHandler) show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, account := h.auth(r)

	// Find funnel by id.
	id, _ := strconv.Atoi(vars["id"])
	f, err := account.Funnel(id)
	if err != nil {
		h.notFound(w, r)
		return
	}

	// Execute the funnel query.
	result, err := f.Query()
	if err != nil {
		http.Error(w, "funnel query: "+err.Error(), http.StatusInternalServerError)
		return
	}

	t := &template.FunnelTemplate{
		Template:     template.New(h.session(r), user, account),
		Funnel:       f,
		FunnelResult: result,
	}
	t.Show(w)
}

func (h *funnelsHandler) edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, account := h.auth(r)

	var f *db.Funnel
	if vars["id"] == "0" {
		f = &db.Funnel{
			Steps: []*db.FunnelStep{{}},
		}
	} else {
		var err error
		id, _ := strconv.Atoi(vars["id"])
		if f, err = account.Funnel(id); err != nil {
			h.notFound(w, r)
			return
		}
	}

	// Find all resources.
	resources, err := account.Resources()
	if err != nil {
		http.Error(w, "resources: "+err.Error(), http.StatusInternalServerError)
		return
	}

	t := &template.FunnelTemplate{
		Template:  template.New(h.session(r), user, account),
		Funnel:    f,
		Resources: resources,
	}
	t.Edit(w)
}

func (h *funnelsHandler) save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, account := h.auth(r)
	session := h.session(r)
	tx := h.transaction(r)

	// Find funnel.
	var f = &db.Funnel{}
	if id, _ := strconv.Atoi(vars["id"]); id != 0 {
		var err error
		f, err = account.Funnel(id)
		if err != nil {
			h.notFound(w, r)
			return
		}
	}

	// Update attributes.
	f.Name = r.FormValue("name")

	// Create steps.
	f.Steps = make([]*db.FunnelStep, 0)
	for {
		condition := r.FormValue(fmt.Sprintf("step[%d].condition", len(f.Steps)))
		if condition == "" {
			break
		}
		f.Steps = append(f.Steps, &db.FunnelStep{Condition: condition})
	}

	// Save.
	var err error
	if f.ID() == 0 {
		err = account.CreateFunnel(f)
	} else {
		err = f.Save()
	}

	if err != nil {
		tx.Rollback()
		session.AddFlash(err.Error())
		session.Save(r, w)
		t := &template.FunnelTemplate{
			Template: template.New(session, user, account),
			Funnel:   f,
		}
		t.Edit(w)
		return
	}

	session.AddFlash("funnel successfully saved")
	session.Save(r, w)
	http.Redirect(w, r, "/funnels", http.StatusFound)
}
