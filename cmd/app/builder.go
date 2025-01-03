package main

import (
	"github.com/TBXark/sphere/internal/biz/task"
	"github.com/TBXark/sphere/internal/server/api"
	"github.com/TBXark/sphere/internal/server/bot"
	"github.com/TBXark/sphere/internal/server/dash"
	"github.com/TBXark/sphere/internal/server/docs"
	"github.com/TBXark/sphere/pkg/utils/boot"
)

func newApplication(dash *dash.Web, api *api.Web, docs *docs.Web, initialize *task.DashInitialize, cleaner *task.ConnectCleaner) *boot.Application {
	return boot.NewApplication(
		dash,
		api,
		docs,
		initialize,
		cleaner,
	)
}

func newAPIApplication(api *api.Web, initialize *task.DashInitialize, cleaner *task.ConnectCleaner) *boot.Application {
	return boot.NewApplication(
		api,
		initialize,
		cleaner,
	)
}

func newDashApplication(dash *dash.Web, initialize *task.DashInitialize, cleaner *task.ConnectCleaner) *boot.Application {
	return boot.NewApplication(
		dash,
		initialize,
		cleaner,
	)
}

func newBotApplication(app *bot.Bot) *boot.Application {
	return boot.NewApplication(app)
}
