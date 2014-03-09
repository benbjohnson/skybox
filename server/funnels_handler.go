package server

import (
	"net/http"
	"strconv"

	"github.com/benbjohnson/skybox/db"
	"github.com/benbjohnson/skybox/server/template"
	"github.com/gorilla/mux"
)

type funnelsHandler struct {
	handler
}

func newFunnelsHandler(s *Server) *funnelsHandler {
	return &funnelsHandler{handler: handler{server: s}}
}

func (h *funnelsHandler) install() {
	h.server.Handle("/funnels/{id}", h.transact(h.authorize(http.HandlerFunc(h.edit)))).Methods("GET")
}

func (h *funnelsHandler) edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, account := h.auth(r)

	var f *db.Funnel
	if vars["id"] == "new" {
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

	t := &template.FunnelTemplate{template.New(h.session(r), user, account), f}
	t.Edit(w)
}
