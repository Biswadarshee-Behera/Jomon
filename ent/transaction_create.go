// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/groupbudget"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/tag"
	"github.com/traPtitech/Jomon/ent/transaction"
	"github.com/traPtitech/Jomon/ent/transactiondetail"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TransactionCreate) SetCreatedAt(t time.Time) *TransactionCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCreatedAt(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetDetailID sets the "detail" edge to the TransactionDetail entity by ID.
func (tc *TransactionCreate) SetDetailID(id uuid.UUID) *TransactionCreate {
	tc.mutation.SetDetailID(id)
	return tc
}

// SetNillableDetailID sets the "detail" edge to the TransactionDetail entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableDetailID(id *uuid.UUID) *TransactionCreate {
	if id != nil {
		tc = tc.SetDetailID(*id)
	}
	return tc
}

// SetDetail sets the "detail" edge to the TransactionDetail entity.
func (tc *TransactionCreate) SetDetail(t *TransactionDetail) *TransactionCreate {
	return tc.SetDetailID(t.ID)
}

// AddTagIDs adds the "tag" edge to the Tag entity by IDs.
func (tc *TransactionCreate) AddTagIDs(ids ...uuid.UUID) *TransactionCreate {
	tc.mutation.AddTagIDs(ids...)
	return tc
}

// AddTag adds the "tag" edges to the Tag entity.
func (tc *TransactionCreate) AddTag(t ...*Tag) *TransactionCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTagIDs(ids...)
}

// SetGroupBudgetID sets the "group_budget" edge to the GroupBudget entity by ID.
func (tc *TransactionCreate) SetGroupBudgetID(id uuid.UUID) *TransactionCreate {
	tc.mutation.SetGroupBudgetID(id)
	return tc
}

// SetNillableGroupBudgetID sets the "group_budget" edge to the GroupBudget entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableGroupBudgetID(id *uuid.UUID) *TransactionCreate {
	if id != nil {
		tc = tc.SetGroupBudgetID(*id)
	}
	return tc
}

// SetGroupBudget sets the "group_budget" edge to the GroupBudget entity.
func (tc *TransactionCreate) SetGroupBudget(g *GroupBudget) *TransactionCreate {
	return tc.SetGroupBudgetID(g.ID)
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (tc *TransactionCreate) SetRequestID(id uuid.UUID) *TransactionCreate {
	tc.mutation.SetRequestID(id)
	return tc
}

// SetNillableRequestID sets the "request" edge to the Request entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableRequestID(id *uuid.UUID) *TransactionCreate {
	if id != nil {
		tc = tc.SetRequestID(*id)
	}
	return tc
}

// SetRequest sets the "request" edge to the Request entity.
func (tc *TransactionCreate) SetRequest(r *Request) *TransactionCreate {
	return tc.SetRequestID(r.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransactionCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := transaction.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := transaction.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transaction.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := tc.mutation.DetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transaction.DetailTable,
			Columns: []string{transaction.DetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transactiondetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transaction.TagTable,
			Columns: []string{transaction.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.GroupBudgetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   transaction.GroupBudgetTable,
			Columns: []string{transaction.GroupBudgetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: groupbudget.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_budget_transaction = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.RequestTable,
			Columns: []string{transaction.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.request_transaction = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
