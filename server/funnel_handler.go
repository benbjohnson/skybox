package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/skybox/skybox/db"
	"github.com/skybox/skybox/server/template"
)

type FunnelHandler struct {
	*Handler
}

func NewFunnelHandler(parent *Handler) *FunnelHandler {
	h := &FunnelHandler{parent}
	h.HandleFunc("/funnels", h.index).Methods("GET")
	h.HandleFunc("/funnels/{id}", h.show).Methods("GET")
	h.HandleFunc("/funnels/{id}/edit", h.edit).Methods("GET")
	h.HandleFunc("/funnels/{id}", h.save).Methods("POST")
	return h
}

func (h *FunnelHandler) index(w http.ResponseWriter, r *http.Request) {
	h.db.With(func(tx *db.Tx) error {
		user, account := h.Authenticate(tx, r)
		if user == nil {
			h.Unauthorized(w, r)
			return ErrUnauthorized
		}
		funnels, _ := account.Funnels()
		t := &template.FunnelsTemplate{
			Template: template.New(h.Flashes(w, r), user, account),
			Funnels:  funnels,
		}
		t.Index(w)
		return nil
	})
}

func (h *FunnelHandler) show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h.db.With(func(tx *db.Tx) error {
		user, account := h.Authenticate(tx, r)
		if user == nil {
			h.Unauthorized(w, r)
			return ErrUnauthorized
		}

		// Find funnel by id.
		id, _ := strconv.Atoi(vars["id"])
		f, err := account.Funnel(id)
		if err != nil {
			h.NotFound(tx, w, r)
			return err
		}

		// Execute the funnel query.
		result, err := f.Query()
		if err != nil {
			http.Error(w, "funnel query: "+err.Error(), http.StatusInternalServerError)
			return err
		}

		t := &template.FunnelTemplate{
			Template: template.New(h.Flashes(w, r), user, account),
			Funnel:   f,
			Result:   result,
		}
		t.Show(w)
		return nil
	})
}

func (h *FunnelHandler) edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h.db.With(func(tx *db.Tx) error {
		user, account := h.Authenticate(tx, r)
		if user == nil {
			h.Unauthorized(w, r)
			return ErrUnauthorized
		}

		var f *db.Funnel
		if vars["id"] == "0" {
			f = &db.Funnel{
				Steps: []*db.FunnelStep{{}},
			}
		} else {
			var err error
			id, _ := strconv.Atoi(vars["id"])
			if f, err = account.Funnel(id); err != nil {
				h.NotFound(tx, w, r)
				return err
			}
		}

		// Find all resources.
		resources, err := account.Resources()
		if err != nil {
			http.Error(w, "resources: "+err.Error(), http.StatusInternalServerError)
			return err
		}

		t := &template.FunnelTemplate{
			Template:  template.New(h.Flashes(w, r), user, account),
			Funnel:    f,
			Resources: resources,
		}
		t.Edit(w)
		return nil
	})
}

func (h *FunnelHandler) save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h.db.Do(func(tx *db.Tx) error {
		user, account := h.Authenticate(tx, r)
		if user == nil {
			h.Unauthorized(w, r)
			return ErrUnauthorized
		}
		session := h.Session(r)

		// Find funnel.
		var f = &db.Funnel{}
		if id, _ := strconv.Atoi(vars["id"]); id != 0 {
			var err error
			f, err = account.Funnel(id)
			if err != nil {
				h.NotFound(tx, w, r)
				return err
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
			t := &template.FunnelTemplate{
				Template: template.New([]string{err.Error()}, user, account),
				Funnel:   f,
			}
			t.Edit(w)
			return err
		}

		session.AddFlash("funnel successfully saved")
		session.Save(r, w)
		http.Redirect(w, r, "/funnels", http.StatusFound)
		return nil
	})
}
