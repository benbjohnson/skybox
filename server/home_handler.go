package server

import (
	"net/http"

	"github.com/benbjohnson/skybox/server/template"
)

type homeHandler struct {
	handler
}

func (h *homeHandler) install() {
	h.server.Handle("/", h.transact(h.authorize(http.HandlerFunc(h.index)))).Methods("GET")
	h.server.Handle("/login", http.HandlerFunc(h.login)).Methods("GET")
	h.server.Handle("/login", h.transact(http.HandlerFunc(h.doLogin))).Methods("POST")
	h.server.Handle("/logout", http.HandlerFunc(h.logout))
}

func (h *homeHandler) index(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	template.New(h.session(r), user, account).Index(w)
}

func (h *homeHandler) login(w http.ResponseWriter, r *http.Request) {
	session := h.session(r)
	if _, ok := session.Values["UserID"]; ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	template.New(session, nil, nil).Login(w)
}

func (h *homeHandler) doLogin(w http.ResponseWriter, r *http.Request) {
	txn, session := h.transaction(r), h.session(r)

	// Retrieve user.
	user, err := txn.UserByUsername(r.FormValue("username"))
	if err != nil {
		session.AddFlash(err.Error())
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Authenticate user.
	if err := user.Authenticate(r.FormValue("password")); err != nil {
		session.AddFlash(err.Error())
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Set session's User ID.
	session.Values["UserID"] = user.ID()

	// Redirect to home page.
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *homeHandler) logout(w http.ResponseWriter, r *http.Request) {
	// TODO(benbjohnson): Remove User ID from session.
}
