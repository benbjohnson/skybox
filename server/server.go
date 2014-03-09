package server

import (
	"net"
	"net/http"
	"path"

	"github.com/benbjohnson/skybox/db"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Server represents an HTTP interface to the database.
type Server struct {
	http.Server
	*mux.Router
	DB       *db.DB
	listener net.Listener
	store    sessions.Store
}

// ListenAndServe opens the server's port and begins listening for requests.
func (s *Server) ListenAndServe() error {
	// Setup cookie store.
	secret, err := s.DB.Secret()
	if err != nil {
		return err
	}
	s.store = sessions.NewCookieStore(secret)

	// Setup routes.
	s.Router = mux.NewRouter()
	s.Handler = s.Router
	s.HandleFunc("/assets/{filename}", s.assetHandler).Methods("GET")
	(&homeHandler{handler{server: s}}).install()
	(&trackHandler{handler{server: s}}).install()
	(&accountHandler{handler{server: s}}).install()
	(&funnelsHandler{handler{server: s}}).install()

	// Start listening on the socket.
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.listener = listener
	return s.Server.Serve(s.listener)
}

// Close closes the listening port and shutsdown the server.
func (s *Server) Close() {
	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}
}

// assetHandler retrieves static files in the "assets" folder.
func (s *Server) assetHandler(w http.ResponseWriter, r *http.Request) {
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
