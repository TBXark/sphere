// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/tbxark/go-base-api/config"
	"github.com/tbxark/go-base-api/internal/biz/bot"
	"github.com/tbxark/go-base-api/internal/pkg/boot"
)

// Injectors from wire.go:

func NewBotApplication(conf *config.Config) (*boot.Application, error) {
	botConfig := conf.Bot
	app := bot.NewApp(botConfig)
	application := CreateApplication(app)
	return application, nil
}
