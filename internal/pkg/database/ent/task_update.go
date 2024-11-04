// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tbxark/sphere/internal/pkg/database/ent/predicate"
	"github.com/tbxark/sphere/internal/pkg/database/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks     []Hook
	mutation  *TaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TaskUpdate) SetUpdatedAt(i int64) *TaskUpdate {
	tu.mutation.ResetUpdatedAt()
	tu.mutation.SetUpdatedAt(i)
	return tu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tu *TaskUpdate) AddUpdatedAt(i int64) *TaskUpdate {
	tu.mutation.AddUpdatedAt(i)
	return tu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tu *TaskUpdate) ClearUpdatedAt() *TaskUpdate {
	tu.mutation.ClearUpdatedAt()
	return tu
}

// SetName sets the "name" field.
func (tu *TaskUpdate) SetName(s string) *TaskUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableName(s *string) *TaskUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(t task.Status) *TaskUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatus(t *task.Status) *TaskUpdate {
	if t != nil {
		tu.SetStatus(*t)
	}
	return tu
}

// SetResult sets the "result" field.
func (tu *TaskUpdate) SetResult(s string) *TaskUpdate {
	tu.mutation.SetResult(s)
	return tu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableResult(s *string) *TaskUpdate {
	if s != nil {
		tu.SetResult(*s)
	}
	return tu
}

// ClearResult clears the value of the "result" field.
func (tu *TaskUpdate) ClearResult() *TaskUpdate {
	tu.mutation.ClearResult()
	return tu
}

// SetError sets the "error" field.
func (tu *TaskUpdate) SetError(s string) *TaskUpdate {
	tu.mutation.SetError(s)
	return tu
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableError(s *string) *TaskUpdate {
	if s != nil {
		tu.SetError(*s)
	}
	return tu
}

// ClearError clears the value of the "error" field.
func (tu *TaskUpdate) ClearError() *TaskUpdate {
	tu.mutation.ClearError()
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok && !tu.mutation.UpdatedAtCleared() {
		v := task.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Result(); ok {
		if err := task.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "Task.result": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Error(); ok {
		if err := task.ErrorValidator(v); err != nil {
			return &ValidationError{Name: "error", err: fmt.Errorf(`ent: validator failed for field "Task.error": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TaskUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.CreatedAtCleared() {
		_spec.ClearField(task.FieldCreatedAt, field.TypeInt64)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(task.FieldUpdatedAt, field.TypeInt64, value)
	}
	if tu.mutation.UpdatedAtCleared() {
		_spec.ClearField(task.FieldUpdatedAt, field.TypeInt64)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Result(); ok {
		_spec.SetField(task.FieldResult, field.TypeString, value)
	}
	if tu.mutation.ResultCleared() {
		_spec.ClearField(task.FieldResult, field.TypeString)
	}
	if value, ok := tu.mutation.Error(); ok {
		_spec.SetField(task.FieldError, field.TypeString, value)
	}
	if tu.mutation.ErrorCleared() {
		_spec.ClearField(task.FieldError, field.TypeString)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TaskUpdateOne) SetUpdatedAt(i int64) *TaskUpdateOne {
	tuo.mutation.ResetUpdatedAt()
	tuo.mutation.SetUpdatedAt(i)
	return tuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tuo *TaskUpdateOne) AddUpdatedAt(i int64) *TaskUpdateOne {
	tuo.mutation.AddUpdatedAt(i)
	return tuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tuo *TaskUpdateOne) ClearUpdatedAt() *TaskUpdateOne {
	tuo.mutation.ClearUpdatedAt()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TaskUpdateOne) SetName(s string) *TaskUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableName(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(t task.Status) *TaskUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatus(t *task.Status) *TaskUpdateOne {
	if t != nil {
		tuo.SetStatus(*t)
	}
	return tuo
}

// SetResult sets the "result" field.
func (tuo *TaskUpdateOne) SetResult(s string) *TaskUpdateOne {
	tuo.mutation.SetResult(s)
	return tuo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableResult(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetResult(*s)
	}
	return tuo
}

// ClearResult clears the value of the "result" field.
func (tuo *TaskUpdateOne) ClearResult() *TaskUpdateOne {
	tuo.mutation.ClearResult()
	return tuo
}

// SetError sets the "error" field.
func (tuo *TaskUpdateOne) SetError(s string) *TaskUpdateOne {
	tuo.mutation.SetError(s)
	return tuo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableError(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetError(*s)
	}
	return tuo
}

// ClearError clears the value of the "error" field.
func (tuo *TaskUpdateOne) ClearError() *TaskUpdateOne {
	tuo.mutation.ClearError()
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok && !tuo.mutation.UpdatedAtCleared() {
		v := task.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Result(); ok {
		if err := task.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "Task.result": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Error(); ok {
		if err := task.ErrorValidator(v); err != nil {
			return &ValidationError{Name: "error", err: fmt.Errorf(`ent: validator failed for field "Task.error": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TaskUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tuo.mutation.CreatedAtCleared() {
		_spec.ClearField(task.FieldCreatedAt, field.TypeInt64)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(task.FieldUpdatedAt, field.TypeInt64, value)
	}
	if tuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(task.FieldUpdatedAt, field.TypeInt64)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Result(); ok {
		_spec.SetField(task.FieldResult, field.TypeString, value)
	}
	if tuo.mutation.ResultCleared() {
		_spec.ClearField(task.FieldResult, field.TypeString)
	}
	if value, ok := tuo.mutation.Error(); ok {
		_spec.SetField(task.FieldError, field.TypeString, value)
	}
	if tuo.mutation.ErrorCleared() {
		_spec.ClearField(task.FieldError, field.TypeString)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
