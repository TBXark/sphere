// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/tbxark/sphere/internal/config"
	bot2 "github.com/tbxark/sphere/internal/server/bot"
	"github.com/tbxark/sphere/internal/service/bot"
	"github.com/tbxark/sphere/pkg/utils/boot"
)

// Injectors from wire.go:

func NewBotApplication(conf *config.Config) (*boot.Application, error) {
	botConfig := conf.Bot
	service := bot.NewService()
	botBot := bot2.NewApp(botConfig, service)
	application := newApplication(botBot)
	return application, nil
}
