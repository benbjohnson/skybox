package server

import (
	"net/http"

	"github.com/skybox/skybox/db"
	"github.com/skybox/skybox/server/template"
)

type AccountHandler struct {
	*Handler
}

func NewAccountHandler(parent *Handler) *AccountHandler {
	h := &AccountHandler{Handler: parent}
	h.Handle("/account", http.HandlerFunc(h.show)).Methods("GET")
	return h
}

func (h *AccountHandler) show(w http.ResponseWriter, r *http.Request) {
	h.db.With(func(tx *db.Tx) error {
		user, account := h.Authenticate(tx, r)
		if user == nil {
			h.Unauthorized(w, r)
			return nil
		}

		template.NewAccountTemplate(h.Session(r), user, account).Show(w)
		return nil
	})
}
