// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/siemens/wfx/generated/api"
	"github.com/siemens/wfx/generated/ent/workflow"
)

// Workflow is the model entity for the Workflow schema.
type Workflow struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// States holds the value of the "states" field.
	States []api.State `json:"states,omitempty"`
	// Transitions holds the value of the "transitions" field.
	Transitions []api.Transition `json:"transitions,omitempty"`
	// Groups holds the value of the "groups" field.
	Groups []api.Group `json:"groups,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowQuery when eager-loading is set.
	Edges        WorkflowEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WorkflowEdges holds the relations/edges for other nodes in the graph.
type WorkflowEdges struct {
	// Jobs holds the value of the jobs edge.
	Jobs []*Job `json:"jobs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// JobsOrErr returns the Jobs value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowEdges) JobsOrErr() ([]*Job, error) {
	if e.loadedTypes[0] {
		return e.Jobs, nil
	}
	return nil, &NotLoadedError{edge: "jobs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workflow) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflow.FieldStates, workflow.FieldTransitions, workflow.FieldGroups:
			values[i] = new([]byte)
		case workflow.FieldID:
			values[i] = new(sql.NullInt64)
		case workflow.FieldName, workflow.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workflow fields.
func (w *Workflow) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflow.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case workflow.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case workflow.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case workflow.FieldStates:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field states", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.States); err != nil {
					return fmt.Errorf("unmarshal field states: %w", err)
				}
			}
		case workflow.FieldTransitions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field transitions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.Transitions); err != nil {
					return fmt.Errorf("unmarshal field transitions: %w", err)
				}
			}
		case workflow.FieldGroups:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field groups", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.Groups); err != nil {
					return fmt.Errorf("unmarshal field groups: %w", err)
				}
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Workflow.
// This includes values selected through modifiers, order, etc.
func (w *Workflow) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryJobs queries the "jobs" edge of the Workflow entity.
func (w *Workflow) QueryJobs() *JobQuery {
	return NewWorkflowClient(w.config).QueryJobs(w)
}

// Update returns a builder for updating this Workflow.
// Note that you need to call Workflow.Unwrap() before calling this method if this Workflow
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workflow) Update() *WorkflowUpdateOne {
	return NewWorkflowClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Workflow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workflow) Unwrap() *Workflow {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Workflow is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workflow) String() string {
	var builder strings.Builder
	builder.WriteString("Workflow(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("name=")
	builder.WriteString(w.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(w.Description)
	builder.WriteString(", ")
	builder.WriteString("states=")
	builder.WriteString(fmt.Sprintf("%v", w.States))
	builder.WriteString(", ")
	builder.WriteString("transitions=")
	builder.WriteString(fmt.Sprintf("%v", w.Transitions))
	builder.WriteString(", ")
	builder.WriteString("groups=")
	builder.WriteString(fmt.Sprintf("%v", w.Groups))
	builder.WriteByte(')')
	return builder.String()
}

// Workflows is a parsable slice of Workflow.
type Workflows []*Workflow
