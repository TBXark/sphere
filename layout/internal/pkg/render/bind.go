package render

import (
	"github.com/TBXark/sphere/database/bind"
	dashv1 "github.com/TBXark/sphere/layout/api/dash/v1"
	"github.com/TBXark/sphere/layout/internal/pkg/database/ent"
)

func CreateAdmin(source *ent.AdminCreate, target *dashv1.AdminEdit, options ...bind.Options) *ent.AdminCreate {
	option := bind.NewGenOptions(options...)
	if !option.IgnoreSetZero("username") || !(target.Username == "") {
		source.SetUsername(target.Username)
	}
	if !option.IgnoreSetZero("nickname") || !(target.Nickname == "") {
		source.SetNickname(target.Nickname)
	}
	if !option.IgnoreSetZero("avatar") || !(target.Avatar == "") {
		source.SetAvatar(target.Avatar)
	}
	if !option.IgnoreSetZero("password") || !(target.Password == "") {
		source.SetPassword(target.Password)
	}
	if !option.IgnoreSetZero("roles") || !(len(target.Roles) == 0) {
		source.SetRoles(target.Roles)
	}
	return source
}

func UpdateOneAdmin(source *ent.AdminUpdateOne, target *dashv1.AdminEdit, options ...bind.Options) *ent.AdminUpdateOne {
	option := bind.NewGenOptions(options...)
	if !option.IgnoreSetZero("username") || !(target.Username == "") {
		source.SetUsername(target.Username)
	}
	if !option.IgnoreSetZero("nickname") || !(target.Nickname == "") {
		source.SetNickname(target.Nickname)
	}
	if !option.IgnoreSetZero("avatar") || !(target.Avatar == "") {
		source.SetAvatar(target.Avatar)
	}
	if !option.IgnoreSetZero("password") || !(target.Password == "") {
		source.SetPassword(target.Password)
	}
	if !option.IgnoreSetZero("roles") || !(len(target.Roles) == 0) {
		source.SetRoles(target.Roles)
	}
	return source
}
