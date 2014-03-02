package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/benbjohnson/skybox/db"
	"github.com/benbjohnson/skybox/server/template"
	"github.com/gorilla/sessions"
)

type handler struct {
	sync.RWMutex
	server       *Server
	transactions map[*http.Request]*db.Transaction
}

// transaction retrieves a transaction for a given request.
func (h *handler) transaction(r *http.Request) *db.Transaction {
	h.RLock()
	defer h.RUnlock()
	return h.transactions[r]
}

// setTransaction sets a transaction for a given request.
func (h *handler) setTransaction(r *http.Request, t *db.Transaction) {
	h.Lock()
	defer h.Unlock()
	if h.transactions == nil {
		h.transactions = make(map[*http.Request]*db.Transaction)
	}
	h.transactions[r] = t
}

// removeTransaction removes a transaction for a request.
func (h *handler) removeTransaction(r *http.Request) {
	h.Lock()
	defer h.Unlock()
	delete(h.transactions, r)
}

// transactional executes a handler in the context of a read/write transaction.
func (h *handler) transact(handler http.Handler) http.Handler {
	return &transactor{parent: h, handler: handler}
}

// rwtransactional executes a handler in the context of a read/write transaction.
func (h *handler) rwtransact(handler http.Handler) http.Handler {
	return &rwtransactor{parent: h, handler: handler}
}

func (h *handler) authorize(handler http.Handler) http.Handler {
	return &authorizer{parent: h, handler: handler}
}

// session returns the current session.
func (h *handler) session(r *http.Request) *sessions.Session {
	session, _ := h.server.store.Get(r, "default")
	return session
}

// auth returns the logged in user and account for a given request.
func (h *handler) auth(r *http.Request) (*db.User, *db.Account) {
	txn := h.transaction(r)
	session := h.session(r)
	id, ok := session.Values["UserID"]
	if !ok {
		return nil, nil
	}
	if id, ok := id.(int); ok {
		u, err := txn.User(id)
		if err != nil {
			log.Println("[warn] session user not found: %v", err)
		}
		a, _ := u.Account()
		return u, a
	}
	return nil, nil
}

// notFound returns a 404 not found page.
func (h *handler) notFound(w http.ResponseWriter, r *http.Request) {
	user, account := h.auth(r)
	t := template.New(h.session(r), user, account)
	t.NotFound(w)
}

// transactor executes a handler in the context of a read-only transaction.
type transactor struct {
	parent  *handler
	handler http.Handler
}

func (t *transactor) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.parent.server.DB.With(func(txn *db.Transaction) error {
		t.parent.setTransaction(req, txn)
		t.handler.ServeHTTP(w, req)
		t.parent.removeTransaction(req)
		return nil
	})
}

// rwtransactor executes a handler in the context of a read/write transaction.
type rwtransactor struct {
	parent  *handler
	handler http.Handler
}

func (t *rwtransactor) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := t.parent.server.DB.Do(func(txn *db.Transaction) error {
		t.parent.setTransaction(req, txn)
		t.handler.ServeHTTP(w, req)
		t.parent.removeTransaction(req)
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// authorizer checks that there is a user id in the session before allowing
// the handler to continue.
type authorizer struct {
	parent  *handler
	handler http.Handler
}

func (a *authorizer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	session, _ := a.parent.server.store.Get(req, "default")
	if _, ok := session.Values["UserID"]; !ok {
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}
	a.handler.ServeHTTP(w, req)
}
