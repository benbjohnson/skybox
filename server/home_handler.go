package server

import (
	"net/http"

	"github.com/benbjohnson/skybox/db"
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
	h.server.Handle("/signup", http.HandlerFunc(h.signup)).Methods("GET")
	h.server.Handle("/signup", h.rwtransact(http.HandlerFunc(h.doSignup))).Methods("POST")
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
	tx, session := h.transaction(r), h.session(r)

	// Retrieve user.
	user, err := tx.UserByEmail(r.FormValue("email"))
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
	session.Save(r, w)

	// Redirect to home page.
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *homeHandler) logout(w http.ResponseWriter, r *http.Request) {
	session := h.session(r)
	delete(session.Values, "UserID")
	session.Save(r, w)
}

func (h *homeHandler) signup(w http.ResponseWriter, r *http.Request) {
	template.New(h.session(r), nil, nil).Signup(w)
}

func (h *homeHandler) doSignup(w http.ResponseWriter, r *http.Request) {
	tx, session := h.transaction(r), h.session(r)

	// Create a new account.
	account := &db.Account{}
	if err := tx.CreateAccount(account); err != nil {
		tx.Rollback()
		session.AddFlash(err.Error())
		http.Redirect(w, r, r.URL.Path, http.StatusInternalServerError)
		return
	}

	// Create a new account.
	user := &db.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Create default project.
	if err := account.CreateProject(&db.Project{Name: "Default Project"}); err != nil {
		tx.Rollback()
		session.AddFlash(err.Error())
		http.Redirect(w, r, r.URL.Path, http.StatusInternalServerError)
		return
	}

	// Create user.
	if err := account.CreateUser(user); err != nil {
		tx.Rollback()
		session.AddFlash(err.Error())
		http.Redirect(w, r, r.URL.Path, http.StatusInternalServerError)
		return
	}

	// Update session.
	session.Values["UserID"] = user.ID()
	session.Save(r, w)

	// Redirect to home page.
	http.Redirect(w, r, "/", http.StatusFound)
}
