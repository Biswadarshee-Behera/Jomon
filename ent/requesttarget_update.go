// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/predicate"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requesttarget"
)

// RequestTargetUpdate is the builder for updating RequestTarget entities.
type RequestTargetUpdate struct {
	config
	hooks    []Hook
	mutation *RequestTargetMutation
}

// Where appends a list predicates to the RequestTargetUpdate builder.
func (rtu *RequestTargetUpdate) Where(ps ...predicate.RequestTarget) *RequestTargetUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetTarget sets the "target" field.
func (rtu *RequestTargetUpdate) SetTarget(s string) *RequestTargetUpdate {
	rtu.mutation.SetTarget(s)
	return rtu
}

// SetAmount sets the "amount" field.
func (rtu *RequestTargetUpdate) SetAmount(i int) *RequestTargetUpdate {
	rtu.mutation.ResetAmount()
	rtu.mutation.SetAmount(i)
	return rtu
}

// AddAmount adds i to the "amount" field.
func (rtu *RequestTargetUpdate) AddAmount(i int) *RequestTargetUpdate {
	rtu.mutation.AddAmount(i)
	return rtu
}

// SetPaidAt sets the "paid_at" field.
func (rtu *RequestTargetUpdate) SetPaidAt(t time.Time) *RequestTargetUpdate {
	rtu.mutation.SetPaidAt(t)
	return rtu
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (rtu *RequestTargetUpdate) SetNillablePaidAt(t *time.Time) *RequestTargetUpdate {
	if t != nil {
		rtu.SetPaidAt(*t)
	}
	return rtu
}

// ClearPaidAt clears the value of the "paid_at" field.
func (rtu *RequestTargetUpdate) ClearPaidAt() *RequestTargetUpdate {
	rtu.mutation.ClearPaidAt()
	return rtu
}

// SetCreatedAt sets the "created_at" field.
func (rtu *RequestTargetUpdate) SetCreatedAt(t time.Time) *RequestTargetUpdate {
	rtu.mutation.SetCreatedAt(t)
	return rtu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtu *RequestTargetUpdate) SetNillableCreatedAt(t *time.Time) *RequestTargetUpdate {
	if t != nil {
		rtu.SetCreatedAt(*t)
	}
	return rtu
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (rtu *RequestTargetUpdate) SetRequestID(id uuid.UUID) *RequestTargetUpdate {
	rtu.mutation.SetRequestID(id)
	return rtu
}

// SetRequest sets the "request" edge to the Request entity.
func (rtu *RequestTargetUpdate) SetRequest(r *Request) *RequestTargetUpdate {
	return rtu.SetRequestID(r.ID)
}

// Mutation returns the RequestTargetMutation object of the builder.
func (rtu *RequestTargetUpdate) Mutation() *RequestTargetMutation {
	return rtu.mutation
}

// ClearRequest clears the "request" edge to the Request entity.
func (rtu *RequestTargetUpdate) ClearRequest() *RequestTargetUpdate {
	rtu.mutation.ClearRequest()
	return rtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *RequestTargetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rtu.hooks) == 0 {
		if err = rtu.check(); err != nil {
			return 0, err
		}
		affected, err = rtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestTargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtu.check(); err != nil {
				return 0, err
			}
			rtu.mutation = mutation
			affected, err = rtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtu.hooks) - 1; i >= 0; i-- {
			if rtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RequestTargetUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RequestTargetUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RequestTargetUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtu *RequestTargetUpdate) check() error {
	if _, ok := rtu.mutation.RequestID(); rtu.mutation.RequestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RequestTarget.request"`)
	}
	return nil
}

func (rtu *RequestTargetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requesttarget.Table,
			Columns: requesttarget.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: requesttarget.FieldID,
			},
		},
	}
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: requesttarget.FieldTarget,
		})
	}
	if value, ok := rtu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: requesttarget.FieldAmount,
		})
	}
	if value, ok := rtu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: requesttarget.FieldAmount,
		})
	}
	if value, ok := rtu.mutation.PaidAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requesttarget.FieldPaidAt,
		})
	}
	if rtu.mutation.PaidAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: requesttarget.FieldPaidAt,
		})
	}
	if value, ok := rtu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requesttarget.FieldCreatedAt,
		})
	}
	if rtu.mutation.RequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requesttarget.RequestTable,
			Columns: []string{requesttarget.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requesttarget.RequestTable,
			Columns: []string{requesttarget.RequestColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{requesttarget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RequestTargetUpdateOne is the builder for updating a single RequestTarget entity.
type RequestTargetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RequestTargetMutation
}

// SetTarget sets the "target" field.
func (rtuo *RequestTargetUpdateOne) SetTarget(s string) *RequestTargetUpdateOne {
	rtuo.mutation.SetTarget(s)
	return rtuo
}

// SetAmount sets the "amount" field.
func (rtuo *RequestTargetUpdateOne) SetAmount(i int) *RequestTargetUpdateOne {
	rtuo.mutation.ResetAmount()
	rtuo.mutation.SetAmount(i)
	return rtuo
}

// AddAmount adds i to the "amount" field.
func (rtuo *RequestTargetUpdateOne) AddAmount(i int) *RequestTargetUpdateOne {
	rtuo.mutation.AddAmount(i)
	return rtuo
}

// SetPaidAt sets the "paid_at" field.
func (rtuo *RequestTargetUpdateOne) SetPaidAt(t time.Time) *RequestTargetUpdateOne {
	rtuo.mutation.SetPaidAt(t)
	return rtuo
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (rtuo *RequestTargetUpdateOne) SetNillablePaidAt(t *time.Time) *RequestTargetUpdateOne {
	if t != nil {
		rtuo.SetPaidAt(*t)
	}
	return rtuo
}

// ClearPaidAt clears the value of the "paid_at" field.
func (rtuo *RequestTargetUpdateOne) ClearPaidAt() *RequestTargetUpdateOne {
	rtuo.mutation.ClearPaidAt()
	return rtuo
}

// SetCreatedAt sets the "created_at" field.
func (rtuo *RequestTargetUpdateOne) SetCreatedAt(t time.Time) *RequestTargetUpdateOne {
	rtuo.mutation.SetCreatedAt(t)
	return rtuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtuo *RequestTargetUpdateOne) SetNillableCreatedAt(t *time.Time) *RequestTargetUpdateOne {
	if t != nil {
		rtuo.SetCreatedAt(*t)
	}
	return rtuo
}

// SetRequestID sets the "request" edge to the Request entity by ID.
func (rtuo *RequestTargetUpdateOne) SetRequestID(id uuid.UUID) *RequestTargetUpdateOne {
	rtuo.mutation.SetRequestID(id)
	return rtuo
}

// SetRequest sets the "request" edge to the Request entity.
func (rtuo *RequestTargetUpdateOne) SetRequest(r *Request) *RequestTargetUpdateOne {
	return rtuo.SetRequestID(r.ID)
}

// Mutation returns the RequestTargetMutation object of the builder.
func (rtuo *RequestTargetUpdateOne) Mutation() *RequestTargetMutation {
	return rtuo.mutation
}

// ClearRequest clears the "request" edge to the Request entity.
func (rtuo *RequestTargetUpdateOne) ClearRequest() *RequestTargetUpdateOne {
	rtuo.mutation.ClearRequest()
	return rtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *RequestTargetUpdateOne) Select(field string, fields ...string) *RequestTargetUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated RequestTarget entity.
func (rtuo *RequestTargetUpdateOne) Save(ctx context.Context) (*RequestTarget, error) {
	var (
		err  error
		node *RequestTarget
	)
	if len(rtuo.hooks) == 0 {
		if err = rtuo.check(); err != nil {
			return nil, err
		}
		node, err = rtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestTargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtuo.check(); err != nil {
				return nil, err
			}
			rtuo.mutation = mutation
			node, err = rtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtuo.hooks) - 1; i >= 0; i-- {
			if rtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rtuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RequestTarget)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RequestTargetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RequestTargetUpdateOne) SaveX(ctx context.Context) *RequestTarget {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *RequestTargetUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RequestTargetUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtuo *RequestTargetUpdateOne) check() error {
	if _, ok := rtuo.mutation.RequestID(); rtuo.mutation.RequestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RequestTarget.request"`)
	}
	return nil
}

func (rtuo *RequestTargetUpdateOne) sqlSave(ctx context.Context) (_node *RequestTarget, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requesttarget.Table,
			Columns: requesttarget.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: requesttarget.FieldID,
			},
		},
	}
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RequestTarget.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, requesttarget.FieldID)
		for _, f := range fields {
			if !requesttarget.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != requesttarget.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: requesttarget.FieldTarget,
		})
	}
	if value, ok := rtuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: requesttarget.FieldAmount,
		})
	}
	if value, ok := rtuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: requesttarget.FieldAmount,
		})
	}
	if value, ok := rtuo.mutation.PaidAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requesttarget.FieldPaidAt,
		})
	}
	if rtuo.mutation.PaidAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: requesttarget.FieldPaidAt,
		})
	}
	if value, ok := rtuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: requesttarget.FieldCreatedAt,
		})
	}
	if rtuo.mutation.RequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requesttarget.RequestTable,
			Columns: []string{requesttarget.RequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: request.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   requesttarget.RequestTable,
			Columns: []string{requesttarget.RequestColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RequestTarget{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{requesttarget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
