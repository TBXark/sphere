package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/TBXark/sphere/utils/idgenerator"
)

type Admin struct {
	ent.Schema
}

func (Admin) Fields() []ent.Field {
	times := DefaultTimeFields()
	return []ent.Field{
		field.Int64("id").Unique().Immutable().DefaultFunc(idgenerator.NextId).Comment("用户ID"),
		field.String("username").Unique().MinLen(1).Comment("用户名"),
		field.String("nickname").Optional().Default("").Comment("昵称"),
		field.String("avatar").Optional().Default("").Comment("头像"),
		field.String("password").Comment("密码").Sensitive(),
		field.Strings("roles").Default([]string{}).Comment("权限").Sensitive(),
		times[0], times[1],
	}
}
