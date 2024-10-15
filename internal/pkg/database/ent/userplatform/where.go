// Code generated by ent, DO NOT EDIT.

package userplatform

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tbxark/go-base-api/internal/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldUserID, v))
}

// Platform applies equality check predicate on the "platform" field. It's identical to PlatformEQ.
func Platform(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldPlatform, v))
}

// PlatformID applies equality check predicate on the "platform_id" field. It's identical to PlatformIDEQ.
func PlatformID(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldPlatformID, v))
}

// SecondID applies equality check predicate on the "second_id" field. It's identical to SecondIDEQ.
func SecondID(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldSecondID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotNull(FieldUpdatedAt))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLTE(FieldUserID, v))
}

// PlatformEQ applies the EQ predicate on the "platform" field.
func PlatformEQ(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldPlatform, v))
}

// PlatformNEQ applies the NEQ predicate on the "platform" field.
func PlatformNEQ(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNEQ(FieldPlatform, v))
}

// PlatformIn applies the In predicate on the "platform" field.
func PlatformIn(vs ...string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIn(FieldPlatform, vs...))
}

// PlatformNotIn applies the NotIn predicate on the "platform" field.
func PlatformNotIn(vs ...string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotIn(FieldPlatform, vs...))
}

// PlatformGT applies the GT predicate on the "platform" field.
func PlatformGT(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGT(FieldPlatform, v))
}

// PlatformGTE applies the GTE predicate on the "platform" field.
func PlatformGTE(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGTE(FieldPlatform, v))
}

// PlatformLT applies the LT predicate on the "platform" field.
func PlatformLT(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLT(FieldPlatform, v))
}

// PlatformLTE applies the LTE predicate on the "platform" field.
func PlatformLTE(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLTE(FieldPlatform, v))
}

// PlatformContains applies the Contains predicate on the "platform" field.
func PlatformContains(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldContains(FieldPlatform, v))
}

// PlatformHasPrefix applies the HasPrefix predicate on the "platform" field.
func PlatformHasPrefix(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldHasPrefix(FieldPlatform, v))
}

// PlatformHasSuffix applies the HasSuffix predicate on the "platform" field.
func PlatformHasSuffix(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldHasSuffix(FieldPlatform, v))
}

// PlatformEqualFold applies the EqualFold predicate on the "platform" field.
func PlatformEqualFold(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEqualFold(FieldPlatform, v))
}

// PlatformContainsFold applies the ContainsFold predicate on the "platform" field.
func PlatformContainsFold(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldContainsFold(FieldPlatform, v))
}

// PlatformIDEQ applies the EQ predicate on the "platform_id" field.
func PlatformIDEQ(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldPlatformID, v))
}

// PlatformIDNEQ applies the NEQ predicate on the "platform_id" field.
func PlatformIDNEQ(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNEQ(FieldPlatformID, v))
}

// PlatformIDIn applies the In predicate on the "platform_id" field.
func PlatformIDIn(vs ...string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIn(FieldPlatformID, vs...))
}

// PlatformIDNotIn applies the NotIn predicate on the "platform_id" field.
func PlatformIDNotIn(vs ...string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotIn(FieldPlatformID, vs...))
}

// PlatformIDGT applies the GT predicate on the "platform_id" field.
func PlatformIDGT(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGT(FieldPlatformID, v))
}

// PlatformIDGTE applies the GTE predicate on the "platform_id" field.
func PlatformIDGTE(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGTE(FieldPlatformID, v))
}

// PlatformIDLT applies the LT predicate on the "platform_id" field.
func PlatformIDLT(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLT(FieldPlatformID, v))
}

// PlatformIDLTE applies the LTE predicate on the "platform_id" field.
func PlatformIDLTE(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLTE(FieldPlatformID, v))
}

// PlatformIDContains applies the Contains predicate on the "platform_id" field.
func PlatformIDContains(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldContains(FieldPlatformID, v))
}

// PlatformIDHasPrefix applies the HasPrefix predicate on the "platform_id" field.
func PlatformIDHasPrefix(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldHasPrefix(FieldPlatformID, v))
}

// PlatformIDHasSuffix applies the HasSuffix predicate on the "platform_id" field.
func PlatformIDHasSuffix(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldHasSuffix(FieldPlatformID, v))
}

// PlatformIDEqualFold applies the EqualFold predicate on the "platform_id" field.
func PlatformIDEqualFold(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEqualFold(FieldPlatformID, v))
}

// PlatformIDContainsFold applies the ContainsFold predicate on the "platform_id" field.
func PlatformIDContainsFold(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldContainsFold(FieldPlatformID, v))
}

// SecondIDEQ applies the EQ predicate on the "second_id" field.
func SecondIDEQ(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEQ(FieldSecondID, v))
}

// SecondIDNEQ applies the NEQ predicate on the "second_id" field.
func SecondIDNEQ(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNEQ(FieldSecondID, v))
}

// SecondIDIn applies the In predicate on the "second_id" field.
func SecondIDIn(vs ...string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIn(FieldSecondID, vs...))
}

// SecondIDNotIn applies the NotIn predicate on the "second_id" field.
func SecondIDNotIn(vs ...string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotIn(FieldSecondID, vs...))
}

// SecondIDGT applies the GT predicate on the "second_id" field.
func SecondIDGT(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGT(FieldSecondID, v))
}

// SecondIDGTE applies the GTE predicate on the "second_id" field.
func SecondIDGTE(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldGTE(FieldSecondID, v))
}

// SecondIDLT applies the LT predicate on the "second_id" field.
func SecondIDLT(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLT(FieldSecondID, v))
}

// SecondIDLTE applies the LTE predicate on the "second_id" field.
func SecondIDLTE(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldLTE(FieldSecondID, v))
}

// SecondIDContains applies the Contains predicate on the "second_id" field.
func SecondIDContains(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldContains(FieldSecondID, v))
}

// SecondIDHasPrefix applies the HasPrefix predicate on the "second_id" field.
func SecondIDHasPrefix(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldHasPrefix(FieldSecondID, v))
}

// SecondIDHasSuffix applies the HasSuffix predicate on the "second_id" field.
func SecondIDHasSuffix(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldHasSuffix(FieldSecondID, v))
}

// SecondIDIsNil applies the IsNil predicate on the "second_id" field.
func SecondIDIsNil() predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldIsNull(FieldSecondID))
}

// SecondIDNotNil applies the NotNil predicate on the "second_id" field.
func SecondIDNotNil() predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldNotNull(FieldSecondID))
}

// SecondIDEqualFold applies the EqualFold predicate on the "second_id" field.
func SecondIDEqualFold(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldEqualFold(FieldSecondID, v))
}

// SecondIDContainsFold applies the ContainsFold predicate on the "second_id" field.
func SecondIDContainsFold(v string) predicate.UserPlatform {
	return predicate.UserPlatform(sql.FieldContainsFold(FieldSecondID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPlatform) predicate.UserPlatform {
	return predicate.UserPlatform(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPlatform) predicate.UserPlatform {
	return predicate.UserPlatform(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPlatform) predicate.UserPlatform {
	return predicate.UserPlatform(sql.NotPredicates(p))
}