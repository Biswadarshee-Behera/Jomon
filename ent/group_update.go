// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupbudget"
	"github.com/traPtitech/Jomon/ent/groupowner"
	"github.com/traPtitech/Jomon/ent/groupuser"
	"github.com/traPtitech/Jomon/ent/predicate"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where adds a new predicate for the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.predicates = append(gu.mutation.predicates, ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetDescription sets the "description" field.
func (gu *GroupUpdate) SetDescription(s string) *GroupUpdate {
	gu.mutation.SetDescription(s)
	return gu
}

// SetBudget sets the "budget" field.
func (gu *GroupUpdate) SetBudget(i int) *GroupUpdate {
	gu.mutation.ResetBudget()
	gu.mutation.SetBudget(i)
	return gu
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableBudget(i *int) *GroupUpdate {
	if i != nil {
		gu.SetBudget(*i)
	}
	return gu
}

// AddBudget adds i to the "budget" field.
func (gu *GroupUpdate) AddBudget(i int) *GroupUpdate {
	gu.mutation.AddBudget(i)
	return gu
}

// ClearBudget clears the value of the "budget" field.
func (gu *GroupUpdate) ClearBudget() *GroupUpdate {
	gu.mutation.ClearBudget()
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GroupUpdate) SetCreatedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetCreatedAt(t)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableCreatedAt(t *time.Time) *GroupUpdate {
	if t != nil {
		gu.SetCreatedAt(*t)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GroupUpdate) SetUpdatedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableUpdatedAt(t *time.Time) *GroupUpdate {
	if t != nil {
		gu.SetUpdatedAt(*t)
	}
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GroupUpdate) SetDeletedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDeletedAt(t *time.Time) *GroupUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gu *GroupUpdate) ClearDeletedAt() *GroupUpdate {
	gu.mutation.ClearDeletedAt()
	return gu
}

// AddGroupBudgetIDs adds the "group_budget" edge to the GroupBudget entity by IDs.
func (gu *GroupUpdate) AddGroupBudgetIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddGroupBudgetIDs(ids...)
	return gu
}

// AddGroupBudget adds the "group_budget" edges to the GroupBudget entity.
func (gu *GroupUpdate) AddGroupBudget(g ...*GroupBudget) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddGroupBudgetIDs(ids...)
}

// AddUserIDs adds the "user" edge to the GroupUser entity by IDs.
func (gu *GroupUpdate) AddUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUser adds the "user" edges to the GroupUser entity.
func (gu *GroupUpdate) AddUser(g ...*GroupUser) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the GroupOwner entity by IDs.
func (gu *GroupUpdate) AddOwnerIDs(ids ...int) *GroupUpdate {
	gu.mutation.AddOwnerIDs(ids...)
	return gu
}

// AddOwner adds the "owner" edges to the GroupOwner entity.
func (gu *GroupUpdate) AddOwner(g ...*GroupOwner) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddOwnerIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearGroupBudget clears all "group_budget" edges to the GroupBudget entity.
func (gu *GroupUpdate) ClearGroupBudget() *GroupUpdate {
	gu.mutation.ClearGroupBudget()
	return gu
}

// RemoveGroupBudgetIDs removes the "group_budget" edge to GroupBudget entities by IDs.
func (gu *GroupUpdate) RemoveGroupBudgetIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveGroupBudgetIDs(ids...)
	return gu
}

// RemoveGroupBudget removes "group_budget" edges to GroupBudget entities.
func (gu *GroupUpdate) RemoveGroupBudget(g ...*GroupBudget) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveGroupBudgetIDs(ids...)
}

// ClearUser clears all "user" edges to the GroupUser entity.
func (gu *GroupUpdate) ClearUser() *GroupUpdate {
	gu.mutation.ClearUser()
	return gu
}

// RemoveUserIDs removes the "user" edge to GroupUser entities by IDs.
func (gu *GroupUpdate) RemoveUserIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUser removes "user" edges to GroupUser entities.
func (gu *GroupUpdate) RemoveUser(g ...*GroupUser) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearOwner clears all "owner" edges to the GroupOwner entity.
func (gu *GroupUpdate) ClearOwner() *GroupUpdate {
	gu.mutation.ClearOwner()
	return gu
}

// RemoveOwnerIDs removes the "owner" edge to GroupOwner entities by IDs.
func (gu *GroupUpdate) RemoveOwnerIDs(ids ...int) *GroupUpdate {
	gu.mutation.RemoveOwnerIDs(ids...)
	return gu
}

// RemoveOwner removes "owner" edges to GroupOwner entities.
func (gu *GroupUpdate) RemoveOwner(g ...*GroupOwner) *GroupUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveOwnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if value, ok := gu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldDescription,
		})
	}
	if value, ok := gu.mutation.Budget(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldBudget,
		})
	}
	if value, ok := gu.mutation.AddedBudget(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldBudget,
		})
	}
	if gu.mutation.BudgetCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: group.FieldBudget,
		})
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldCreatedAt,
		})
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldUpdatedAt,
		})
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldDeletedAt,
		})
	}
	if gu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: group.FieldDeletedAt,
		})
	}
	if gu.mutation.GroupBudgetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedGroupBudgetIDs(); len(nodes) > 0 && !gu.mutation.GroupBudgetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GroupBudgetIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedUserIDs(); len(nodes) > 0 && !gu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !gu.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetDescription sets the "description" field.
func (guo *GroupUpdateOne) SetDescription(s string) *GroupUpdateOne {
	guo.mutation.SetDescription(s)
	return guo
}

// SetBudget sets the "budget" field.
func (guo *GroupUpdateOne) SetBudget(i int) *GroupUpdateOne {
	guo.mutation.ResetBudget()
	guo.mutation.SetBudget(i)
	return guo
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableBudget(i *int) *GroupUpdateOne {
	if i != nil {
		guo.SetBudget(*i)
	}
	return guo
}

// AddBudget adds i to the "budget" field.
func (guo *GroupUpdateOne) AddBudget(i int) *GroupUpdateOne {
	guo.mutation.AddBudget(i)
	return guo
}

// ClearBudget clears the value of the "budget" field.
func (guo *GroupUpdateOne) ClearBudget() *GroupUpdateOne {
	guo.mutation.ClearBudget()
	return guo
}

// SetCreatedAt sets the "created_at" field.
func (guo *GroupUpdateOne) SetCreatedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetCreatedAt(t)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableCreatedAt(t *time.Time) *GroupUpdateOne {
	if t != nil {
		guo.SetCreatedAt(*t)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GroupUpdateOne) SetUpdatedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableUpdatedAt(t *time.Time) *GroupUpdateOne {
	if t != nil {
		guo.SetUpdatedAt(*t)
	}
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GroupUpdateOne) SetDeletedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (guo *GroupUpdateOne) ClearDeletedAt() *GroupUpdateOne {
	guo.mutation.ClearDeletedAt()
	return guo
}

// AddGroupBudgetIDs adds the "group_budget" edge to the GroupBudget entity by IDs.
func (guo *GroupUpdateOne) AddGroupBudgetIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddGroupBudgetIDs(ids...)
	return guo
}

// AddGroupBudget adds the "group_budget" edges to the GroupBudget entity.
func (guo *GroupUpdateOne) AddGroupBudget(g ...*GroupBudget) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddGroupBudgetIDs(ids...)
}

// AddUserIDs adds the "user" edge to the GroupUser entity by IDs.
func (guo *GroupUpdateOne) AddUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUser adds the "user" edges to the GroupUser entity.
func (guo *GroupUpdateOne) AddUser(g ...*GroupUser) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// AddOwnerIDs adds the "owner" edge to the GroupOwner entity by IDs.
func (guo *GroupUpdateOne) AddOwnerIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.AddOwnerIDs(ids...)
	return guo
}

// AddOwner adds the "owner" edges to the GroupOwner entity.
func (guo *GroupUpdateOne) AddOwner(g ...*GroupOwner) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddOwnerIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearGroupBudget clears all "group_budget" edges to the GroupBudget entity.
func (guo *GroupUpdateOne) ClearGroupBudget() *GroupUpdateOne {
	guo.mutation.ClearGroupBudget()
	return guo
}

// RemoveGroupBudgetIDs removes the "group_budget" edge to GroupBudget entities by IDs.
func (guo *GroupUpdateOne) RemoveGroupBudgetIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveGroupBudgetIDs(ids...)
	return guo
}

// RemoveGroupBudget removes "group_budget" edges to GroupBudget entities.
func (guo *GroupUpdateOne) RemoveGroupBudget(g ...*GroupBudget) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveGroupBudgetIDs(ids...)
}

// ClearUser clears all "user" edges to the GroupUser entity.
func (guo *GroupUpdateOne) ClearUser() *GroupUpdateOne {
	guo.mutation.ClearUser()
	return guo
}

// RemoveUserIDs removes the "user" edge to GroupUser entities by IDs.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUser removes "user" edges to GroupUser entities.
func (guo *GroupUpdateOne) RemoveUser(g ...*GroupUser) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearOwner clears all "owner" edges to the GroupOwner entity.
func (guo *GroupUpdateOne) ClearOwner() *GroupUpdateOne {
	guo.mutation.ClearOwner()
	return guo
}

// RemoveOwnerIDs removes the "owner" edge to GroupOwner entities by IDs.
func (guo *GroupUpdateOne) RemoveOwnerIDs(ids ...int) *GroupUpdateOne {
	guo.mutation.RemoveOwnerIDs(ids...)
	return guo
}

// RemoveOwner removes "owner" edges to GroupOwner entities.
func (guo *GroupUpdateOne) RemoveOwner(g ...*GroupOwner) *GroupUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveOwnerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Group.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if value, ok := guo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldDescription,
		})
	}
	if value, ok := guo.mutation.Budget(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldBudget,
		})
	}
	if value, ok := guo.mutation.AddedBudget(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldBudget,
		})
	}
	if guo.mutation.BudgetCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: group.FieldBudget,
		})
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldCreatedAt,
		})
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldUpdatedAt,
		})
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldDeletedAt,
		})
	}
	if guo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: group.FieldDeletedAt,
		})
	}
	if guo.mutation.GroupBudgetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedGroupBudgetIDs(); len(nodes) > 0 && !guo.mutation.GroupBudgetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GroupBudgetIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedUserIDs(); len(nodes) > 0 && !guo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedOwnerIDs(); len(nodes) > 0 && !guo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
