// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupbudget"
	"github.com/traPtitech/Jomon/ent/groupowner"
	"github.com/traPtitech/Jomon/ent/groupuser"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetDescription sets the "description" field.
func (gc *GroupCreate) SetDescription(s string) *GroupCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetBudget sets the "budget" field.
func (gc *GroupCreate) SetBudget(i int) *GroupCreate {
	gc.mutation.SetBudget(i)
	return gc
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (gc *GroupCreate) SetNillableBudget(i *int) *GroupCreate {
	if i != nil {
		gc.SetBudget(*i)
	}
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GroupCreate) SetUpdatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableUpdatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GroupCreate) SetDeletedAt(t time.Time) *GroupCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableDeletedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// AddGroupBudgetIDs adds the "group_budget" edge to the GroupBudget entity by IDs.
func (gc *GroupCreate) AddGroupBudgetIDs(ids ...int) *GroupCreate {
	gc.mutation.AddGroupBudgetIDs(ids...)
	return gc
}

// AddGroupBudget adds the "group_budget" edges to the GroupBudget entity.
func (gc *GroupCreate) AddGroupBudget(g ...*GroupBudget) *GroupCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddGroupBudgetIDs(ids...)
}

// AddUserIDs adds the "user" edge to the GroupUser entity by IDs.
func (gc *GroupCreate) AddUserIDs(ids ...int) *GroupCreate {
	gc.mutation.AddUserIDs(ids...)
	return gc
}

// AddUser adds the "user" edges to the GroupUser entity.
func (gc *GroupCreate) AddUser(g ...*GroupUser) *GroupCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the GroupOwner entity by IDs.
func (gc *GroupCreate) AddOwnerIDs(ids ...int) *GroupCreate {
	gc.mutation.AddOwnerIDs(ids...)
	return gc
}

// AddOwner adds the "owner" edges to the GroupOwner entity.
func (gc *GroupCreate) AddOwner(g ...*GroupOwner) *GroupCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddOwnerIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := group.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := group.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := gc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: group.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := gc.mutation.Budget(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldBudget,
		})
		_node.Budget = &value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := gc.mutation.GroupBudgetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupBudgetTable,
			Columns: []string{group.GroupBudgetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupbudget.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.UserTable,
			Columns: []string{group.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupowner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	builders []*GroupCreate
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
