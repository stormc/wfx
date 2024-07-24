// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/siemens/wfx/generated/api"
	"github.com/siemens/wfx/generated/ent/job"
	"github.com/siemens/wfx/generated/ent/workflow"
)

// WorkflowCreate is the builder for creating a Workflow entity.
type WorkflowCreate struct {
	config
	mutation *WorkflowMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wc *WorkflowCreate) SetName(s string) *WorkflowCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkflowCreate) SetDescription(s string) *WorkflowCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableDescription(s *string) *WorkflowCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetStates sets the "states" field.
func (wc *WorkflowCreate) SetStates(a []api.State) *WorkflowCreate {
	wc.mutation.SetStates(a)
	return wc
}

// SetTransitions sets the "transitions" field.
func (wc *WorkflowCreate) SetTransitions(a []api.Transition) *WorkflowCreate {
	wc.mutation.SetTransitions(a)
	return wc
}

// SetGroups sets the "groups" field.
func (wc *WorkflowCreate) SetGroups(a []api.Group) *WorkflowCreate {
	wc.mutation.SetGroups(a)
	return wc
}

// AddJobIDs adds the "jobs" edge to the Job entity by IDs.
func (wc *WorkflowCreate) AddJobIDs(ids ...string) *WorkflowCreate {
	wc.mutation.AddJobIDs(ids...)
	return wc
}

// AddJobs adds the "jobs" edges to the Job entity.
func (wc *WorkflowCreate) AddJobs(j ...*Job) *WorkflowCreate {
	ids := make([]string, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return wc.AddJobIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wc *WorkflowCreate) Mutation() *WorkflowMutation {
	return wc.mutation
}

// Save creates the Workflow in the database.
func (wc *WorkflowCreate) Save(ctx context.Context) (*Workflow, error) {
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkflowCreate) SaveX(ctx context.Context) *Workflow {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkflowCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkflowCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkflowCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Workflow.name"`)}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := workflow.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Workflow.name": %w`, err)}
		}
	}
	if v, ok := wc.mutation.Description(); ok {
		if err := workflow.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Workflow.description": %w`, err)}
		}
	}
	if _, ok := wc.mutation.States(); !ok {
		return &ValidationError{Name: "states", err: errors.New(`ent: missing required field "Workflow.states"`)}
	}
	if _, ok := wc.mutation.Transitions(); !ok {
		return &ValidationError{Name: "transitions", err: errors.New(`ent: missing required field "Workflow.transitions"`)}
	}
	if _, ok := wc.mutation.Groups(); !ok {
		return &ValidationError{Name: "groups", err: errors.New(`ent: missing required field "Workflow.groups"`)}
	}
	return nil
}

func (wc *WorkflowCreate) sqlSave(ctx context.Context) (*Workflow, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkflowCreate) createSpec() (*Workflow, *sqlgraph.CreateSpec) {
	var (
		_node = &Workflow{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(workflow.Table, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(workflow.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(workflow.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.States(); ok {
		_spec.SetField(workflow.FieldStates, field.TypeJSON, value)
		_node.States = value
	}
	if value, ok := wc.mutation.Transitions(); ok {
		_spec.SetField(workflow.FieldTransitions, field.TypeJSON, value)
		_node.Transitions = value
	}
	if value, ok := wc.mutation.Groups(); ok {
		_spec.SetField(workflow.FieldGroups, field.TypeJSON, value)
		_node.Groups = value
	}
	if nodes := wc.mutation.JobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.JobsTable,
			Columns: []string{workflow.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkflowCreateBulk is the builder for creating many Workflow entities in bulk.
type WorkflowCreateBulk struct {
	config
	err      error
	builders []*WorkflowCreate
}

// Save creates the Workflow entities in the database.
func (wcb *WorkflowCreateBulk) Save(ctx context.Context) ([]*Workflow, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workflow, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkflowMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) SaveX(ctx context.Context) []*Workflow {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkflowCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
