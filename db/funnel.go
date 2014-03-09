package db

import (
	"bytes"
	"fmt"
	"strings"
)

// SessionIdleTime is the amount of idle time that delimits sessions.
const SessionIdleTime = "2 HOURS"

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
	Tx        *Tx
	id        int
	ProjectID int           `json:"projectID"`
	Name      string        `json:"name"`
	Steps     []*FunnelStep `json:"steps"`
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
	value := f.Tx.Bucket("funnels").Get(itob(f.id))
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
	return f.Tx.Bucket("funnels").Put(itob(f.id), marshal(f))
}

// Delete removes the Funnel from the database.
func (f *Funnel) Delete() error {
	// Remove project entry.
	err := f.Tx.Bucket("funnels").Delete(itob(f.id))
	assert(err == nil, "funnel delete error: %s", err)

	// Remove funnel id from indices.
	removeFromForeignKeyIndex(f.Tx, "project.funnels", itob(f.ProjectID), f.id)

	return nil
}

// Query executes a query against the funnel and returns the result.
func (f *Funnel) Query() (*FunnelResult, error) {
	// Generate query.
	querystring := f.QueryString()

	// Retrieve Sky table.
	p := &Project{id: f.ProjectID, Tx: f.Tx}
	t := p.SkyTable()

	// Execute query against Sky.
	raw, err := t.Query(querystring)
	if err != nil {
		return nil, err
	}

	// Deserialize into results.
	result := &FunnelResult{
		Name:  f.Name,
		Steps: make([]*FunnelStepResult, 0),
	}
	for i, step := range f.Steps {
		stepResult := &FunnelStepResult{Condition: step.Condition}
		if rawStep, ok := raw[fmt.Sprintf("step%d", i)].(map[string]interface{}); ok {
			if count, ok := rawStep["count"].(float64); ok {
				stepResult.Count = int(count)
			}
		}
		result.Steps = append(result.Steps, stepResult)
	}

	return result, nil
}

// QueryString generates a funnel query for use in Sky.
func (f *Funnel) QueryString() string {
	var buf bytes.Buffer

	fmt.Fprintln(&buf, "FOR EACH SESSION DELIMITED BY", SessionIdleTime)
	fmt.Fprintln(&buf, "  FOR EACH EVENT")

	for i, step := range f.Steps {
		var within = ""
		if i > 0 {
			within = " WITHIN 1..100000 STEPS"
		}
		// Write step condition.
		fmt.Fprintf(&buf, "%s    WHEN %s%s THEN\n", strings.Repeat("  ", i), step.Condition, within)

		// Write step selection.
		fmt.Fprintf(&buf, "%s      SELECT count() INTO \"step%d\"\n", strings.Repeat("  ", i), i)
	}

	for i, _ := range f.Steps {
		// Write condition block close.
		fmt.Fprintf(&buf, "%s    END\n", strings.Repeat("  ", len(f.Steps)-i-1))
	}

	fmt.Fprintln(&buf, "  END")
	fmt.Fprintln(&buf, "END")

	return buf.String()
}

type Funnels []*Funnel

func (s Funnels) Len() int           { return len(s) }
func (s Funnels) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Funnels) Less(i, j int) bool { return s[i].Name < s[j].Name }

// FunnelResult represents the results of an executed funnel query.
type FunnelResult struct {
	Name  string              `json:"name"`
	Steps []*FunnelStepResult `json:"steps"`
}

// FunnelStepResult represents the result of one step in an executed funnel query.
type FunnelStepResult struct {
	Condition string `json:"condition"`
	Count     int    `json:"count"`
}
