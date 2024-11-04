// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/tbxark/sphere/internal/pkg/database/ent/admin"
	"github.com/tbxark/sphere/internal/pkg/database/ent/keyvaluestore"
	"github.com/tbxark/sphere/internal/pkg/database/ent/schema"
	"github.com/tbxark/sphere/internal/pkg/database/ent/task"
	"github.com/tbxark/sphere/internal/pkg/database/ent/user"
	"github.com/tbxark/sphere/internal/pkg/database/ent/userplatform"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminMixin := schema.Admin{}.Mixin()
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminMixinFields0[0].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() int64)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminMixinFields0[1].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() int64)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() int64)
	// adminDescUsername is the schema descriptor for username field.
	adminDescUsername := adminFields[1].Descriptor()
	// admin.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	admin.UsernameValidator = adminDescUsername.Validators[0].(func(string) error)
	// adminDescNickname is the schema descriptor for nickname field.
	adminDescNickname := adminFields[2].Descriptor()
	// admin.DefaultNickname holds the default value on creation for the nickname field.
	admin.DefaultNickname = adminDescNickname.Default.(string)
	// adminDescAvatar is the schema descriptor for avatar field.
	adminDescAvatar := adminFields[3].Descriptor()
	// admin.DefaultAvatar holds the default value on creation for the avatar field.
	admin.DefaultAvatar = adminDescAvatar.Default.(string)
	// adminDescRoles is the schema descriptor for roles field.
	adminDescRoles := adminFields[5].Descriptor()
	// admin.DefaultRoles holds the default value on creation for the roles field.
	admin.DefaultRoles = adminDescRoles.Default.([]string)
	// adminDescID is the schema descriptor for id field.
	adminDescID := adminFields[0].Descriptor()
	// admin.DefaultID holds the default value on creation for the id field.
	admin.DefaultID = adminDescID.Default.(func() int64)
	keyvaluestoreMixin := schema.KeyValueStore{}.Mixin()
	keyvaluestoreMixinFields0 := keyvaluestoreMixin[0].Fields()
	_ = keyvaluestoreMixinFields0
	keyvaluestoreFields := schema.KeyValueStore{}.Fields()
	_ = keyvaluestoreFields
	// keyvaluestoreDescCreatedAt is the schema descriptor for created_at field.
	keyvaluestoreDescCreatedAt := keyvaluestoreMixinFields0[0].Descriptor()
	// keyvaluestore.DefaultCreatedAt holds the default value on creation for the created_at field.
	keyvaluestore.DefaultCreatedAt = keyvaluestoreDescCreatedAt.Default.(func() int64)
	// keyvaluestoreDescUpdatedAt is the schema descriptor for updated_at field.
	keyvaluestoreDescUpdatedAt := keyvaluestoreMixinFields0[1].Descriptor()
	// keyvaluestore.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	keyvaluestore.DefaultUpdatedAt = keyvaluestoreDescUpdatedAt.Default.(func() int64)
	// keyvaluestore.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	keyvaluestore.UpdateDefaultUpdatedAt = keyvaluestoreDescUpdatedAt.UpdateDefault.(func() int64)
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskMixinFields0[0].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() int64)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskMixinFields0[1].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() int64)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() int64)
	// taskDescName is the schema descriptor for name field.
	taskDescName := taskFields[1].Descriptor()
	// task.DefaultName holds the default value on creation for the name field.
	task.DefaultName = taskDescName.Default.(string)
	// taskDescResult is the schema descriptor for result field.
	taskDescResult := taskFields[3].Descriptor()
	// task.DefaultResult holds the default value on creation for the result field.
	task.DefaultResult = taskDescResult.Default.(string)
	// task.ResultValidator is a validator for the "result" field. It is called by the builders before save.
	task.ResultValidator = taskDescResult.Validators[0].(func(string) error)
	// taskDescError is the schema descriptor for error field.
	taskDescError := taskFields[4].Descriptor()
	// task.DefaultError holds the default value on creation for the error field.
	task.DefaultError = taskDescError.Default.(string)
	// task.ErrorValidator is a validator for the "error" field. It is called by the builders before save.
	task.ErrorValidator = taskDescError.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() int64)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() int64)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() int64)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescRemark is the schema descriptor for remark field.
	userDescRemark := userFields[2].Descriptor()
	// user.DefaultRemark holds the default value on creation for the remark field.
	user.DefaultRemark = userDescRemark.Default.(string)
	// user.RemarkValidator is a validator for the "remark" field. It is called by the builders before save.
	user.RemarkValidator = userDescRemark.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[3].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[4].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescFlags is the schema descriptor for flags field.
	userDescFlags := userFields[5].Descriptor()
	// user.DefaultFlags holds the default value on creation for the flags field.
	user.DefaultFlags = userDescFlags.Default.(uint64)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() int64)
	userplatformMixin := schema.UserPlatform{}.Mixin()
	userplatformMixinFields0 := userplatformMixin[0].Fields()
	_ = userplatformMixinFields0
	userplatformFields := schema.UserPlatform{}.Fields()
	_ = userplatformFields
	// userplatformDescCreatedAt is the schema descriptor for created_at field.
	userplatformDescCreatedAt := userplatformMixinFields0[0].Descriptor()
	// userplatform.DefaultCreatedAt holds the default value on creation for the created_at field.
	userplatform.DefaultCreatedAt = userplatformDescCreatedAt.Default.(func() int64)
	// userplatformDescUpdatedAt is the schema descriptor for updated_at field.
	userplatformDescUpdatedAt := userplatformMixinFields0[1].Descriptor()
	// userplatform.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userplatform.DefaultUpdatedAt = userplatformDescUpdatedAt.Default.(func() int64)
	// userplatform.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userplatform.UpdateDefaultUpdatedAt = userplatformDescUpdatedAt.UpdateDefault.(func() int64)
	// userplatformDescSecondID is the schema descriptor for second_id field.
	userplatformDescSecondID := userplatformFields[3].Descriptor()
	// userplatform.DefaultSecondID holds the default value on creation for the second_id field.
	userplatform.DefaultSecondID = userplatformDescSecondID.Default.(string)
}
