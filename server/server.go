package server

import (
	"net"
	"net/http"
	"path"

	"github.com/benbjohnson/skybox/db"
	"github.com/benbjohnson/skybox/server/templates"
	"github.com/gorilla/mux"
)

// Server represents an HTTP interface to the database.
type Server struct {
	http.Server
	DB       *db.DB
	listener net.Listener
}

// ListenAndServe opens the server's port and begins listening for requests.
func (s *Server) ListenAndServe() error {
	router := mux.NewRouter()
	router.HandleFunc("/assets/{filename}", s.assetHandler).Methods("GET")
	router.HandleFunc("/", s.indexHandler).Methods("GET")
	s.Handler = router

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

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.Index(w)
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
