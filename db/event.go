package db

import (
	"fmt"
	"time"

	"github.com/skydb/gosky"
)

const AnonymousPrefix = "@"

// Event represents an action or state change trigger by a tracked user.
type Event struct {
	Timestamp time.Time              `json:"timestamp"`
	UserID    string                 `json:"userID"`
	DeviceID  string                 `json:"deviceID"`
	Channel   string                 `json:"channel"`
	Resource  string                 `json:"resource"`
	Action    string                 `json:"action"`
	Data      map[string]interface{} `json:"data"`
}

// ID returns the identifier used to store the event in Sky.
func (e *Event) ID() string {
	if e.UserID != "" {
		return e.UserID
	}
	return e.AnonymousID()
}

// AnonymousID returns the identifier used when the ID is not available.
func (e *Event) AnonymousID() string {
	if e.DeviceID == "" {
		return ""
	}
	return fmt.Sprintf("%s%s", AnonymousPrefix, e.DeviceID)
}

// IsAnonymous returns whether the event is for an unidentified user.
func (e *Event) IsAnonymous() bool {
	return (e.UserID == "")
}

// Serialize converts the Skybox event into a Sky event.
func (e *Event) Serialize() *sky.Event {
	event := &sky.Event{
		Timestamp: e.Timestamp,
		Data: map[string]interface{}{
			"channel":  e.Channel,
			"resource": e.Resource,
			"action":   e.Action,
		},
	}
	for k, v := range e.Data {
		event.Data[k] = v
	}
	return event
}

// Deserialize converts a Sky event into a Skybox event.
func (e *Event) Deserialize(id string, event *sky.Event) {
	if event == nil {
		event = &sky.Event{}
	}

	// Parse the user or device id from the Sky identifier.
	if id != "" {
		if id[0] == AnonymousPrefix[0] {
			e.UserID = ""
			e.DeviceID = id[1:]
		} else {
			e.UserID = id
			e.DeviceID = ""
		}
	}

	// Parse the remaining properties.
	e.Timestamp = event.Timestamp
	e.Data = make(map[string]interface{})
	for k, v := range event.Data {
		switch k {
		case "channel":
			e.Channel, _ = v.(string)
		case "resource":
			e.Resource, _ = v.(string)
		case "action":
			e.Action, _ = v.(string)
		default:
			e.Data[k] = v
		}
	}
}
