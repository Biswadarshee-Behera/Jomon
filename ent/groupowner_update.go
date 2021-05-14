// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupowner"
	"github.com/traPtitech/Jomon/ent/predicate"
)

// GroupOwnerUpdate is the builder for updating GroupOwner entities.
type GroupOwnerUpdate struct {
	config
	hooks    []Hook
	mutation *GroupOwnerMutation
}

// Where adds a new predicate for the GroupOwnerUpdate builder.
func (gou *GroupOwnerUpdate) Where(ps ...predicate.GroupOwner) *GroupOwnerUpdate {
	gou.mutation.predicates = append(gou.mutation.predicates, ps...)
	return gou
}

// SetOwner sets the "owner" field.
func (gou *GroupOwnerUpdate) SetOwner(s string) *GroupOwnerUpdate {
	gou.mutation.SetOwner(s)
	return gou
}

// SetGroupID sets the "group_id" field.
func (gou *GroupOwnerUpdate) SetGroupID(i int) *GroupOwnerUpdate {
	gou.mutation.ResetGroupID()
	gou.mutation.SetGroupID(i)
	return gou
}

// SetCreatedAt sets the "created_at" field.
func (gou *GroupOwnerUpdate) SetCreatedAt(t time.Time) *GroupOwnerUpdate {
	gou.mutation.SetCreatedAt(t)
	return gou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gou *GroupOwnerUpdate) SetNillableCreatedAt(t *time.Time) *GroupOwnerUpdate {
	if t != nil {
		gou.SetCreatedAt(*t)
	}
	return gou
}

// SetGroup sets the "group" edge to the Group entity.
func (gou *GroupOwnerUpdate) SetGroup(g *Group) *GroupOwnerUpdate {
	return gou.SetGroupID(g.ID)
}

// Mutation returns the GroupOwnerMutation object of the builder.
func (gou *GroupOwnerUpdate) Mutation() *GroupOwnerMutation {
	return gou.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gou *GroupOwnerUpdate) ClearGroup() *GroupOwnerUpdate {
	gou.mutation.ClearGroup()
	return gou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gou *GroupOwnerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gou.hooks) == 0 {
		if err = gou.check(); err != nil {
			return 0, err
		}
		affected, err = gou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupOwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gou.check(); err != nil {
				return 0, err
			}
			gou.mutation = mutation
			affected, err = gou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gou.hooks) - 1; i >= 0; i-- {
			mut = gou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gou *GroupOwnerUpdate) SaveX(ctx context.Context) int {
	affected, err := gou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gou *GroupOwnerUpdate) Exec(ctx context.Context) error {
	_, err := gou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gou *GroupOwnerUpdate) ExecX(ctx context.Context) {
	if err := gou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gou *GroupOwnerUpdate) check() error {
	if _, ok := gou.mutation.GroupID(); gou.mutation.GroupCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"group\"")
	}
	return nil
}

func (gou *GroupOwnerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupowner.Table,
			Columns: groupowner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupowner.FieldID,
			},
		},
	}
	if ps := gou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gou.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupowner.FieldOwner,
		})
	}
	if value, ok := gou.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupowner.FieldCreatedAt,
		})
	}
	if gou.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupowner.GroupTable,
			Columns: []string{groupowner.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gou.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupowner.GroupTable,
			Columns: []string{groupowner.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupowner.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupOwnerUpdateOne is the builder for updating a single GroupOwner entity.
type GroupOwnerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupOwnerMutation
}

// SetOwner sets the "owner" field.
func (gouo *GroupOwnerUpdateOne) SetOwner(s string) *GroupOwnerUpdateOne {
	gouo.mutation.SetOwner(s)
	return gouo
}

// SetGroupID sets the "group_id" field.
func (gouo *GroupOwnerUpdateOne) SetGroupID(i int) *GroupOwnerUpdateOne {
	gouo.mutation.ResetGroupID()
	gouo.mutation.SetGroupID(i)
	return gouo
}

// SetCreatedAt sets the "created_at" field.
func (gouo *GroupOwnerUpdateOne) SetCreatedAt(t time.Time) *GroupOwnerUpdateOne {
	gouo.mutation.SetCreatedAt(t)
	return gouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gouo *GroupOwnerUpdateOne) SetNillableCreatedAt(t *time.Time) *GroupOwnerUpdateOne {
	if t != nil {
		gouo.SetCreatedAt(*t)
	}
	return gouo
}

// SetGroup sets the "group" edge to the Group entity.
func (gouo *GroupOwnerUpdateOne) SetGroup(g *Group) *GroupOwnerUpdateOne {
	return gouo.SetGroupID(g.ID)
}

// Mutation returns the GroupOwnerMutation object of the builder.
func (gouo *GroupOwnerUpdateOne) Mutation() *GroupOwnerMutation {
	return gouo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gouo *GroupOwnerUpdateOne) ClearGroup() *GroupOwnerUpdateOne {
	gouo.mutation.ClearGroup()
	return gouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gouo *GroupOwnerUpdateOne) Select(field string, fields ...string) *GroupOwnerUpdateOne {
	gouo.fields = append([]string{field}, fields...)
	return gouo
}

// Save executes the query and returns the updated GroupOwner entity.
func (gouo *GroupOwnerUpdateOne) Save(ctx context.Context) (*GroupOwner, error) {
	var (
		err  error
		node *GroupOwner
	)
	if len(gouo.hooks) == 0 {
		if err = gouo.check(); err != nil {
			return nil, err
		}
		node, err = gouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupOwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gouo.check(); err != nil {
				return nil, err
			}
			gouo.mutation = mutation
			node, err = gouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gouo.hooks) - 1; i >= 0; i-- {
			mut = gouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gouo *GroupOwnerUpdateOne) SaveX(ctx context.Context) *GroupOwner {
	node, err := gouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gouo *GroupOwnerUpdateOne) Exec(ctx context.Context) error {
	_, err := gouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gouo *GroupOwnerUpdateOne) ExecX(ctx context.Context) {
	if err := gouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gouo *GroupOwnerUpdateOne) check() error {
	if _, ok := gouo.mutation.GroupID(); gouo.mutation.GroupCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"group\"")
	}
	return nil
}

func (gouo *GroupOwnerUpdateOne) sqlSave(ctx context.Context) (_node *GroupOwner, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupowner.Table,
			Columns: groupowner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupowner.FieldID,
			},
		},
	}
	id, ok := gouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GroupOwner.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := gouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupowner.FieldID)
		for _, f := range fields {
			if !groupowner.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupowner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gouo.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupowner.FieldOwner,
		})
	}
	if value, ok := gouo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupowner.FieldCreatedAt,
		})
	}
	if gouo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupowner.GroupTable,
			Columns: []string{groupowner.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gouo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupowner.GroupTable,
			Columns: []string{groupowner.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupOwner{config: gouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupowner.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
