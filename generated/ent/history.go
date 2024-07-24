// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/siemens/wfx/generated/api"
	"github.com/siemens/wfx/generated/ent/history"
	"github.com/siemens/wfx/generated/ent/job"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// modification time
	Mtime time.Time `json:"mtime,omitempty"`
	// Status holds the value of the "status" field.
	Status api.JobStatus `json:"status,omitempty"`
	// Definition holds the value of the "definition" field.
	Definition map[string]interface{} `json:"definition,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistoryQuery when eager-loading is set.
	Edges        HistoryEdges `json:"edges"`
	job_history  *string
	selectValues sql.SelectValues
}

// HistoryEdges holds the relations/edges for other nodes in the graph.
type HistoryEdges struct {
	// Job holds the value of the job edge.
	Job *Job `json:"job,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// JobOrErr returns the Job value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) JobOrErr() (*Job, error) {
	if e.Job != nil {
		return e.Job, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: job.Label}
	}
	return nil, &NotLoadedError{edge: "job"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case history.FieldStatus, history.FieldDefinition:
			values[i] = new([]byte)
		case history.FieldID:
			values[i] = new(sql.NullInt64)
		case history.FieldMtime:
			values[i] = new(sql.NullTime)
		case history.ForeignKeys[0]: // job_history
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case history.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				h.Mtime = value.Time
			}
		case history.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case history.FieldDefinition:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field definition", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.Definition); err != nil {
					return fmt.Errorf("unmarshal field definition: %w", err)
				}
			}
		case history.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_history", values[i])
			} else if value.Valid {
				h.job_history = new(string)
				*h.job_history = value.String
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the History.
// This includes values selected through modifiers, order, etc.
func (h *History) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryJob queries the "job" edge of the History entity.
func (h *History) QueryJob() *JobQuery {
	return NewHistoryClient(h.config).QueryJob(h)
}

// Update returns a builder for updating this History.
// Note that you need to call History.Unwrap() before calling this method if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return NewHistoryClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the History entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("mtime=")
	builder.WriteString(h.Mtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", h.Status))
	builder.WriteString(", ")
	builder.WriteString("definition=")
	builder.WriteString(fmt.Sprintf("%v", h.Definition))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History
