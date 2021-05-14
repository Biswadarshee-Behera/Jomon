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
	"github.com/traPtitech/Jomon/ent/predicate"
	"github.com/traPtitech/Jomon/ent/tag"
	"github.com/traPtitech/Jomon/ent/transaction"
	"github.com/traPtitech/Jomon/ent/transactiontag"
)

// TransactionTagUpdate is the builder for updating TransactionTag entities.
type TransactionTagUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionTagMutation
}

// Where adds a new predicate for the TransactionTagUpdate builder.
func (ttu *TransactionTagUpdate) Where(ps ...predicate.TransactionTag) *TransactionTagUpdate {
	ttu.mutation.predicates = append(ttu.mutation.predicates, ps...)
	return ttu
}

// SetTransactionID sets the "transaction_id" field.
func (ttu *TransactionTagUpdate) SetTransactionID(i int) *TransactionTagUpdate {
	ttu.mutation.ResetTransactionID()
	ttu.mutation.SetTransactionID(i)
	return ttu
}

// SetCreatedAt sets the "created_at" field.
func (ttu *TransactionTagUpdate) SetCreatedAt(t time.Time) *TransactionTagUpdate {
	ttu.mutation.SetCreatedAt(t)
	return ttu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttu *TransactionTagUpdate) SetNillableCreatedAt(t *time.Time) *TransactionTagUpdate {
	if t != nil {
		ttu.SetCreatedAt(*t)
	}
	return ttu
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (ttu *TransactionTagUpdate) SetTransaction(t *Transaction) *TransactionTagUpdate {
	return ttu.SetTransactionID(t.ID)
}

// SetTagID sets the "tag" edge to the Tag entity by ID.
func (ttu *TransactionTagUpdate) SetTagID(id int) *TransactionTagUpdate {
	ttu.mutation.SetTagID(id)
	return ttu
}

// SetNillableTagID sets the "tag" edge to the Tag entity by ID if the given value is not nil.
func (ttu *TransactionTagUpdate) SetNillableTagID(id *int) *TransactionTagUpdate {
	if id != nil {
		ttu = ttu.SetTagID(*id)
	}
	return ttu
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttu *TransactionTagUpdate) SetTag(t *Tag) *TransactionTagUpdate {
	return ttu.SetTagID(t.ID)
}

// Mutation returns the TransactionTagMutation object of the builder.
func (ttu *TransactionTagUpdate) Mutation() *TransactionTagMutation {
	return ttu.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (ttu *TransactionTagUpdate) ClearTransaction() *TransactionTagUpdate {
	ttu.mutation.ClearTransaction()
	return ttu
}

// ClearTag clears the "tag" edge to the Tag entity.
func (ttu *TransactionTagUpdate) ClearTag() *TransactionTagUpdate {
	ttu.mutation.ClearTag()
	return ttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TransactionTagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttu.hooks) == 0 {
		if err = ttu.check(); err != nil {
			return 0, err
		}
		affected, err = ttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttu.check(); err != nil {
				return 0, err
			}
			ttu.mutation = mutation
			affected, err = ttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttu.hooks) - 1; i >= 0; i-- {
			mut = ttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TransactionTagUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TransactionTagUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TransactionTagUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttu *TransactionTagUpdate) check() error {
	if _, ok := ttu.mutation.TransactionID(); ttu.mutation.TransactionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"transaction\"")
	}
	return nil
}

func (ttu *TransactionTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactiontag.Table,
			Columns: transactiontag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactiontag.FieldID,
			},
		},
	}
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactiontag.FieldCreatedAt,
		})
	}
	if ttu.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transactiontag.TransactionTable,
			Columns: []string{transactiontag.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transactiontag.TransactionTable,
			Columns: []string{transactiontag.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ttu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transactiontag.TagTable,
			Columns: []string{transactiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transactiontag.TagTable,
			Columns: []string{transactiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiontag.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TransactionTagUpdateOne is the builder for updating a single TransactionTag entity.
type TransactionTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionTagMutation
}

// SetTransactionID sets the "transaction_id" field.
func (ttuo *TransactionTagUpdateOne) SetTransactionID(i int) *TransactionTagUpdateOne {
	ttuo.mutation.ResetTransactionID()
	ttuo.mutation.SetTransactionID(i)
	return ttuo
}

// SetCreatedAt sets the "created_at" field.
func (ttuo *TransactionTagUpdateOne) SetCreatedAt(t time.Time) *TransactionTagUpdateOne {
	ttuo.mutation.SetCreatedAt(t)
	return ttuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ttuo *TransactionTagUpdateOne) SetNillableCreatedAt(t *time.Time) *TransactionTagUpdateOne {
	if t != nil {
		ttuo.SetCreatedAt(*t)
	}
	return ttuo
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (ttuo *TransactionTagUpdateOne) SetTransaction(t *Transaction) *TransactionTagUpdateOne {
	return ttuo.SetTransactionID(t.ID)
}

// SetTagID sets the "tag" edge to the Tag entity by ID.
func (ttuo *TransactionTagUpdateOne) SetTagID(id int) *TransactionTagUpdateOne {
	ttuo.mutation.SetTagID(id)
	return ttuo
}

// SetNillableTagID sets the "tag" edge to the Tag entity by ID if the given value is not nil.
func (ttuo *TransactionTagUpdateOne) SetNillableTagID(id *int) *TransactionTagUpdateOne {
	if id != nil {
		ttuo = ttuo.SetTagID(*id)
	}
	return ttuo
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttuo *TransactionTagUpdateOne) SetTag(t *Tag) *TransactionTagUpdateOne {
	return ttuo.SetTagID(t.ID)
}

// Mutation returns the TransactionTagMutation object of the builder.
func (ttuo *TransactionTagUpdateOne) Mutation() *TransactionTagMutation {
	return ttuo.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (ttuo *TransactionTagUpdateOne) ClearTransaction() *TransactionTagUpdateOne {
	ttuo.mutation.ClearTransaction()
	return ttuo
}

// ClearTag clears the "tag" edge to the Tag entity.
func (ttuo *TransactionTagUpdateOne) ClearTag() *TransactionTagUpdateOne {
	ttuo.mutation.ClearTag()
	return ttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TransactionTagUpdateOne) Select(field string, fields ...string) *TransactionTagUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TransactionTag entity.
func (ttuo *TransactionTagUpdateOne) Save(ctx context.Context) (*TransactionTag, error) {
	var (
		err  error
		node *TransactionTag
	)
	if len(ttuo.hooks) == 0 {
		if err = ttuo.check(); err != nil {
			return nil, err
		}
		node, err = ttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttuo.check(); err != nil {
				return nil, err
			}
			ttuo.mutation = mutation
			node, err = ttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttuo.hooks) - 1; i >= 0; i-- {
			mut = ttuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TransactionTagUpdateOne) SaveX(ctx context.Context) *TransactionTag {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TransactionTagUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TransactionTagUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttuo *TransactionTagUpdateOne) check() error {
	if _, ok := ttuo.mutation.TransactionID(); ttuo.mutation.TransactionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"transaction\"")
	}
	return nil
}

func (ttuo *TransactionTagUpdateOne) sqlSave(ctx context.Context) (_node *TransactionTag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactiontag.Table,
			Columns: transactiontag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactiontag.FieldID,
			},
		},
	}
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TransactionTag.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactiontag.FieldID)
		for _, f := range fields {
			if !transactiontag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transactiontag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactiontag.FieldCreatedAt,
		})
	}
	if ttuo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transactiontag.TransactionTable,
			Columns: []string{transactiontag.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transactiontag.TransactionTable,
			Columns: []string{transactiontag.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ttuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transactiontag.TagTable,
			Columns: []string{transactiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transactiontag.TagTable,
			Columns: []string{transactiontag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TransactionTag{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiontag.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}