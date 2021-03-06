// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/gaspaying"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GasPayingQuery is the builder for querying GasPaying entities.
type GasPayingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GasPaying
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GasPayingQuery builder.
func (gpq *GasPayingQuery) Where(ps ...predicate.GasPaying) *GasPayingQuery {
	gpq.predicates = append(gpq.predicates, ps...)
	return gpq
}

// Limit adds a limit step to the query.
func (gpq *GasPayingQuery) Limit(limit int) *GasPayingQuery {
	gpq.limit = &limit
	return gpq
}

// Offset adds an offset step to the query.
func (gpq *GasPayingQuery) Offset(offset int) *GasPayingQuery {
	gpq.offset = &offset
	return gpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gpq *GasPayingQuery) Unique(unique bool) *GasPayingQuery {
	gpq.unique = &unique
	return gpq
}

// Order adds an order step to the query.
func (gpq *GasPayingQuery) Order(o ...OrderFunc) *GasPayingQuery {
	gpq.order = append(gpq.order, o...)
	return gpq
}

// First returns the first GasPaying entity from the query.
// Returns a *NotFoundError when no GasPaying was found.
func (gpq *GasPayingQuery) First(ctx context.Context) (*GasPaying, error) {
	nodes, err := gpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gaspaying.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gpq *GasPayingQuery) FirstX(ctx context.Context) *GasPaying {
	node, err := gpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GasPaying ID from the query.
// Returns a *NotFoundError when no GasPaying ID was found.
func (gpq *GasPayingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gaspaying.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gpq *GasPayingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GasPaying entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GasPaying entity is found.
// Returns a *NotFoundError when no GasPaying entities are found.
func (gpq *GasPayingQuery) Only(ctx context.Context) (*GasPaying, error) {
	nodes, err := gpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gaspaying.Label}
	default:
		return nil, &NotSingularError{gaspaying.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gpq *GasPayingQuery) OnlyX(ctx context.Context) *GasPaying {
	node, err := gpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GasPaying ID in the query.
// Returns a *NotSingularError when more than one GasPaying ID is found.
// Returns a *NotFoundError when no entities are found.
func (gpq *GasPayingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = &NotSingularError{gaspaying.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gpq *GasPayingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GasPayings.
func (gpq *GasPayingQuery) All(ctx context.Context) ([]*GasPaying, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gpq *GasPayingQuery) AllX(ctx context.Context) []*GasPaying {
	nodes, err := gpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GasPaying IDs.
func (gpq *GasPayingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gpq.Select(gaspaying.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gpq *GasPayingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gpq *GasPayingQuery) Count(ctx context.Context) (int, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gpq *GasPayingQuery) CountX(ctx context.Context) int {
	count, err := gpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gpq *GasPayingQuery) Exist(ctx context.Context) (bool, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gpq *GasPayingQuery) ExistX(ctx context.Context) bool {
	exist, err := gpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GasPayingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gpq *GasPayingQuery) Clone() *GasPayingQuery {
	if gpq == nil {
		return nil
	}
	return &GasPayingQuery{
		config:     gpq.config,
		limit:      gpq.limit,
		offset:     gpq.offset,
		order:      append([]OrderFunc{}, gpq.order...),
		predicates: append([]predicate.GasPaying{}, gpq.predicates...),
		// clone intermediate query.
		sql:    gpq.sql.Clone(),
		path:   gpq.path,
		unique: gpq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderID uuid.UUID `json:"order_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GasPaying.Query().
//		GroupBy(gaspaying.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gpq *GasPayingQuery) GroupBy(field string, fields ...string) *GasPayingGroupBy {
	group := &GasPayingGroupBy{config: gpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderID uuid.UUID `json:"order_id,omitempty"`
//	}
//
//	client.GasPaying.Query().
//		Select(gaspaying.FieldOrderID).
//		Scan(ctx, &v)
//
func (gpq *GasPayingQuery) Select(fields ...string) *GasPayingSelect {
	gpq.fields = append(gpq.fields, fields...)
	return &GasPayingSelect{GasPayingQuery: gpq}
}

func (gpq *GasPayingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gpq.fields {
		if !gaspaying.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gpq.path != nil {
		prev, err := gpq.path(ctx)
		if err != nil {
			return err
		}
		gpq.sql = prev
	}
	return nil
}

func (gpq *GasPayingQuery) sqlAll(ctx context.Context) ([]*GasPaying, error) {
	var (
		nodes = []*GasPaying{}
		_spec = gpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GasPaying{config: gpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, gpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gpq *GasPayingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gpq.querySpec()
	_spec.Node.Columns = gpq.fields
	if len(gpq.fields) > 0 {
		_spec.Unique = gpq.unique != nil && *gpq.unique
	}
	return sqlgraph.CountNodes(ctx, gpq.driver, _spec)
}

func (gpq *GasPayingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gpq *GasPayingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gaspaying.Table,
			Columns: gaspaying.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gaspaying.FieldID,
			},
		},
		From:   gpq.sql,
		Unique: true,
	}
	if unique := gpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gaspaying.FieldID)
		for i := range fields {
			if fields[i] != gaspaying.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gpq *GasPayingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gpq.driver.Dialect())
	t1 := builder.Table(gaspaying.Table)
	columns := gpq.fields
	if len(columns) == 0 {
		columns = gaspaying.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gpq.sql != nil {
		selector = gpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gpq.unique != nil && *gpq.unique {
		selector.Distinct()
	}
	for _, p := range gpq.predicates {
		p(selector)
	}
	for _, p := range gpq.order {
		p(selector)
	}
	if offset := gpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GasPayingGroupBy is the group-by builder for GasPaying entities.
type GasPayingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gpgb *GasPayingGroupBy) Aggregate(fns ...AggregateFunc) *GasPayingGroupBy {
	gpgb.fns = append(gpgb.fns, fns...)
	return gpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gpgb *GasPayingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gpgb.path(ctx)
	if err != nil {
		return err
	}
	gpgb.sql = query
	return gpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gpgb.fields) > 1 {
		return nil, errors.New("ent: GasPayingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) StringsX(ctx context.Context) []string {
	v, err := gpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) StringX(ctx context.Context) string {
	v, err := gpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gpgb.fields) > 1 {
		return nil, errors.New("ent: GasPayingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) IntsX(ctx context.Context) []int {
	v, err := gpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) IntX(ctx context.Context) int {
	v, err := gpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gpgb.fields) > 1 {
		return nil, errors.New("ent: GasPayingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gpgb.fields) > 1 {
		return nil, errors.New("ent: GasPayingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gpgb *GasPayingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gpgb *GasPayingGroupBy) BoolX(ctx context.Context) bool {
	v, err := gpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gpgb *GasPayingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gpgb.fields {
		if !gaspaying.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gpgb *GasPayingGroupBy) sqlQuery() *sql.Selector {
	selector := gpgb.sql.Select()
	aggregation := make([]string, 0, len(gpgb.fns))
	for _, fn := range gpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gpgb.fields)+len(gpgb.fns))
		for _, f := range gpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gpgb.fields...)...)
}

// GasPayingSelect is the builder for selecting fields of GasPaying entities.
type GasPayingSelect struct {
	*GasPayingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gps *GasPayingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gps.prepareQuery(ctx); err != nil {
		return err
	}
	gps.sql = gps.GasPayingQuery.sqlQuery(ctx)
	return gps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gps *GasPayingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gps.fields) > 1 {
		return nil, errors.New("ent: GasPayingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gps *GasPayingSelect) StringsX(ctx context.Context) []string {
	v, err := gps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gps *GasPayingSelect) StringX(ctx context.Context) string {
	v, err := gps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gps.fields) > 1 {
		return nil, errors.New("ent: GasPayingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gps *GasPayingSelect) IntsX(ctx context.Context) []int {
	v, err := gps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gps *GasPayingSelect) IntX(ctx context.Context) int {
	v, err := gps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gps.fields) > 1 {
		return nil, errors.New("ent: GasPayingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gps *GasPayingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gps *GasPayingSelect) Float64X(ctx context.Context) float64 {
	v, err := gps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gps.fields) > 1 {
		return nil, errors.New("ent: GasPayingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gps *GasPayingSelect) BoolsX(ctx context.Context) []bool {
	v, err := gps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gps *GasPayingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gaspaying.Label}
	default:
		err = fmt.Errorf("ent: GasPayingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gps *GasPayingSelect) BoolX(ctx context.Context) bool {
	v, err := gps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gps *GasPayingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gps.sql.Query()
	if err := gps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
