package server

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/skybox/skybox/db"
	"github.com/skybox/skybox/server/template"
)

type RootHandler struct {
	*Handler
}

func NewRootHandler(parent *Handler) *RootHandler {
	h := &RootHandler{Handler: parent}
	h.HandleFunc("/", h.index).Methods("GET")
	h.HandleFunc("/assets/{filename}", h.asset).Methods("GET")
	h.HandleFunc("/skybox.js", h.skyboxjs).Methods("GET")
	h.HandleFunc("/login", h.login).Methods("GET")
	h.HandleFunc("/login", h.doLogin).Methods("POST")
	h.HandleFunc("/logout", h.logout)
	h.HandleFunc("/signup", h.signup).Methods("GET")
	h.HandleFunc("/signup", h.doSignup).Methods("POST")
	return h
}

func (h *RootHandler) index(w http.ResponseWriter, r *http.Request) {
	h.db.With(func(tx *db.Tx) error {
		user, account := h.Authenticate(tx, r)
		if user == nil {
			template.New(h.Session(r), user, account).Index(w)
		} else {
			template.New(h.Session(r), user, account).Dashboard(w)
		}
		return nil
	})
}

// asset retrieves static files in the "assets" folder.
func (h *RootHandler) asset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, err := Asset(vars["filename"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch path.Ext(vars["filename"]) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	}
	w.Write(b)
}

// skyboxjs retrieves the skybox.js static file.
func (h *RootHandler) skyboxjs(w http.ResponseWriter, r *http.Request) {
	b, _ := Asset("skybox.js")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
}

func (h *RootHandler) login(w http.ResponseWriter, r *http.Request) {
	h.db.With(func(tx *db.Tx) error {
		user, _ := h.Authenticate(tx, r)
		if user != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}
		template.New(h.Session(r), nil, nil).Login(w)
		return nil
	})
}

func (h *RootHandler) doLogin(w http.ResponseWriter, r *http.Request) {
	session := h.Session(r)

	h.db.With(func(tx *db.Tx) error {
		user, _ := h.Authenticate(tx, r)
		if user != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}

		// Retrieve user.
		user, err := tx.UserByEmail(r.FormValue("email"))
		if err != nil {
			session.AddFlash(err.Error())
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
			return nil
		}

		// Authenticate user.
		if err := user.Authenticate(r.FormValue("password")); err != nil {
			session.AddFlash(err.Error())
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
			return nil
		}

		// Set session's User ID.
		session.Values["UserID"] = user.ID()
		session.Save(r, w)

		// Redirect to home page.
		http.Redirect(w, r, "/", http.StatusFound)
		return nil
	})
}

func (h *RootHandler) logout(w http.ResponseWriter, r *http.Request) {
	session := h.Session(r)
	delete(session.Values, "UserID")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *RootHandler) signup(w http.ResponseWriter, r *http.Request) {
	h.db.With(func(tx *db.Tx) error {
		user, _ := h.Authenticate(tx, r)
		if user != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}

		template.New(h.Session(r), nil, nil).Signup(w)
		return nil
	})
}

func (h *RootHandler) doSignup(w http.ResponseWriter, r *http.Request) {
	session := h.Session(r)
	h.db.Do(func(tx *db.Tx) error {
		// Create a new account.
		account := &db.Account{}
		if err := tx.CreateAccount(account); err != nil {
			session.AddFlash(err.Error())
			session.Save(r, w)
			http.Redirect(w, r, r.URL.Path, http.StatusInternalServerError)
			return err
		}

		// Create a new account.
		user := &db.User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		// Create user.
		if err := account.CreateUser(user); err != nil {
			session.AddFlash(err.Error())
			session.Save(r, w)
			http.Redirect(w, r, r.URL.Path, http.StatusInternalServerError)
			return err
		}

		// Update session.
		session.Values["UserID"] = user.ID()
		session.Save(r, w)

		// Redirect to home page.
		http.Redirect(w, r, "/", http.StatusFound)
		return nil
	})
}
