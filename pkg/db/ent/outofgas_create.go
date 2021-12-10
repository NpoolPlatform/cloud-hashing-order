// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/outofgas"
	"github.com/google/uuid"
)

// OutOfGasCreate is the builder for creating a OutOfGas entity.
type OutOfGasCreate struct {
	config
	mutation *OutOfGasMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrderID sets the "order_id" field.
func (oogc *OutOfGasCreate) SetOrderID(u uuid.UUID) *OutOfGasCreate {
	oogc.mutation.SetOrderID(u)
	return oogc
}

// SetStart sets the "start" field.
func (oogc *OutOfGasCreate) SetStart(u uint32) *OutOfGasCreate {
	oogc.mutation.SetStart(u)
	return oogc
}

// SetEnd sets the "end" field.
func (oogc *OutOfGasCreate) SetEnd(u uint32) *OutOfGasCreate {
	oogc.mutation.SetEnd(u)
	return oogc
}

// SetCreateAt sets the "create_at" field.
func (oogc *OutOfGasCreate) SetCreateAt(u uint32) *OutOfGasCreate {
	oogc.mutation.SetCreateAt(u)
	return oogc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (oogc *OutOfGasCreate) SetNillableCreateAt(u *uint32) *OutOfGasCreate {
	if u != nil {
		oogc.SetCreateAt(*u)
	}
	return oogc
}

// SetUpdateAt sets the "update_at" field.
func (oogc *OutOfGasCreate) SetUpdateAt(u uint32) *OutOfGasCreate {
	oogc.mutation.SetUpdateAt(u)
	return oogc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (oogc *OutOfGasCreate) SetNillableUpdateAt(u *uint32) *OutOfGasCreate {
	if u != nil {
		oogc.SetUpdateAt(*u)
	}
	return oogc
}

// SetDeleteAt sets the "delete_at" field.
func (oogc *OutOfGasCreate) SetDeleteAt(u uint32) *OutOfGasCreate {
	oogc.mutation.SetDeleteAt(u)
	return oogc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (oogc *OutOfGasCreate) SetNillableDeleteAt(u *uint32) *OutOfGasCreate {
	if u != nil {
		oogc.SetDeleteAt(*u)
	}
	return oogc
}

// SetID sets the "id" field.
func (oogc *OutOfGasCreate) SetID(u uuid.UUID) *OutOfGasCreate {
	oogc.mutation.SetID(u)
	return oogc
}

// Mutation returns the OutOfGasMutation object of the builder.
func (oogc *OutOfGasCreate) Mutation() *OutOfGasMutation {
	return oogc.mutation
}

// Save creates the OutOfGas in the database.
func (oogc *OutOfGasCreate) Save(ctx context.Context) (*OutOfGas, error) {
	var (
		err  error
		node *OutOfGas
	)
	oogc.defaults()
	if len(oogc.hooks) == 0 {
		if err = oogc.check(); err != nil {
			return nil, err
		}
		node, err = oogc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutOfGasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oogc.check(); err != nil {
				return nil, err
			}
			oogc.mutation = mutation
			if node, err = oogc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oogc.hooks) - 1; i >= 0; i-- {
			if oogc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oogc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oogc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oogc *OutOfGasCreate) SaveX(ctx context.Context) *OutOfGas {
	v, err := oogc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oogc *OutOfGasCreate) Exec(ctx context.Context) error {
	_, err := oogc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oogc *OutOfGasCreate) ExecX(ctx context.Context) {
	if err := oogc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oogc *OutOfGasCreate) defaults() {
	if _, ok := oogc.mutation.CreateAt(); !ok {
		v := outofgas.DefaultCreateAt()
		oogc.mutation.SetCreateAt(v)
	}
	if _, ok := oogc.mutation.UpdateAt(); !ok {
		v := outofgas.DefaultUpdateAt()
		oogc.mutation.SetUpdateAt(v)
	}
	if _, ok := oogc.mutation.DeleteAt(); !ok {
		v := outofgas.DefaultDeleteAt()
		oogc.mutation.SetDeleteAt(v)
	}
	if _, ok := oogc.mutation.ID(); !ok {
		v := outofgas.DefaultID()
		oogc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oogc *OutOfGasCreate) check() error {
	if _, ok := oogc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "order_id"`)}
	}
	if _, ok := oogc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "start"`)}
	}
	if _, ok := oogc.mutation.End(); !ok {
		return &ValidationError{Name: "end", err: errors.New(`ent: missing required field "end"`)}
	}
	if _, ok := oogc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := oogc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := oogc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (oogc *OutOfGasCreate) sqlSave(ctx context.Context) (*OutOfGas, error) {
	_node, _spec := oogc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oogc.driver, _spec); err != nil {
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

func (oogc *OutOfGasCreate) createSpec() (*OutOfGas, *sqlgraph.CreateSpec) {
	var (
		_node = &OutOfGas{config: oogc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outofgas.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		}
	)
	_spec.OnConflict = oogc.conflict
	if id, ok := oogc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oogc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: outofgas.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := oogc.mutation.Start(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
		_node.Start = value
	}
	if value, ok := oogc.mutation.End(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
		_node.End = value
	}
	if value, ok := oogc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := oogc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := oogc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OutOfGas.Create().
//		SetOrderID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OutOfGasUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
//
func (oogc *OutOfGasCreate) OnConflict(opts ...sql.ConflictOption) *OutOfGasUpsertOne {
	oogc.conflict = opts
	return &OutOfGasUpsertOne{
		create: oogc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OutOfGas.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (oogc *OutOfGasCreate) OnConflictColumns(columns ...string) *OutOfGasUpsertOne {
	oogc.conflict = append(oogc.conflict, sql.ConflictColumns(columns...))
	return &OutOfGasUpsertOne{
		create: oogc,
	}
}

type (
	// OutOfGasUpsertOne is the builder for "upsert"-ing
	//  one OutOfGas node.
	OutOfGasUpsertOne struct {
		create *OutOfGasCreate
	}

	// OutOfGasUpsert is the "OnConflict" setter.
	OutOfGasUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrderID sets the "order_id" field.
func (u *OutOfGasUpsert) SetOrderID(v uuid.UUID) *OutOfGasUpsert {
	u.Set(outofgas.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OutOfGasUpsert) UpdateOrderID() *OutOfGasUpsert {
	u.SetExcluded(outofgas.FieldOrderID)
	return u
}

// SetStart sets the "start" field.
func (u *OutOfGasUpsert) SetStart(v uint32) *OutOfGasUpsert {
	u.Set(outofgas.FieldStart, v)
	return u
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *OutOfGasUpsert) UpdateStart() *OutOfGasUpsert {
	u.SetExcluded(outofgas.FieldStart)
	return u
}

// SetEnd sets the "end" field.
func (u *OutOfGasUpsert) SetEnd(v uint32) *OutOfGasUpsert {
	u.Set(outofgas.FieldEnd, v)
	return u
}

// UpdateEnd sets the "end" field to the value that was provided on create.
func (u *OutOfGasUpsert) UpdateEnd() *OutOfGasUpsert {
	u.SetExcluded(outofgas.FieldEnd)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *OutOfGasUpsert) SetCreateAt(v uint32) *OutOfGasUpsert {
	u.Set(outofgas.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *OutOfGasUpsert) UpdateCreateAt() *OutOfGasUpsert {
	u.SetExcluded(outofgas.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *OutOfGasUpsert) SetUpdateAt(v uint32) *OutOfGasUpsert {
	u.Set(outofgas.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *OutOfGasUpsert) UpdateUpdateAt() *OutOfGasUpsert {
	u.SetExcluded(outofgas.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *OutOfGasUpsert) SetDeleteAt(v uint32) *OutOfGasUpsert {
	u.Set(outofgas.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *OutOfGasUpsert) UpdateDeleteAt() *OutOfGasUpsert {
	u.SetExcluded(outofgas.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OutOfGas.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(outofgas.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *OutOfGasUpsertOne) UpdateNewValues() *OutOfGasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(outofgas.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.OutOfGas.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *OutOfGasUpsertOne) Ignore() *OutOfGasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OutOfGasUpsertOne) DoNothing() *OutOfGasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OutOfGasCreate.OnConflict
// documentation for more info.
func (u *OutOfGasUpsertOne) Update(set func(*OutOfGasUpsert)) *OutOfGasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OutOfGasUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OutOfGasUpsertOne) SetOrderID(v uuid.UUID) *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OutOfGasUpsertOne) UpdateOrderID() *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateOrderID()
	})
}

// SetStart sets the "start" field.
func (u *OutOfGasUpsertOne) SetStart(v uint32) *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *OutOfGasUpsertOne) UpdateStart() *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateStart()
	})
}

// SetEnd sets the "end" field.
func (u *OutOfGasUpsertOne) SetEnd(v uint32) *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetEnd(v)
	})
}

// UpdateEnd sets the "end" field to the value that was provided on create.
func (u *OutOfGasUpsertOne) UpdateEnd() *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateEnd()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *OutOfGasUpsertOne) SetCreateAt(v uint32) *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *OutOfGasUpsertOne) UpdateCreateAt() *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *OutOfGasUpsertOne) SetUpdateAt(v uint32) *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *OutOfGasUpsertOne) UpdateUpdateAt() *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *OutOfGasUpsertOne) SetDeleteAt(v uint32) *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *OutOfGasUpsertOne) UpdateDeleteAt() *OutOfGasUpsertOne {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *OutOfGasUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OutOfGasCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OutOfGasUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OutOfGasUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: OutOfGasUpsertOne.ID is not supported by MySQL driver. Use OutOfGasUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OutOfGasUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OutOfGasCreateBulk is the builder for creating many OutOfGas entities in bulk.
type OutOfGasCreateBulk struct {
	config
	builders []*OutOfGasCreate
	conflict []sql.ConflictOption
}

// Save creates the OutOfGas entities in the database.
func (oogcb *OutOfGasCreateBulk) Save(ctx context.Context) ([]*OutOfGas, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oogcb.builders))
	nodes := make([]*OutOfGas, len(oogcb.builders))
	mutators := make([]Mutator, len(oogcb.builders))
	for i := range oogcb.builders {
		func(i int, root context.Context) {
			builder := oogcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutOfGasMutation)
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
					_, err = mutators[i+1].Mutate(root, oogcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = oogcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oogcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oogcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oogcb *OutOfGasCreateBulk) SaveX(ctx context.Context) []*OutOfGas {
	v, err := oogcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oogcb *OutOfGasCreateBulk) Exec(ctx context.Context) error {
	_, err := oogcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oogcb *OutOfGasCreateBulk) ExecX(ctx context.Context) {
	if err := oogcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OutOfGas.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OutOfGasUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
//
func (oogcb *OutOfGasCreateBulk) OnConflict(opts ...sql.ConflictOption) *OutOfGasUpsertBulk {
	oogcb.conflict = opts
	return &OutOfGasUpsertBulk{
		create: oogcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OutOfGas.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (oogcb *OutOfGasCreateBulk) OnConflictColumns(columns ...string) *OutOfGasUpsertBulk {
	oogcb.conflict = append(oogcb.conflict, sql.ConflictColumns(columns...))
	return &OutOfGasUpsertBulk{
		create: oogcb,
	}
}

// OutOfGasUpsertBulk is the builder for "upsert"-ing
// a bulk of OutOfGas nodes.
type OutOfGasUpsertBulk struct {
	create *OutOfGasCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OutOfGas.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(outofgas.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *OutOfGasUpsertBulk) UpdateNewValues() *OutOfGasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(outofgas.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OutOfGas.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *OutOfGasUpsertBulk) Ignore() *OutOfGasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OutOfGasUpsertBulk) DoNothing() *OutOfGasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OutOfGasCreateBulk.OnConflict
// documentation for more info.
func (u *OutOfGasUpsertBulk) Update(set func(*OutOfGasUpsert)) *OutOfGasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OutOfGasUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OutOfGasUpsertBulk) SetOrderID(v uuid.UUID) *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OutOfGasUpsertBulk) UpdateOrderID() *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateOrderID()
	})
}

// SetStart sets the "start" field.
func (u *OutOfGasUpsertBulk) SetStart(v uint32) *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *OutOfGasUpsertBulk) UpdateStart() *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateStart()
	})
}

// SetEnd sets the "end" field.
func (u *OutOfGasUpsertBulk) SetEnd(v uint32) *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetEnd(v)
	})
}

// UpdateEnd sets the "end" field to the value that was provided on create.
func (u *OutOfGasUpsertBulk) UpdateEnd() *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateEnd()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *OutOfGasUpsertBulk) SetCreateAt(v uint32) *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *OutOfGasUpsertBulk) UpdateCreateAt() *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *OutOfGasUpsertBulk) SetUpdateAt(v uint32) *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *OutOfGasUpsertBulk) UpdateUpdateAt() *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *OutOfGasUpsertBulk) SetDeleteAt(v uint32) *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *OutOfGasUpsertBulk) UpdateDeleteAt() *OutOfGasUpsertBulk {
	return u.Update(func(s *OutOfGasUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *OutOfGasUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OutOfGasCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OutOfGasCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OutOfGasUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
