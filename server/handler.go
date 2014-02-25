package server

import (
	"log"
	"net/http"

	"github.com/benbjohnson/skybox/db"
)

type handler struct {
	server *Server
}

func (h *handler) authorize(handler http.Handler) http.Handler {
	return &authorizer{handler}
}

// auth returns the logged in user and account for a given request.
func (h *handler) auth(r *http.Request) (*db.User, *db.Account) {
	session, _ := h.server.store.Get(r, "default")
	id, ok := session.Values["UserID"]
	if !ok {
		return nil, nil
	}
	if id, ok := id.(int); ok {
		u, err := h.server.DB.User(id)
		if err != nil {
			log.Println("[warn] session user not found: %v", err)
		}
		a, _ := u.Account()
		return u, a
	}
	return nil, nil
}

type authorizer struct {
	handler http.Handler
}

func (a *authorizer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO(benbjohnson): Check if there is a user id.
}
