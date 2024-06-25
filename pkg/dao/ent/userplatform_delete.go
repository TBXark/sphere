// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tbxark/go-base-api/pkg/dao/ent/predicate"
	"github.com/tbxark/go-base-api/pkg/dao/ent/userplatform"
)

// UserPlatformDelete is the builder for deleting a UserPlatform entity.
type UserPlatformDelete struct {
	config
	hooks    []Hook
	mutation *UserPlatformMutation
}

// Where appends a list predicates to the UserPlatformDelete builder.
func (upd *UserPlatformDelete) Where(ps ...predicate.UserPlatform) *UserPlatformDelete {
	upd.mutation.Where(ps...)
	return upd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upd *UserPlatformDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, upd.sqlExec, upd.mutation, upd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (upd *UserPlatformDelete) ExecX(ctx context.Context) int {
	n, err := upd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upd *UserPlatformDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userplatform.Table, sqlgraph.NewFieldSpec(userplatform.FieldID, field.TypeInt))
	if ps := upd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, upd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	upd.mutation.done = true
	return affected, err
}

// UserPlatformDeleteOne is the builder for deleting a single UserPlatform entity.
type UserPlatformDeleteOne struct {
	upd *UserPlatformDelete
}

// Where appends a list predicates to the UserPlatformDelete builder.
func (updo *UserPlatformDeleteOne) Where(ps ...predicate.UserPlatform) *UserPlatformDeleteOne {
	updo.upd.mutation.Where(ps...)
	return updo
}

// Exec executes the deletion query.
func (updo *UserPlatformDeleteOne) Exec(ctx context.Context) error {
	n, err := updo.upd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userplatform.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (updo *UserPlatformDeleteOne) ExecX(ctx context.Context) {
	if err := updo.Exec(ctx); err != nil {
		panic(err)
	}
}
