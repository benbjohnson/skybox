package db

var (
	// ErrFunnelNotFound is returned when a funnel does not exist.
	ErrFunnelNotFound = &Error{"funnel not found", nil}

	// ErrFunnelNameRequired is returned when a funnel has a blank name.
	ErrFunnelNameRequired = &Error{"funnel name required", nil}

	// ErrFunnelStepsRequired is returned when a funnel has no steps.
	ErrFunnelStepsRequired = &Error{"funnel steps required", nil}
)

// Funnel represents a multi-step query.
// A Funnel belongs to a Project.
type Funnel struct {
	Transaction *Transaction
	id          int
	ProjectID   int           `json:"projectID"`
	Name        string        `json:"name"`
	Steps       []*FunnelStep `json:"steps"`
}

// FunnelStep represents a single step in a funnel.
type FunnelStep struct {
	Condition string `json:"condition"`
}

// ID returns the funnel identifier.
func (f *Funnel) ID() int {
	return f.id
}

// Validate validates all fields of the funnel.
func (f *Funnel) Validate() error {
	if len(f.Name) == 0 {
		return ErrFunnelNameRequired
	} else if len(f.Steps) == 0 {
		return ErrFunnelStepsRequired
	}
	return nil
}

func (f *Funnel) get() ([]byte, error) {
	value := f.Transaction.Bucket("funnels").Get(itob(f.id))
	if value == nil {
		return nil, ErrFunnelNotFound
	}
	return value, nil
}

// Load retrieves a funnel from the database.
func (f *Funnel) Load() error {
	value, err := f.get()
	if err != nil {
		return err
	}
	unmarshal(value, &f)
	return nil
}

// Save commits the Funnel to the database.
func (f *Funnel) Save() error {
	assert(f.id > 0, "uninitialized funnel cannot be saved")
	return f.Transaction.Bucket("funnels").Put(itob(f.id), marshal(f))
}

// Delete removes the Funnel from the database.
func (f *Funnel) Delete() error {
	// Remove project entry.
	err := f.Transaction.Bucket("funnels").Delete(itob(f.id))
	assert(err == nil, "funnel delete error: %s", err)

	// Remove funnel id from indices.
	removeFromForeignKeyIndex(f.Transaction, "project.funnels", itob(f.ProjectID), f.id)

	return nil
}

type Funnels []*Funnel

func (s Funnels) Len() int           { return len(s) }
func (s Funnels) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Funnels) Less(i, j int) bool { return s[i].Name < s[j].Name }
