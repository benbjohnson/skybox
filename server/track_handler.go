package server

import (
	"net/http"
	"time"

	"github.com/benbjohnson/skybox/db"
)

type trackHandler struct {
	handler
}

func (h *trackHandler) install() {
	h.server.Handle("/track.png", h.transact(http.HandlerFunc(h.track))).Methods("GET")
}

func (h *trackHandler) track(w http.ResponseWriter, r *http.Request) {
	tx := h.transaction(r)

	// Find project by API key.
	p, err := tx.ProjectByAPIKey(r.FormValue("apiKey"))
	if err != nil {
		http.Error(w, "invalid api key", http.StatusBadRequest)
		return
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
	if err := p.Track(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write png to response.
	b, _ := Asset("pixel.png")
	w.Write(b)
}
