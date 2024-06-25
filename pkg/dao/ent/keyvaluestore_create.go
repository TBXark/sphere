// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tbxark/go-base-api/pkg/dao/ent/keyvaluestore"
)

// KeyValueStoreCreate is the builder for creating a KeyValueStore entity.
type KeyValueStoreCreate struct {
	config
	mutation *KeyValueStoreMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (kvsc *KeyValueStoreCreate) SetCreatedAt(i int64) *KeyValueStoreCreate {
	kvsc.mutation.SetCreatedAt(i)
	return kvsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kvsc *KeyValueStoreCreate) SetNillableCreatedAt(i *int64) *KeyValueStoreCreate {
	if i != nil {
		kvsc.SetCreatedAt(*i)
	}
	return kvsc
}

// SetUpdatedAt sets the "updated_at" field.
func (kvsc *KeyValueStoreCreate) SetUpdatedAt(i int64) *KeyValueStoreCreate {
	kvsc.mutation.SetUpdatedAt(i)
	return kvsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kvsc *KeyValueStoreCreate) SetNillableUpdatedAt(i *int64) *KeyValueStoreCreate {
	if i != nil {
		kvsc.SetUpdatedAt(*i)
	}
	return kvsc
}

// SetKey sets the "key" field.
func (kvsc *KeyValueStoreCreate) SetKey(s string) *KeyValueStoreCreate {
	kvsc.mutation.SetKey(s)
	return kvsc
}

// SetValue sets the "value" field.
func (kvsc *KeyValueStoreCreate) SetValue(b []byte) *KeyValueStoreCreate {
	kvsc.mutation.SetValue(b)
	return kvsc
}

// Mutation returns the KeyValueStoreMutation object of the builder.
func (kvsc *KeyValueStoreCreate) Mutation() *KeyValueStoreMutation {
	return kvsc.mutation
}

// Save creates the KeyValueStore in the database.
func (kvsc *KeyValueStoreCreate) Save(ctx context.Context) (*KeyValueStore, error) {
	kvsc.defaults()
	return withHooks(ctx, kvsc.sqlSave, kvsc.mutation, kvsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kvsc *KeyValueStoreCreate) SaveX(ctx context.Context) *KeyValueStore {
	v, err := kvsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kvsc *KeyValueStoreCreate) Exec(ctx context.Context) error {
	_, err := kvsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kvsc *KeyValueStoreCreate) ExecX(ctx context.Context) {
	if err := kvsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kvsc *KeyValueStoreCreate) defaults() {
	if _, ok := kvsc.mutation.CreatedAt(); !ok {
		v := keyvaluestore.DefaultCreatedAt()
		kvsc.mutation.SetCreatedAt(v)
	}
	if _, ok := kvsc.mutation.UpdatedAt(); !ok {
		v := keyvaluestore.DefaultUpdatedAt()
		kvsc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kvsc *KeyValueStoreCreate) check() error {
	if _, ok := kvsc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "KeyValueStore.key"`)}
	}
	return nil
}

func (kvsc *KeyValueStoreCreate) sqlSave(ctx context.Context) (*KeyValueStore, error) {
	if err := kvsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kvsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kvsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	kvsc.mutation.id = &_node.ID
	kvsc.mutation.done = true
	return _node, nil
}

func (kvsc *KeyValueStoreCreate) createSpec() (*KeyValueStore, *sqlgraph.CreateSpec) {
	var (
		_node = &KeyValueStore{config: kvsc.config}
		_spec = sqlgraph.NewCreateSpec(keyvaluestore.Table, sqlgraph.NewFieldSpec(keyvaluestore.FieldID, field.TypeInt))
	)
	_spec.OnConflict = kvsc.conflict
	if value, ok := kvsc.mutation.CreatedAt(); ok {
		_spec.SetField(keyvaluestore.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := kvsc.mutation.UpdatedAt(); ok {
		_spec.SetField(keyvaluestore.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := kvsc.mutation.Key(); ok {
		_spec.SetField(keyvaluestore.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := kvsc.mutation.Value(); ok {
		_spec.SetField(keyvaluestore.FieldValue, field.TypeBytes, value)
		_node.Value = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.KeyValueStore.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KeyValueStoreUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (kvsc *KeyValueStoreCreate) OnConflict(opts ...sql.ConflictOption) *KeyValueStoreUpsertOne {
	kvsc.conflict = opts
	return &KeyValueStoreUpsertOne{
		create: kvsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.KeyValueStore.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (kvsc *KeyValueStoreCreate) OnConflictColumns(columns ...string) *KeyValueStoreUpsertOne {
	kvsc.conflict = append(kvsc.conflict, sql.ConflictColumns(columns...))
	return &KeyValueStoreUpsertOne{
		create: kvsc,
	}
}

type (
	// KeyValueStoreUpsertOne is the builder for "upsert"-ing
	//  one KeyValueStore node.
	KeyValueStoreUpsertOne struct {
		create *KeyValueStoreCreate
	}

	// KeyValueStoreUpsert is the "OnConflict" setter.
	KeyValueStoreUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *KeyValueStoreUpsert) SetUpdatedAt(v int64) *KeyValueStoreUpsert {
	u.Set(keyvaluestore.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *KeyValueStoreUpsert) UpdateUpdatedAt() *KeyValueStoreUpsert {
	u.SetExcluded(keyvaluestore.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *KeyValueStoreUpsert) AddUpdatedAt(v int64) *KeyValueStoreUpsert {
	u.Add(keyvaluestore.FieldUpdatedAt, v)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *KeyValueStoreUpsert) ClearUpdatedAt() *KeyValueStoreUpsert {
	u.SetNull(keyvaluestore.FieldUpdatedAt)
	return u
}

// SetKey sets the "key" field.
func (u *KeyValueStoreUpsert) SetKey(v string) *KeyValueStoreUpsert {
	u.Set(keyvaluestore.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *KeyValueStoreUpsert) UpdateKey() *KeyValueStoreUpsert {
	u.SetExcluded(keyvaluestore.FieldKey)
	return u
}

// SetValue sets the "value" field.
func (u *KeyValueStoreUpsert) SetValue(v []byte) *KeyValueStoreUpsert {
	u.Set(keyvaluestore.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *KeyValueStoreUpsert) UpdateValue() *KeyValueStoreUpsert {
	u.SetExcluded(keyvaluestore.FieldValue)
	return u
}

// ClearValue clears the value of the "value" field.
func (u *KeyValueStoreUpsert) ClearValue() *KeyValueStoreUpsert {
	u.SetNull(keyvaluestore.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.KeyValueStore.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *KeyValueStoreUpsertOne) UpdateNewValues() *KeyValueStoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(keyvaluestore.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.KeyValueStore.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *KeyValueStoreUpsertOne) Ignore() *KeyValueStoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KeyValueStoreUpsertOne) DoNothing() *KeyValueStoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KeyValueStoreCreate.OnConflict
// documentation for more info.
func (u *KeyValueStoreUpsertOne) Update(set func(*KeyValueStoreUpsert)) *KeyValueStoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KeyValueStoreUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *KeyValueStoreUpsertOne) SetUpdatedAt(v int64) *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *KeyValueStoreUpsertOne) AddUpdatedAt(v int64) *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *KeyValueStoreUpsertOne) UpdateUpdatedAt() *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *KeyValueStoreUpsertOne) ClearUpdatedAt() *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetKey sets the "key" field.
func (u *KeyValueStoreUpsertOne) SetKey(v string) *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *KeyValueStoreUpsertOne) UpdateKey() *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *KeyValueStoreUpsertOne) SetValue(v []byte) *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *KeyValueStoreUpsertOne) UpdateValue() *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.UpdateValue()
	})
}

// ClearValue clears the value of the "value" field.
func (u *KeyValueStoreUpsertOne) ClearValue() *KeyValueStoreUpsertOne {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.ClearValue()
	})
}

// Exec executes the query.
func (u *KeyValueStoreUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for KeyValueStoreCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KeyValueStoreUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *KeyValueStoreUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *KeyValueStoreUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// KeyValueStoreCreateBulk is the builder for creating many KeyValueStore entities in bulk.
type KeyValueStoreCreateBulk struct {
	config
	err      error
	builders []*KeyValueStoreCreate
	conflict []sql.ConflictOption
}

// Save creates the KeyValueStore entities in the database.
func (kvscb *KeyValueStoreCreateBulk) Save(ctx context.Context) ([]*KeyValueStore, error) {
	if kvscb.err != nil {
		return nil, kvscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kvscb.builders))
	nodes := make([]*KeyValueStore, len(kvscb.builders))
	mutators := make([]Mutator, len(kvscb.builders))
	for i := range kvscb.builders {
		func(i int, root context.Context) {
			builder := kvscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KeyValueStoreMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kvscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = kvscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kvscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, kvscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kvscb *KeyValueStoreCreateBulk) SaveX(ctx context.Context) []*KeyValueStore {
	v, err := kvscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kvscb *KeyValueStoreCreateBulk) Exec(ctx context.Context) error {
	_, err := kvscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kvscb *KeyValueStoreCreateBulk) ExecX(ctx context.Context) {
	if err := kvscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.KeyValueStore.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KeyValueStoreUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (kvscb *KeyValueStoreCreateBulk) OnConflict(opts ...sql.ConflictOption) *KeyValueStoreUpsertBulk {
	kvscb.conflict = opts
	return &KeyValueStoreUpsertBulk{
		create: kvscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.KeyValueStore.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (kvscb *KeyValueStoreCreateBulk) OnConflictColumns(columns ...string) *KeyValueStoreUpsertBulk {
	kvscb.conflict = append(kvscb.conflict, sql.ConflictColumns(columns...))
	return &KeyValueStoreUpsertBulk{
		create: kvscb,
	}
}

// KeyValueStoreUpsertBulk is the builder for "upsert"-ing
// a bulk of KeyValueStore nodes.
type KeyValueStoreUpsertBulk struct {
	create *KeyValueStoreCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.KeyValueStore.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *KeyValueStoreUpsertBulk) UpdateNewValues() *KeyValueStoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(keyvaluestore.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.KeyValueStore.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *KeyValueStoreUpsertBulk) Ignore() *KeyValueStoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KeyValueStoreUpsertBulk) DoNothing() *KeyValueStoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KeyValueStoreCreateBulk.OnConflict
// documentation for more info.
func (u *KeyValueStoreUpsertBulk) Update(set func(*KeyValueStoreUpsert)) *KeyValueStoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KeyValueStoreUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *KeyValueStoreUpsertBulk) SetUpdatedAt(v int64) *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *KeyValueStoreUpsertBulk) AddUpdatedAt(v int64) *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *KeyValueStoreUpsertBulk) UpdateUpdatedAt() *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *KeyValueStoreUpsertBulk) ClearUpdatedAt() *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetKey sets the "key" field.
func (u *KeyValueStoreUpsertBulk) SetKey(v string) *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *KeyValueStoreUpsertBulk) UpdateKey() *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.UpdateKey()
	})
}

// SetValue sets the "value" field.
func (u *KeyValueStoreUpsertBulk) SetValue(v []byte) *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *KeyValueStoreUpsertBulk) UpdateValue() *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.UpdateValue()
	})
}

// ClearValue clears the value of the "value" field.
func (u *KeyValueStoreUpsertBulk) ClearValue() *KeyValueStoreUpsertBulk {
	return u.Update(func(s *KeyValueStoreUpsert) {
		s.ClearValue()
	})
}

// Exec executes the query.
func (u *KeyValueStoreUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the KeyValueStoreCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for KeyValueStoreCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KeyValueStoreUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
