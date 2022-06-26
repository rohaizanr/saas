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
	"github.com/go-saas/saas/examples/ent/shared/ent/predicate"
	"github.com/go-saas/saas/examples/ent/shared/ent/tenantconn"
)

// TenantConnUpdate is the builder for updating TenantConn entities.
type TenantConnUpdate struct {
	config
	hooks    []Hook
	mutation *TenantConnMutation
}

// Where appends a list predicates to the TenantConnUpdate builder.
func (tcu *TenantConnUpdate) Where(ps ...predicate.TenantConn) *TenantConnUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetUpdateTime sets the "update_time" field.
func (tcu *TenantConnUpdate) SetUpdateTime(t time.Time) *TenantConnUpdate {
	tcu.mutation.SetUpdateTime(t)
	return tcu
}

// SetKey sets the "key" field.
func (tcu *TenantConnUpdate) SetKey(s string) *TenantConnUpdate {
	tcu.mutation.SetKey(s)
	return tcu
}

// SetValue sets the "value" field.
func (tcu *TenantConnUpdate) SetValue(s string) *TenantConnUpdate {
	tcu.mutation.SetValue(s)
	return tcu
}

// Mutation returns the TenantConnMutation object of the builder.
func (tcu *TenantConnUpdate) Mutation() *TenantConnMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TenantConnUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tcu.defaults()
	if len(tcu.hooks) == 0 {
		affected, err = tcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TenantConnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcu.mutation = mutation
			affected, err = tcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcu.hooks) - 1; i >= 0; i-- {
			if tcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TenantConnUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TenantConnUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TenantConnUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TenantConnUpdate) defaults() {
	if _, ok := tcu.mutation.UpdateTime(); !ok {
		v := tenantconn.UpdateDefaultUpdateTime()
		tcu.mutation.SetUpdateTime(v)
	}
}

func (tcu *TenantConnUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tenantconn.Table,
			Columns: tenantconn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tenantconn.FieldID,
			},
		},
	}
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tenantconn.FieldUpdateTime,
		})
	}
	if value, ok := tcu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenantconn.FieldKey,
		})
	}
	if value, ok := tcu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenantconn.FieldValue,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tenantconn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TenantConnUpdateOne is the builder for updating a single TenantConn entity.
type TenantConnUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TenantConnMutation
}

// SetUpdateTime sets the "update_time" field.
func (tcuo *TenantConnUpdateOne) SetUpdateTime(t time.Time) *TenantConnUpdateOne {
	tcuo.mutation.SetUpdateTime(t)
	return tcuo
}

// SetKey sets the "key" field.
func (tcuo *TenantConnUpdateOne) SetKey(s string) *TenantConnUpdateOne {
	tcuo.mutation.SetKey(s)
	return tcuo
}

// SetValue sets the "value" field.
func (tcuo *TenantConnUpdateOne) SetValue(s string) *TenantConnUpdateOne {
	tcuo.mutation.SetValue(s)
	return tcuo
}

// Mutation returns the TenantConnMutation object of the builder.
func (tcuo *TenantConnUpdateOne) Mutation() *TenantConnMutation {
	return tcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TenantConnUpdateOne) Select(field string, fields ...string) *TenantConnUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TenantConn entity.
func (tcuo *TenantConnUpdateOne) Save(ctx context.Context) (*TenantConn, error) {
	var (
		err  error
		node *TenantConn
	)
	tcuo.defaults()
	if len(tcuo.hooks) == 0 {
		node, err = tcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TenantConnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcuo.mutation = mutation
			node, err = tcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tcuo.hooks) - 1; i >= 0; i-- {
			if tcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TenantConnUpdateOne) SaveX(ctx context.Context) *TenantConn {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TenantConnUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TenantConnUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TenantConnUpdateOne) defaults() {
	if _, ok := tcuo.mutation.UpdateTime(); !ok {
		v := tenantconn.UpdateDefaultUpdateTime()
		tcuo.mutation.SetUpdateTime(v)
	}
}

func (tcuo *TenantConnUpdateOne) sqlSave(ctx context.Context) (_node *TenantConn, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tenantconn.Table,
			Columns: tenantconn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tenantconn.FieldID,
			},
		},
	}
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TenantConn.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tenantconn.FieldID)
		for _, f := range fields {
			if !tenantconn.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tenantconn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tenantconn.FieldUpdateTime,
		})
	}
	if value, ok := tcuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenantconn.FieldKey,
		})
	}
	if value, ok := tcuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenantconn.FieldValue,
		})
	}
	_node = &TenantConn{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tenantconn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
