// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/traPtitech/Jomon/ent/file"
	"github.com/traPtitech/Jomon/ent/requestfile"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetMimeType sets the "mime_type" field.
func (fc *FileCreate) SetMimeType(s string) *FileCreate {
	fc.mutation.SetMimeType(s)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FileCreate) SetCreatedAt(t time.Time) *FileCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetDeletedAt sets the "deleted_at" field.
func (fc *FileCreate) SetDeletedAt(t time.Time) *FileCreate {
	fc.mutation.SetDeletedAt(t)
	return fc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableDeletedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetDeletedAt(*t)
	}
	return fc
}

// SetRequestFileID sets the "request_file" edge to the RequestFile entity by ID.
func (fc *FileCreate) SetRequestFileID(id int) *FileCreate {
	fc.mutation.SetRequestFileID(id)
	return fc
}

// SetNillableRequestFileID sets the "request_file" edge to the RequestFile entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableRequestFileID(id *int) *FileCreate {
	if id != nil {
		fc = fc.SetRequestFileID(*id)
	}
	return fc
}

// SetRequestFile sets the "request_file" edge to the RequestFile entity.
func (fc *FileCreate) SetRequestFile(r *RequestFile) *FileCreate {
	return fc.SetRequestFileID(r.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := file.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.MimeType(); !ok {
		return &ValidationError{Name: "mime_type", err: errors.New("ent: missing required field \"mime_type\"")}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: file.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: file.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.MimeType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldMimeType,
		})
		_node.MimeType = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := fc.mutation.RequestFileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   file.RequestFileTable,
			Columns: []string{file.RequestFileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: requestfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.request_file_file = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	builders []*FileCreate
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}