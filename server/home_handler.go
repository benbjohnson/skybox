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
}

func (h *homeHandler) index(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	template.New(user, account).Index(w)
}
