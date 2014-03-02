package server

import (
	"net/http"

	"github.com/benbjohnson/skybox/db"
)

type trackHandler struct {
	handler
}

func (h *homeHandler) install() {
	h.server.Handle("/track.png", h.transact(http.HandlerFunc(h.track))).Methods("GET")
}

func (h *homeHandler) track(w http.ResponseWriter, r *http.Request) {
	txn := h.transaction(r)

	// Find project by API key.
	p, err := txn.ProjectByAPIKey(r.FormValue("apiKey"))
	if err != nil {
		http.Error(w, "invalid api key", http.StatusBadRequest)
		return
	}

	// TODO: Extract event from URL parameters.
	// e := &Event{UserID, DeviceID, Channel, Resource, Action, Data:{Domain, Path}}

	// TODO: Send event to Sky.
	// p.Track(e)

	// TODO: Write track.png to writer.
}
