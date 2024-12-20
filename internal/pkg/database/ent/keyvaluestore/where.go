// Code generated by ent, DO NOT EDIT.

package keyvaluestore

import (
	"entgo.io/ent/dialect/sql"
	"github.com/TBXark/sphere/internal/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldUpdatedAt, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v []byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldValue, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotNull(FieldUpdatedAt))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v []byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v []byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...[]byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...[]byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v []byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v []byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v []byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v []byte) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldLTE(FieldValue, v))
}

// ValueIsNil applies the IsNil predicate on the "value" field.
func ValueIsNil() predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldIsNull(FieldValue))
}

// ValueNotNil applies the NotNil predicate on the "value" field.
func ValueNotNil() predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.FieldNotNull(FieldValue))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KeyValueStore) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KeyValueStore) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.KeyValueStore) predicate.KeyValueStore {
	return predicate.KeyValueStore(sql.NotPredicates(p))
}
