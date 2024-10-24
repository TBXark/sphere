package authorizer

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/constraints"
	"time"
)

type UID interface {
	constraints.Integer | string
}

var (
	ErrorExpiredToken = errors.New("expired token")
)

const (
	ContextKeyUID     = "uid"
	ContextKeySubject = "subject"
	ContextKeyRoles   = "roles"
)

type RBACClaims[T UID] struct {
	jwt.RegisteredClaims
	UID   T        `json:"uid,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

func NewRBACClaims[T UID](uid T, subject string, roles []string, expiresAt time.Time) *RBACClaims[T] {
	return &RBACClaims[T]{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UID:   uid,
		Roles: roles,
	}
}