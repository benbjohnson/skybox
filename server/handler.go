package server

import (
	"github.com/benbjohnson/skybox/db"
	"github.com/gorilla/sessions"
)

type handler struct {
	server *Server
}

func newHandler(s *Server) {
	return &handler{server: s}
}

func (h *handler) authorize(handler http.Handler) http.Handler {
	return &authorizer{handler}
}

// user returns the logged in user for a given request.
func (h *handler) user(r *http.Request) *db.User {
	session, _ := h.server.store.Get(r, "default")
	id, ok := session.Values["UserID"]
	if !ok {
		return nil
	}
	if id, ok := id.(int); ok {
		u, err := h.server.DB.User(id)
		if err != nil {
			log.Println("[warn] session user not found: %v", err)
		}
		return u
	} else {
		return nil, nil
	}
}

type authorizer struct {
	handler http.Handler
}

func (a *authorizer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}
