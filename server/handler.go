package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/skybox/skybox/db"
	"github.com/skybox/skybox/server/template"
)

var (
	// ErrUnauthorized is returned when a user is not permitted to perform
	// an action.
	ErrUnauthorized = errors.New("unauthorized")
)

// Handler represents an HTTP interface to the database.
type Handler struct {
	*mux.Router
	db    *db.DB
	store sessions.Store
}

// NewHandler creates a new Handler instance.
func NewHandler(db *db.DB) (*Handler, error) {
	secret, err := db.Secret()
	if err != nil {
		return nil, err
	}

	// Setup the handler.
	h := &Handler{
		Router: mux.NewRouter(),
		db:     db,
		store:  sessions.NewCookieStore(secret),
	}
	h.Handle("/track.png", http.HandlerFunc(h.track)).Methods("GET")
	NewRootHandler(h)
	NewAccountHandler(h)
	NewFunnelHandler(h)
	return h, nil
}

// DB returns the database associated with the handler.
func (h *Handler) DB() *db.DB {
	return h.db
}

func (h *Handler) track(w http.ResponseWriter, r *http.Request) {
	h.db.Do(func(tx *db.Tx) error {
		// Find account by API key.
		a, err := tx.AccountByAPIKey(r.FormValue("apiKey"))
		if err != nil {
			http.Error(w, "invalid api key", http.StatusBadRequest)
			return err
		}

		// Extract event from URL parameters.
		e := &db.Event{
			UserID:    r.FormValue("user.id"),
			DeviceID:  r.FormValue("device.id"),
			Timestamp: time.Now().UTC(),
			Channel:   r.FormValue("channel"),
			Resource:  r.FormValue("resource"),
			Action:    r.FormValue("action"),
			Data:      make(map[string]interface{}),
		}
		if domain := r.FormValue("domain"); len(domain) > 0 {
			e.Data["domain"] = domain
		}
		if path := r.FormValue("path"); len(path) > 0 {
			e.Data["path"] = path
		}

		// Send event to Sky.
		if err := a.Track(e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		// Write png to response.
		b, _ := Asset("pixel.png")
		w.Write(b)

		return nil
	})
}

// auth returns the logged in user and account for a given request.
func (h *Handler) Authenticate(tx *db.Tx, r *http.Request) (*db.User, *db.Account) {
	session := h.Session(r)
	id, ok := session.Values["UserID"]
	if !ok {
		return nil, nil
	}
	if id, ok := id.(int); ok {
		u, err := tx.User(id)
		if err != nil {
			log.Println("[warn] session user not found: %v", err)
		}
		a, _ := u.Account()
		return u, a
	}
	return nil, nil
}

// Session returns the current session for a request.
func (h *Handler) Session(r *http.Request) *sessions.Session {
	session, _ := h.store.Get(r, "default")
	return session
}

// Flashes retrieves all flashes from the session and clears them.
func (h *Handler) Flashes(w http.ResponseWriter, r *http.Request) []string {
	session := h.Session(r)
	flashes := make([]string, 0)
	for _, flash := range session.Flashes() {
		if flash, ok := flash.(string); ok {
			flashes = append(flashes, flash)
		}
	}
	session.Save(r, w)
	return flashes
}

// Unauthorized redirects to the home page.
func (h *Handler) Unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

// NotFound returns a 404 not found page.
func (h *Handler) NotFound(tx *db.Tx, w http.ResponseWriter, r *http.Request) {
	user, account := h.Authenticate(tx, r)
	template.New([]string{"page not found"}, user, account).NotFound(w)
}
