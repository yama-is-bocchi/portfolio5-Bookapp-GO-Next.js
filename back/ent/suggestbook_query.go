// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Bookapp/ent/predicate"
	"Bookapp/ent/suggestbook"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SuggestBookQuery is the builder for querying SuggestBook entities.
type SuggestBookQuery struct {
	config
	ctx        *QueryContext
	order      []suggestbook.OrderOption
	inters     []Interceptor
	predicates []predicate.SuggestBook
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SuggestBookQuery builder.
func (sbq *SuggestBookQuery) Where(ps ...predicate.SuggestBook) *SuggestBookQuery {
	sbq.predicates = append(sbq.predicates, ps...)
	return sbq
}

// Limit the number of records to be returned by this query.
func (sbq *SuggestBookQuery) Limit(limit int) *SuggestBookQuery {
	sbq.ctx.Limit = &limit
	return sbq
}

// Offset to start from.
func (sbq *SuggestBookQuery) Offset(offset int) *SuggestBookQuery {
	sbq.ctx.Offset = &offset
	return sbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sbq *SuggestBookQuery) Unique(unique bool) *SuggestBookQuery {
	sbq.ctx.Unique = &unique
	return sbq
}

// Order specifies how the records should be ordered.
func (sbq *SuggestBookQuery) Order(o ...suggestbook.OrderOption) *SuggestBookQuery {
	sbq.order = append(sbq.order, o...)
	return sbq
}

// First returns the first SuggestBook entity from the query.
// Returns a *NotFoundError when no SuggestBook was found.
func (sbq *SuggestBookQuery) First(ctx context.Context) (*SuggestBook, error) {
	nodes, err := sbq.Limit(1).All(setContextOp(ctx, sbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{suggestbook.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sbq *SuggestBookQuery) FirstX(ctx context.Context) *SuggestBook {
	node, err := sbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SuggestBook ID from the query.
// Returns a *NotFoundError when no SuggestBook ID was found.
func (sbq *SuggestBookQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sbq.Limit(1).IDs(setContextOp(ctx, sbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{suggestbook.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sbq *SuggestBookQuery) FirstIDX(ctx context.Context) int {
	id, err := sbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SuggestBook entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SuggestBook entity is found.
// Returns a *NotFoundError when no SuggestBook entities are found.
func (sbq *SuggestBookQuery) Only(ctx context.Context) (*SuggestBook, error) {
	nodes, err := sbq.Limit(2).All(setContextOp(ctx, sbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{suggestbook.Label}
	default:
		return nil, &NotSingularError{suggestbook.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sbq *SuggestBookQuery) OnlyX(ctx context.Context) *SuggestBook {
	node, err := sbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SuggestBook ID in the query.
// Returns a *NotSingularError when more than one SuggestBook ID is found.
// Returns a *NotFoundError when no entities are found.
func (sbq *SuggestBookQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sbq.Limit(2).IDs(setContextOp(ctx, sbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{suggestbook.Label}
	default:
		err = &NotSingularError{suggestbook.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sbq *SuggestBookQuery) OnlyIDX(ctx context.Context) int {
	id, err := sbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SuggestBooks.
func (sbq *SuggestBookQuery) All(ctx context.Context) ([]*SuggestBook, error) {
	ctx = setContextOp(ctx, sbq.ctx, "All")
	if err := sbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SuggestBook, *SuggestBookQuery]()
	return withInterceptors[[]*SuggestBook](ctx, sbq, qr, sbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sbq *SuggestBookQuery) AllX(ctx context.Context) []*SuggestBook {
	nodes, err := sbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SuggestBook IDs.
func (sbq *SuggestBookQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sbq.ctx.Unique == nil && sbq.path != nil {
		sbq.Unique(true)
	}
	ctx = setContextOp(ctx, sbq.ctx, "IDs")
	if err = sbq.Select(suggestbook.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sbq *SuggestBookQuery) IDsX(ctx context.Context) []int {
	ids, err := sbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sbq *SuggestBookQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sbq.ctx, "Count")
	if err := sbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sbq, querierCount[*SuggestBookQuery](), sbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sbq *SuggestBookQuery) CountX(ctx context.Context) int {
	count, err := sbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sbq *SuggestBookQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sbq.ctx, "Exist")
	switch _, err := sbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sbq *SuggestBookQuery) ExistX(ctx context.Context) bool {
	exist, err := sbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SuggestBookQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sbq *SuggestBookQuery) Clone() *SuggestBookQuery {
	if sbq == nil {
		return nil
	}
	return &SuggestBookQuery{
		config:     sbq.config,
		ctx:        sbq.ctx.Clone(),
		order:      append([]suggestbook.OrderOption{}, sbq.order...),
		inters:     append([]Interceptor{}, sbq.inters...),
		predicates: append([]predicate.SuggestBook{}, sbq.predicates...),
		// clone intermediate query.
		sql:  sbq.sql.Clone(),
		path: sbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SuggestBook.Query().
//		GroupBy(suggestbook.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sbq *SuggestBookQuery) GroupBy(field string, fields ...string) *SuggestBookGroupBy {
	sbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SuggestBookGroupBy{build: sbq}
	grbuild.flds = &sbq.ctx.Fields
	grbuild.label = suggestbook.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.SuggestBook.Query().
//		Select(suggestbook.FieldTitle).
//		Scan(ctx, &v)
func (sbq *SuggestBookQuery) Select(fields ...string) *SuggestBookSelect {
	sbq.ctx.Fields = append(sbq.ctx.Fields, fields...)
	sbuild := &SuggestBookSelect{SuggestBookQuery: sbq}
	sbuild.label = suggestbook.Label
	sbuild.flds, sbuild.scan = &sbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SuggestBookSelect configured with the given aggregations.
func (sbq *SuggestBookQuery) Aggregate(fns ...AggregateFunc) *SuggestBookSelect {
	return sbq.Select().Aggregate(fns...)
}

func (sbq *SuggestBookQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sbq); err != nil {
				return err
			}
		}
	}
	for _, f := range sbq.ctx.Fields {
		if !suggestbook.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sbq.path != nil {
		prev, err := sbq.path(ctx)
		if err != nil {
			return err
		}
		sbq.sql = prev
	}
	return nil
}

func (sbq *SuggestBookQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SuggestBook, error) {
	var (
		nodes = []*SuggestBook{}
		_spec = sbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SuggestBook).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SuggestBook{config: sbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sbq *SuggestBookQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sbq.querySpec()
	_spec.Node.Columns = sbq.ctx.Fields
	if len(sbq.ctx.Fields) > 0 {
		_spec.Unique = sbq.ctx.Unique != nil && *sbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sbq.driver, _spec)
}

func (sbq *SuggestBookQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(suggestbook.Table, suggestbook.Columns, sqlgraph.NewFieldSpec(suggestbook.FieldID, field.TypeInt))
	_spec.From = sbq.sql
	if unique := sbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sbq.path != nil {
		_spec.Unique = true
	}
	if fields := sbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, suggestbook.FieldID)
		for i := range fields {
			if fields[i] != suggestbook.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sbq *SuggestBookQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sbq.driver.Dialect())
	t1 := builder.Table(suggestbook.Table)
	columns := sbq.ctx.Fields
	if len(columns) == 0 {
		columns = suggestbook.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sbq.sql != nil {
		selector = sbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sbq.ctx.Unique != nil && *sbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sbq.predicates {
		p(selector)
	}
	for _, p := range sbq.order {
		p(selector)
	}
	if offset := sbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SuggestBookGroupBy is the group-by builder for SuggestBook entities.
type SuggestBookGroupBy struct {
	selector
	build *SuggestBookQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sbgb *SuggestBookGroupBy) Aggregate(fns ...AggregateFunc) *SuggestBookGroupBy {
	sbgb.fns = append(sbgb.fns, fns...)
	return sbgb
}

// Scan applies the selector query and scans the result into the given value.
func (sbgb *SuggestBookGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sbgb.build.ctx, "GroupBy")
	if err := sbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SuggestBookQuery, *SuggestBookGroupBy](ctx, sbgb.build, sbgb, sbgb.build.inters, v)
}

func (sbgb *SuggestBookGroupBy) sqlScan(ctx context.Context, root *SuggestBookQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sbgb.fns))
	for _, fn := range sbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sbgb.flds)+len(sbgb.fns))
		for _, f := range *sbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SuggestBookSelect is the builder for selecting fields of SuggestBook entities.
type SuggestBookSelect struct {
	*SuggestBookQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sbs *SuggestBookSelect) Aggregate(fns ...AggregateFunc) *SuggestBookSelect {
	sbs.fns = append(sbs.fns, fns...)
	return sbs
}

// Scan applies the selector query and scans the result into the given value.
func (sbs *SuggestBookSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sbs.ctx, "Select")
	if err := sbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SuggestBookQuery, *SuggestBookSelect](ctx, sbs.SuggestBookQuery, sbs, sbs.inters, v)
}

func (sbs *SuggestBookSelect) sqlScan(ctx context.Context, root *SuggestBookQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sbs.fns))
	for _, fn := range sbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
