// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/TBXark/sphere/cache/memory"
	"github.com/TBXark/sphere/layout/internal/biz/task/conncleaner"
	"github.com/TBXark/sphere/layout/internal/biz/task/dashinit"
	"github.com/TBXark/sphere/layout/internal/config"
	"github.com/TBXark/sphere/layout/internal/pkg/dao"
	"github.com/TBXark/sphere/layout/internal/pkg/database/client"
	api2 "github.com/TBXark/sphere/layout/internal/server/api"
	bot2 "github.com/TBXark/sphere/layout/internal/server/bot"
	dash2 "github.com/TBXark/sphere/layout/internal/server/dash"
	"github.com/TBXark/sphere/layout/internal/server/docs"
	"github.com/TBXark/sphere/layout/internal/service/api"
	"github.com/TBXark/sphere/layout/internal/service/bot"
	"github.com/TBXark/sphere/layout/internal/service/dash"
	"github.com/TBXark/sphere/storage/qiniu"
	"github.com/TBXark/sphere/utils/boot"
	"github.com/TBXark/sphere/wechat"
)

// Injectors from wire.go:

func NewApplication(conf *config.Config) (*boot.Application, error) {
	dashConfig := conf.Dash
	clientConfig := conf.Database
	entClient, err := client.NewDataBaseClient(clientConfig)
	if err != nil {
		return nil, err
	}
	daoDao := dao.NewDao(entClient)
	wechatConfig := conf.WxMini
	wechatWechat := wechat.NewWechat(wechatConfig)
	qiniuConfig := conf.Storage
	qiniuClient := qiniu.NewClient(qiniuConfig)
	cache := memory.NewByteCache()
	service := dash.NewService(daoDao, wechatWechat, qiniuClient, cache)
	web := dash2.NewWebServer(dashConfig, service)
	apiConfig := conf.API
	apiService := api.NewService(daoDao, wechatWechat, qiniuClient, cache)
	apiWeb := api2.NewWebServer(apiConfig, apiService)
	docsConfig := conf.Docs
	docsWeb := docs.NewWebServer(docsConfig)
	dashInitialize := dashinit.NewDashInitialize(daoDao)
	connectCleaner := conncleaner.NewConnectCleaner(entClient)
	application := newApplication(web, apiWeb, docsWeb, dashInitialize, connectCleaner)
	return application, nil
}

func NewAPIApplication(conf *config.Config) (*boot.Application, error) {
	apiConfig := conf.API
	clientConfig := conf.Database
	entClient, err := client.NewDataBaseClient(clientConfig)
	if err != nil {
		return nil, err
	}
	daoDao := dao.NewDao(entClient)
	wechatConfig := conf.WxMini
	wechatWechat := wechat.NewWechat(wechatConfig)
	qiniuConfig := conf.Storage
	qiniuClient := qiniu.NewClient(qiniuConfig)
	cache := memory.NewByteCache()
	service := api.NewService(daoDao, wechatWechat, qiniuClient, cache)
	web := api2.NewWebServer(apiConfig, service)
	dashInitialize := dashinit.NewDashInitialize(daoDao)
	connectCleaner := conncleaner.NewConnectCleaner(entClient)
	application := newAPIApplication(web, dashInitialize, connectCleaner)
	return application, nil
}

func NewDashApplication(conf *config.Config) (*boot.Application, error) {
	dashConfig := conf.Dash
	clientConfig := conf.Database
	entClient, err := client.NewDataBaseClient(clientConfig)
	if err != nil {
		return nil, err
	}
	daoDao := dao.NewDao(entClient)
	wechatConfig := conf.WxMini
	wechatWechat := wechat.NewWechat(wechatConfig)
	qiniuConfig := conf.Storage
	qiniuClient := qiniu.NewClient(qiniuConfig)
	cache := memory.NewByteCache()
	service := dash.NewService(daoDao, wechatWechat, qiniuClient, cache)
	web := dash2.NewWebServer(dashConfig, service)
	dashInitialize := dashinit.NewDashInitialize(daoDao)
	connectCleaner := conncleaner.NewConnectCleaner(entClient)
	application := newDashApplication(web, dashInitialize, connectCleaner)
	return application, nil
}

func NewBotApplication(conf *config.Config) (*boot.Application, error) {
	botConfig := conf.Bot
	service := bot.NewService()
	botBot, err := bot2.NewApp(botConfig, service)
	if err != nil {
		return nil, err
	}
	application := newBotApplication(botBot)
	return application, nil
}
