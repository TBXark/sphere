package tma_auth

import (
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"strconv"
	"time"
)

const AuthorizationPrefixTMA = "tma"

type TmaAuth struct {
	token string
	expIn time.Duration
}

func NewTmaAuth(token string) *TmaAuth {
	return &TmaAuth{
		token: token,
		expIn: time.Hour * 24,
	}
}

func (t *TmaAuth) Validate(token string) (map[string]any, error) {
	err := initdata.Validate(token, t.token, t.expIn)
	if err != nil {
		return nil, err
	}
	initData, err := initdata.Parse(token)
	if err != nil {
		return nil, err
	}

	res := make(map[string]any, 4)
	res["uid"] = strconv.Itoa(int(initData.Chat.ID))
	res["username"] = initData.Chat.Username
	res["roles"] = string(initData.ChatType)
	res["exp"] = initData.AuthDate().Add(t.expIn)

	return res, nil
}

func (t *TmaAuth) ParseRolesString(roles string) map[string]struct{} {
	return make(map[string]struct{})
}
