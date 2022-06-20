// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/YadaYuki/omochi/app/ent/invertindexcompressed"
	"github.com/YadaYuki/omochi/app/ent/predicate"
	"github.com/YadaYuki/omochi/app/ent/term"
	"github.com/google/uuid"
)

// InvertIndexCompressedQuery is the builder for querying InvertIndexCompressed entities.
type InvertIndexCompressedQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.InvertIndexCompressed
	// eager-loading edges.
	withTerm *TermQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvertIndexCompressedQuery builder.
func (iicq *InvertIndexCompressedQuery) Where(ps ...predicate.InvertIndexCompressed) *InvertIndexCompressedQuery {
	iicq.predicates = append(iicq.predicates, ps...)
	return iicq
}

// Limit adds a limit step to the query.
func (iicq *InvertIndexCompressedQuery) Limit(limit int) *InvertIndexCompressedQuery {
	iicq.limit = &limit
	return iicq
}

// Offset adds an offset step to the query.
func (iicq *InvertIndexCompressedQuery) Offset(offset int) *InvertIndexCompressedQuery {
	iicq.offset = &offset
	return iicq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iicq *InvertIndexCompressedQuery) Unique(unique bool) *InvertIndexCompressedQuery {
	iicq.unique = &unique
	return iicq
}

// Order adds an order step to the query.
func (iicq *InvertIndexCompressedQuery) Order(o ...OrderFunc) *InvertIndexCompressedQuery {
	iicq.order = append(iicq.order, o...)
	return iicq
}

// QueryTerm chains the current query on the "term" edge.
func (iicq *InvertIndexCompressedQuery) QueryTerm() *TermQuery {
	query := &TermQuery{config: iicq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iicq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iicq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invertindexcompressed.Table, invertindexcompressed.FieldID, selector),
			sqlgraph.To(term.Table, term.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, invertindexcompressed.TermTable, invertindexcompressed.TermColumn),
		)
		fromU = sqlgraph.SetNeighbors(iicq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InvertIndexCompressed entity from the query.
// Returns a *NotFoundError when no InvertIndexCompressed was found.
func (iicq *InvertIndexCompressedQuery) First(ctx context.Context) (*InvertIndexCompressed, error) {
	nodes, err := iicq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invertindexcompressed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) FirstX(ctx context.Context) *InvertIndexCompressed {
	node, err := iicq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InvertIndexCompressed ID from the query.
// Returns a *NotFoundError when no InvertIndexCompressed ID was found.
func (iicq *InvertIndexCompressedQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iicq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invertindexcompressed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iicq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InvertIndexCompressed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InvertIndexCompressed entity is found.
// Returns a *NotFoundError when no InvertIndexCompressed entities are found.
func (iicq *InvertIndexCompressedQuery) Only(ctx context.Context) (*InvertIndexCompressed, error) {
	nodes, err := iicq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invertindexcompressed.Label}
	default:
		return nil, &NotSingularError{invertindexcompressed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) OnlyX(ctx context.Context) *InvertIndexCompressed {
	node, err := iicq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InvertIndexCompressed ID in the query.
// Returns a *NotSingularError when more than one InvertIndexCompressed ID is found.
// Returns a *NotFoundError when no entities are found.
func (iicq *InvertIndexCompressedQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iicq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = &NotSingularError{invertindexcompressed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iicq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InvertIndexCompresseds.
func (iicq *InvertIndexCompressedQuery) All(ctx context.Context) ([]*InvertIndexCompressed, error) {
	if err := iicq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iicq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) AllX(ctx context.Context) []*InvertIndexCompressed {
	nodes, err := iicq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InvertIndexCompressed IDs.
func (iicq *InvertIndexCompressedQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := iicq.Select(invertindexcompressed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iicq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iicq *InvertIndexCompressedQuery) Count(ctx context.Context) (int, error) {
	if err := iicq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iicq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) CountX(ctx context.Context) int {
	count, err := iicq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iicq *InvertIndexCompressedQuery) Exist(ctx context.Context) (bool, error) {
	if err := iicq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iicq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iicq *InvertIndexCompressedQuery) ExistX(ctx context.Context) bool {
	exist, err := iicq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvertIndexCompressedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iicq *InvertIndexCompressedQuery) Clone() *InvertIndexCompressedQuery {
	if iicq == nil {
		return nil
	}
	return &InvertIndexCompressedQuery{
		config:     iicq.config,
		limit:      iicq.limit,
		offset:     iicq.offset,
		order:      append([]OrderFunc{}, iicq.order...),
		predicates: append([]predicate.InvertIndexCompressed{}, iicq.predicates...),
		withTerm:   iicq.withTerm.Clone(),
		// clone intermediate query.
		sql:    iicq.sql.Clone(),
		path:   iicq.path,
		unique: iicq.unique,
	}
}

// WithTerm tells the query-builder to eager-load the nodes that are connected to
// the "term" edge. The optional arguments are used to configure the query builder of the edge.
func (iicq *InvertIndexCompressedQuery) WithTerm(opts ...func(*TermQuery)) *InvertIndexCompressedQuery {
	query := &TermQuery{config: iicq.config}
	for _, opt := range opts {
		opt(query)
	}
	iicq.withTerm = query
	return iicq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InvertIndexCompressed.Query().
//		GroupBy(invertindexcompressed.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iicq *InvertIndexCompressedQuery) GroupBy(field string, fields ...string) *InvertIndexCompressedGroupBy {
	group := &InvertIndexCompressedGroupBy{config: iicq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iicq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iicq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.InvertIndexCompressed.Query().
//		Select(invertindexcompressed.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (iicq *InvertIndexCompressedQuery) Select(fields ...string) *InvertIndexCompressedSelect {
	iicq.fields = append(iicq.fields, fields...)
	return &InvertIndexCompressedSelect{InvertIndexCompressedQuery: iicq}
}

func (iicq *InvertIndexCompressedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iicq.fields {
		if !invertindexcompressed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iicq.path != nil {
		prev, err := iicq.path(ctx)
		if err != nil {
			return err
		}
		iicq.sql = prev
	}
	return nil
}

func (iicq *InvertIndexCompressedQuery) sqlAll(ctx context.Context) ([]*InvertIndexCompressed, error) {
	var (
		nodes       = []*InvertIndexCompressed{}
		_spec       = iicq.querySpec()
		loadedTypes = [1]bool{
			iicq.withTerm != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &InvertIndexCompressed{config: iicq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, iicq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iicq.withTerm; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*InvertIndexCompressed)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Term(func(s *sql.Selector) {
			s.Where(sql.InValues(invertindexcompressed.TermColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.invert_index_compressed_term
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "invert_index_compressed_term" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "invert_index_compressed_term" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Term = n
		}
	}

	return nodes, nil
}

func (iicq *InvertIndexCompressedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iicq.querySpec()
	_spec.Node.Columns = iicq.fields
	if len(iicq.fields) > 0 {
		_spec.Unique = iicq.unique != nil && *iicq.unique
	}
	return sqlgraph.CountNodes(ctx, iicq.driver, _spec)
}

func (iicq *InvertIndexCompressedQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iicq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iicq *InvertIndexCompressedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invertindexcompressed.Table,
			Columns: invertindexcompressed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: invertindexcompressed.FieldID,
			},
		},
		From:   iicq.sql,
		Unique: true,
	}
	if unique := iicq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iicq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invertindexcompressed.FieldID)
		for i := range fields {
			if fields[i] != invertindexcompressed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iicq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iicq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iicq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iicq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iicq *InvertIndexCompressedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iicq.driver.Dialect())
	t1 := builder.Table(invertindexcompressed.Table)
	columns := iicq.fields
	if len(columns) == 0 {
		columns = invertindexcompressed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iicq.sql != nil {
		selector = iicq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iicq.unique != nil && *iicq.unique {
		selector.Distinct()
	}
	for _, p := range iicq.predicates {
		p(selector)
	}
	for _, p := range iicq.order {
		p(selector)
	}
	if offset := iicq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iicq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvertIndexCompressedGroupBy is the group-by builder for InvertIndexCompressed entities.
type InvertIndexCompressedGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iicgb *InvertIndexCompressedGroupBy) Aggregate(fns ...AggregateFunc) *InvertIndexCompressedGroupBy {
	iicgb.fns = append(iicgb.fns, fns...)
	return iicgb
}

// Scan applies the group-by query and scans the result into the given value.
func (iicgb *InvertIndexCompressedGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := iicgb.path(ctx)
	if err != nil {
		return err
	}
	iicgb.sql = query
	return iicgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := iicgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(iicgb.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := iicgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) StringsX(ctx context.Context) []string {
	v, err := iicgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = iicgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) StringX(ctx context.Context) string {
	v, err := iicgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(iicgb.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := iicgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) IntsX(ctx context.Context) []int {
	v, err := iicgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = iicgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) IntX(ctx context.Context) int {
	v, err := iicgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(iicgb.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := iicgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := iicgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = iicgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) Float64X(ctx context.Context) float64 {
	v, err := iicgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(iicgb.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := iicgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := iicgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (iicgb *InvertIndexCompressedGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = iicgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (iicgb *InvertIndexCompressedGroupBy) BoolX(ctx context.Context) bool {
	v, err := iicgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (iicgb *InvertIndexCompressedGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range iicgb.fields {
		if !invertindexcompressed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := iicgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iicgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (iicgb *InvertIndexCompressedGroupBy) sqlQuery() *sql.Selector {
	selector := iicgb.sql.Select()
	aggregation := make([]string, 0, len(iicgb.fns))
	for _, fn := range iicgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(iicgb.fields)+len(iicgb.fns))
		for _, f := range iicgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(iicgb.fields...)...)
}

// InvertIndexCompressedSelect is the builder for selecting fields of InvertIndexCompressed entities.
type InvertIndexCompressedSelect struct {
	*InvertIndexCompressedQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (iics *InvertIndexCompressedSelect) Scan(ctx context.Context, v interface{}) error {
	if err := iics.prepareQuery(ctx); err != nil {
		return err
	}
	iics.sql = iics.InvertIndexCompressedQuery.sqlQuery(ctx)
	return iics.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) ScanX(ctx context.Context, v interface{}) {
	if err := iics.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) Strings(ctx context.Context) ([]string, error) {
	if len(iics.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := iics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) StringsX(ctx context.Context) []string {
	v, err := iics.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = iics.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) StringX(ctx context.Context) string {
	v, err := iics.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) Ints(ctx context.Context) ([]int, error) {
	if len(iics.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := iics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) IntsX(ctx context.Context) []int {
	v, err := iics.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = iics.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) IntX(ctx context.Context) int {
	v, err := iics.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(iics.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := iics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) Float64sX(ctx context.Context) []float64 {
	v, err := iics.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = iics.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) Float64X(ctx context.Context) float64 {
	v, err := iics.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(iics.fields) > 1 {
		return nil, errors.New("ent: InvertIndexCompressedSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := iics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) BoolsX(ctx context.Context) []bool {
	v, err := iics.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (iics *InvertIndexCompressedSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = iics.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{invertindexcompressed.Label}
	default:
		err = fmt.Errorf("ent: InvertIndexCompressedSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (iics *InvertIndexCompressedSelect) BoolX(ctx context.Context) bool {
	v, err := iics.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (iics *InvertIndexCompressedSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := iics.sql.Query()
	if err := iics.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
