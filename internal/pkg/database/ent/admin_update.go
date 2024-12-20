// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/sphere/internal/pkg/database/ent/admin"
	"github.com/TBXark/sphere/internal/pkg/database/ent/predicate"
)

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks     []Hook
	mutation  *AdminMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AdminUpdate builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AdminUpdate) SetUpdatedAt(i int64) *AdminUpdate {
	au.mutation.ResetUpdatedAt()
	au.mutation.SetUpdatedAt(i)
	return au
}

// AddUpdatedAt adds i to the "updated_at" field.
func (au *AdminUpdate) AddUpdatedAt(i int64) *AdminUpdate {
	au.mutation.AddUpdatedAt(i)
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *AdminUpdate) ClearUpdatedAt() *AdminUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetUsername sets the "username" field.
func (au *AdminUpdate) SetUsername(s string) *AdminUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (au *AdminUpdate) SetNillableUsername(s *string) *AdminUpdate {
	if s != nil {
		au.SetUsername(*s)
	}
	return au
}

// SetNickname sets the "nickname" field.
func (au *AdminUpdate) SetNickname(s string) *AdminUpdate {
	au.mutation.SetNickname(s)
	return au
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (au *AdminUpdate) SetNillableNickname(s *string) *AdminUpdate {
	if s != nil {
		au.SetNickname(*s)
	}
	return au
}

// ClearNickname clears the value of the "nickname" field.
func (au *AdminUpdate) ClearNickname() *AdminUpdate {
	au.mutation.ClearNickname()
	return au
}

// SetAvatar sets the "avatar" field.
func (au *AdminUpdate) SetAvatar(s string) *AdminUpdate {
	au.mutation.SetAvatar(s)
	return au
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (au *AdminUpdate) SetNillableAvatar(s *string) *AdminUpdate {
	if s != nil {
		au.SetAvatar(*s)
	}
	return au
}

// ClearAvatar clears the value of the "avatar" field.
func (au *AdminUpdate) ClearAvatar() *AdminUpdate {
	au.mutation.ClearAvatar()
	return au
}

// SetPassword sets the "password" field.
func (au *AdminUpdate) SetPassword(s string) *AdminUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AdminUpdate) SetNillablePassword(s *string) *AdminUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// SetRoles sets the "roles" field.
func (au *AdminUpdate) SetRoles(s []string) *AdminUpdate {
	au.mutation.SetRoles(s)
	return au
}

// AppendRoles appends s to the "roles" field.
func (au *AdminUpdate) AppendRoles(s []string) *AdminUpdate {
	au.mutation.AppendRoles(s)
	return au
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AdminUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok && !au.mutation.UpdatedAtCleared() {
		v := admin.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AdminUpdate) check() error {
	if v, ok := au.mutation.Username(); ok {
		if err := admin.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Admin.username": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AdminUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdminUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(admin.Table, admin.Columns, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if au.mutation.CreatedAtCleared() {
		_spec.ClearField(admin.FieldCreatedAt, field.TypeInt64)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(admin.FieldUpdatedAt, field.TypeInt64, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(admin.FieldUpdatedAt, field.TypeInt64)
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
	}
	if value, ok := au.mutation.Nickname(); ok {
		_spec.SetField(admin.FieldNickname, field.TypeString, value)
	}
	if au.mutation.NicknameCleared() {
		_spec.ClearField(admin.FieldNickname, field.TypeString)
	}
	if value, ok := au.mutation.Avatar(); ok {
		_spec.SetField(admin.FieldAvatar, field.TypeString, value)
	}
	if au.mutation.AvatarCleared() {
		_spec.ClearField(admin.FieldAvatar, field.TypeString)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
	}
	if value, ok := au.mutation.Roles(); ok {
		_spec.SetField(admin.FieldRoles, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedRoles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, admin.FieldRoles, value)
		})
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AdminMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AdminUpdateOne) SetUpdatedAt(i int64) *AdminUpdateOne {
	auo.mutation.ResetUpdatedAt()
	auo.mutation.SetUpdatedAt(i)
	return auo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (auo *AdminUpdateOne) AddUpdatedAt(i int64) *AdminUpdateOne {
	auo.mutation.AddUpdatedAt(i)
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *AdminUpdateOne) ClearUpdatedAt() *AdminUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetUsername sets the "username" field.
func (auo *AdminUpdateOne) SetUsername(s string) *AdminUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableUsername(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetUsername(*s)
	}
	return auo
}

// SetNickname sets the "nickname" field.
func (auo *AdminUpdateOne) SetNickname(s string) *AdminUpdateOne {
	auo.mutation.SetNickname(s)
	return auo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableNickname(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetNickname(*s)
	}
	return auo
}

// ClearNickname clears the value of the "nickname" field.
func (auo *AdminUpdateOne) ClearNickname() *AdminUpdateOne {
	auo.mutation.ClearNickname()
	return auo
}

// SetAvatar sets the "avatar" field.
func (auo *AdminUpdateOne) SetAvatar(s string) *AdminUpdateOne {
	auo.mutation.SetAvatar(s)
	return auo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableAvatar(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetAvatar(*s)
	}
	return auo
}

// ClearAvatar clears the value of the "avatar" field.
func (auo *AdminUpdateOne) ClearAvatar() *AdminUpdateOne {
	auo.mutation.ClearAvatar()
	return auo
}

// SetPassword sets the "password" field.
func (auo *AdminUpdateOne) SetPassword(s string) *AdminUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillablePassword(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// SetRoles sets the "roles" field.
func (auo *AdminUpdateOne) SetRoles(s []string) *AdminUpdateOne {
	auo.mutation.SetRoles(s)
	return auo
}

// AppendRoles appends s to the "roles" field.
func (auo *AdminUpdateOne) AppendRoles(s []string) *AdminUpdateOne {
	auo.mutation.AppendRoles(s)
	return auo
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// Where appends a list predicates to the AdminUpdate builder.
func (auo *AdminUpdateOne) Where(ps ...predicate.Admin) *AdminUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdminUpdateOne) Select(field string, fields ...string) *AdminUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Admin entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AdminUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok && !auo.mutation.UpdatedAtCleared() {
		v := admin.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AdminUpdateOne) check() error {
	if v, ok := auo.mutation.Username(); ok {
		if err := admin.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Admin.username": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AdminUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdminUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (_node *Admin, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(admin.Table, admin.Columns, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Admin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, admin.FieldID)
		for _, f := range fields {
			if !admin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != admin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auo.mutation.CreatedAtCleared() {
		_spec.ClearField(admin.FieldCreatedAt, field.TypeInt64)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(admin.FieldUpdatedAt, field.TypeInt64, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(admin.FieldUpdatedAt, field.TypeInt64)
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
	}
	if value, ok := auo.mutation.Nickname(); ok {
		_spec.SetField(admin.FieldNickname, field.TypeString, value)
	}
	if auo.mutation.NicknameCleared() {
		_spec.ClearField(admin.FieldNickname, field.TypeString)
	}
	if value, ok := auo.mutation.Avatar(); ok {
		_spec.SetField(admin.FieldAvatar, field.TypeString, value)
	}
	if auo.mutation.AvatarCleared() {
		_spec.ClearField(admin.FieldAvatar, field.TypeString)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
	}
	if value, ok := auo.mutation.Roles(); ok {
		_spec.SetField(admin.FieldRoles, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedRoles(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, admin.FieldRoles, value)
		})
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Admin{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
