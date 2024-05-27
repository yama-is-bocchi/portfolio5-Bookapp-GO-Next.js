// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Bookapp/ent/lock"
	"Bookapp/ent/predicate"
	"Bookapp/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LockUpdate is the builder for updating Lock entities.
type LockUpdate struct {
	config
	hooks    []Hook
	mutation *LockMutation
}

// Where appends a list predicates to the LockUpdate builder.
func (lu *LockUpdate) Where(ps ...predicate.Lock) *LockUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUserID sets the "user_id" field.
func (lu *LockUpdate) SetUserID(i int) *LockUpdate {
	lu.mutation.SetUserID(i)
	return lu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (lu *LockUpdate) SetNillableUserID(i *int) *LockUpdate {
	if i != nil {
		lu.SetUserID(*i)
	}
	return lu
}

// SetDate sets the "date" field.
func (lu *LockUpdate) SetDate(t time.Time) *LockUpdate {
	lu.mutation.SetDate(t)
	return lu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (lu *LockUpdate) SetNillableDate(t *time.Time) *LockUpdate {
	if t != nil {
		lu.SetDate(*t)
	}
	return lu
}

// SetUser sets the "user" edge to the User entity.
func (lu *LockUpdate) SetUser(u *User) *LockUpdate {
	return lu.SetUserID(u.ID)
}

// Mutation returns the LockMutation object of the builder.
func (lu *LockUpdate) Mutation() *LockMutation {
	return lu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lu *LockUpdate) ClearUser() *LockUpdate {
	lu.mutation.ClearUser()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LockUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LockUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LockUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LockUpdate) check() error {
	if _, ok := lu.mutation.UserID(); lu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Lock.user"`)
	}
	return nil
}

func (lu *LockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lock.Table, lock.Columns, sqlgraph.NewFieldSpec(lock.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Date(); ok {
		_spec.SetField(lock.FieldDate, field.TypeTime, value)
	}
	if lu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lock.UserTable,
			Columns: []string{lock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lock.UserTable,
			Columns: []string{lock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LockUpdateOne is the builder for updating a single Lock entity.
type LockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LockMutation
}

// SetUserID sets the "user_id" field.
func (luo *LockUpdateOne) SetUserID(i int) *LockUpdateOne {
	luo.mutation.SetUserID(i)
	return luo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (luo *LockUpdateOne) SetNillableUserID(i *int) *LockUpdateOne {
	if i != nil {
		luo.SetUserID(*i)
	}
	return luo
}

// SetDate sets the "date" field.
func (luo *LockUpdateOne) SetDate(t time.Time) *LockUpdateOne {
	luo.mutation.SetDate(t)
	return luo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (luo *LockUpdateOne) SetNillableDate(t *time.Time) *LockUpdateOne {
	if t != nil {
		luo.SetDate(*t)
	}
	return luo
}

// SetUser sets the "user" edge to the User entity.
func (luo *LockUpdateOne) SetUser(u *User) *LockUpdateOne {
	return luo.SetUserID(u.ID)
}

// Mutation returns the LockMutation object of the builder.
func (luo *LockUpdateOne) Mutation() *LockMutation {
	return luo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (luo *LockUpdateOne) ClearUser() *LockUpdateOne {
	luo.mutation.ClearUser()
	return luo
}

// Where appends a list predicates to the LockUpdate builder.
func (luo *LockUpdateOne) Where(ps ...predicate.Lock) *LockUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LockUpdateOne) Select(field string, fields ...string) *LockUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Lock entity.
func (luo *LockUpdateOne) Save(ctx context.Context) (*Lock, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LockUpdateOne) SaveX(ctx context.Context) *Lock {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LockUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LockUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LockUpdateOne) check() error {
	if _, ok := luo.mutation.UserID(); luo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Lock.user"`)
	}
	return nil
}

func (luo *LockUpdateOne) sqlSave(ctx context.Context) (_node *Lock, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lock.Table, lock.Columns, sqlgraph.NewFieldSpec(lock.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Lock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lock.FieldID)
		for _, f := range fields {
			if !lock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Date(); ok {
		_spec.SetField(lock.FieldDate, field.TypeTime, value)
	}
	if luo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lock.UserTable,
			Columns: []string{lock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lock.UserTable,
			Columns: []string{lock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Lock{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
