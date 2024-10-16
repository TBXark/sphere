// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tbxark/sphere/internal/pkg/database/ent/admin"
)

// Admin is the model entity for the Admin schema.
type Admin struct {
	config `json:"-"`
	// ID of the ent.
	// 用户ID
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
	// 昵称
	Nickname string `json:"nickname,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty"`
	// 密码
	Password string `json:"-"`
	// 权限
	Roles        []string `json:"roles,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Admin) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case admin.FieldRoles:
			values[i] = new([]byte)
		case admin.FieldID, admin.FieldCreatedAt, admin.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case admin.FieldUsername, admin.FieldNickname, admin.FieldAvatar, admin.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Admin fields.
func (a *Admin) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case admin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case admin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Int64
			}
		case admin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Int64
			}
		case admin.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				a.Username = value.String
			}
		case admin.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				a.Nickname = value.String
			}
		case admin.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				a.Avatar = value.String
			}
		case admin.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = value.String
			}
		case admin.FieldRoles:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field roles", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Roles); err != nil {
					return fmt.Errorf("unmarshal field roles: %w", err)
				}
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Admin.
// This includes values selected through modifiers, order, etc.
func (a *Admin) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Admin.
// Note that you need to call Admin.Unwrap() before calling this method if this Admin
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Admin) Update() *AdminUpdateOne {
	return NewAdminClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Admin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Admin) Unwrap() *Admin {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Admin is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Admin) String() string {
	var builder strings.Builder
	builder.WriteString("Admin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(a.Username)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(a.Nickname)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(a.Avatar)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("roles=")
	builder.WriteString(fmt.Sprintf("%v", a.Roles))
	builder.WriteByte(')')
	return builder.String()
}

// Admins is a parsable slice of Admin.
type Admins []*Admin
