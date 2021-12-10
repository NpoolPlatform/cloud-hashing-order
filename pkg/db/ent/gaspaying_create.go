// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/gaspaying"
	"github.com/google/uuid"
)

// GasPayingCreate is the builder for creating a GasPaying entity.
type GasPayingCreate struct {
	config
	mutation *GasPayingMutation
	hooks    []Hook
}

// SetOrderID sets the "order_id" field.
func (gpc *GasPayingCreate) SetOrderID(u uuid.UUID) *GasPayingCreate {
	gpc.mutation.SetOrderID(u)
	return gpc
}

// SetFeeTypeID sets the "fee_type_id" field.
func (gpc *GasPayingCreate) SetFeeTypeID(u uuid.UUID) *GasPayingCreate {
	gpc.mutation.SetFeeTypeID(u)
	return gpc
}

// SetPaymentID sets the "payment_id" field.
func (gpc *GasPayingCreate) SetPaymentID(u uuid.UUID) *GasPayingCreate {
	gpc.mutation.SetPaymentID(u)
	return gpc
}

// SetDurationMinutes sets the "duration_minutes" field.
func (gpc *GasPayingCreate) SetDurationMinutes(u uint32) *GasPayingCreate {
	gpc.mutation.SetDurationMinutes(u)
	return gpc
}

// SetCreateAt sets the "create_at" field.
func (gpc *GasPayingCreate) SetCreateAt(u uint32) *GasPayingCreate {
	gpc.mutation.SetCreateAt(u)
	return gpc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (gpc *GasPayingCreate) SetNillableCreateAt(u *uint32) *GasPayingCreate {
	if u != nil {
		gpc.SetCreateAt(*u)
	}
	return gpc
}

// SetUpdateAt sets the "update_at" field.
func (gpc *GasPayingCreate) SetUpdateAt(u uint32) *GasPayingCreate {
	gpc.mutation.SetUpdateAt(u)
	return gpc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (gpc *GasPayingCreate) SetNillableUpdateAt(u *uint32) *GasPayingCreate {
	if u != nil {
		gpc.SetUpdateAt(*u)
	}
	return gpc
}

// SetDeleteAt sets the "delete_at" field.
func (gpc *GasPayingCreate) SetDeleteAt(u uint32) *GasPayingCreate {
	gpc.mutation.SetDeleteAt(u)
	return gpc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (gpc *GasPayingCreate) SetNillableDeleteAt(u *uint32) *GasPayingCreate {
	if u != nil {
		gpc.SetDeleteAt(*u)
	}
	return gpc
}

// SetID sets the "id" field.
func (gpc *GasPayingCreate) SetID(u uuid.UUID) *GasPayingCreate {
	gpc.mutation.SetID(u)
	return gpc
}

// Mutation returns the GasPayingMutation object of the builder.
func (gpc *GasPayingCreate) Mutation() *GasPayingMutation {
	return gpc.mutation
}

// Save creates the GasPaying in the database.
func (gpc *GasPayingCreate) Save(ctx context.Context) (*GasPaying, error) {
	var (
		err  error
		node *GasPaying
	)
	gpc.defaults()
	if len(gpc.hooks) == 0 {
		if err = gpc.check(); err != nil {
			return nil, err
		}
		node, err = gpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GasPayingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gpc.check(); err != nil {
				return nil, err
			}
			gpc.mutation = mutation
			if node, err = gpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gpc.hooks) - 1; i >= 0; i-- {
			if gpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gpc *GasPayingCreate) SaveX(ctx context.Context) *GasPaying {
	v, err := gpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gpc *GasPayingCreate) Exec(ctx context.Context) error {
	_, err := gpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpc *GasPayingCreate) ExecX(ctx context.Context) {
	if err := gpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gpc *GasPayingCreate) defaults() {
	if _, ok := gpc.mutation.CreateAt(); !ok {
		v := gaspaying.DefaultCreateAt()
		gpc.mutation.SetCreateAt(v)
	}
	if _, ok := gpc.mutation.UpdateAt(); !ok {
		v := gaspaying.DefaultUpdateAt()
		gpc.mutation.SetUpdateAt(v)
	}
	if _, ok := gpc.mutation.DeleteAt(); !ok {
		v := gaspaying.DefaultDeleteAt()
		gpc.mutation.SetDeleteAt(v)
	}
	if _, ok := gpc.mutation.ID(); !ok {
		v := gaspaying.DefaultID()
		gpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gpc *GasPayingCreate) check() error {
	if _, ok := gpc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "order_id"`)}
	}
	if _, ok := gpc.mutation.FeeTypeID(); !ok {
		return &ValidationError{Name: "fee_type_id", err: errors.New(`ent: missing required field "fee_type_id"`)}
	}
	if _, ok := gpc.mutation.PaymentID(); !ok {
		return &ValidationError{Name: "payment_id", err: errors.New(`ent: missing required field "payment_id"`)}
	}
	if _, ok := gpc.mutation.DurationMinutes(); !ok {
		return &ValidationError{Name: "duration_minutes", err: errors.New(`ent: missing required field "duration_minutes"`)}
	}
	if _, ok := gpc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := gpc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := gpc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (gpc *GasPayingCreate) sqlSave(ctx context.Context) (*GasPaying, error) {
	_node, _spec := gpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (gpc *GasPayingCreate) createSpec() (*GasPaying, *sqlgraph.CreateSpec) {
	var (
		_node = &GasPaying{config: gpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gaspaying.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gaspaying.FieldID,
			},
		}
	)
	if id, ok := gpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gpc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gaspaying.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := gpc.mutation.FeeTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gaspaying.FieldFeeTypeID,
		})
		_node.FeeTypeID = value
	}
	if value, ok := gpc.mutation.PaymentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: gaspaying.FieldPaymentID,
		})
		_node.PaymentID = value
	}
	if value, ok := gpc.mutation.DurationMinutes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gaspaying.FieldDurationMinutes,
		})
		_node.DurationMinutes = value
	}
	if value, ok := gpc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gaspaying.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := gpc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gaspaying.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := gpc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: gaspaying.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// GasPayingCreateBulk is the builder for creating many GasPaying entities in bulk.
type GasPayingCreateBulk struct {
	config
	builders []*GasPayingCreate
}

// Save creates the GasPaying entities in the database.
func (gpcb *GasPayingCreateBulk) Save(ctx context.Context) ([]*GasPaying, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gpcb.builders))
	nodes := make([]*GasPaying, len(gpcb.builders))
	mutators := make([]Mutator, len(gpcb.builders))
	for i := range gpcb.builders {
		func(i int, root context.Context) {
			builder := gpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GasPayingMutation)
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
					_, err = mutators[i+1].Mutate(root, gpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gpcb *GasPayingCreateBulk) SaveX(ctx context.Context) []*GasPaying {
	v, err := gpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gpcb *GasPayingCreateBulk) Exec(ctx context.Context) error {
	_, err := gpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpcb *GasPayingCreateBulk) ExecX(ctx context.Context) {
	if err := gpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
