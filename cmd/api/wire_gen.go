// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/tbxark/sphere/internal/biz/task"
	"github.com/tbxark/sphere/internal/config"
	"github.com/tbxark/sphere/internal/pkg/dao"
	"github.com/tbxark/sphere/internal/pkg/database/client"
	api2 "github.com/tbxark/sphere/internal/server/api"
	"github.com/tbxark/sphere/internal/service/api"
	"github.com/tbxark/sphere/pkg/cache/memory"
	"github.com/tbxark/sphere/pkg/storage/qiniu"
	"github.com/tbxark/sphere/pkg/utils/boot"
	"github.com/tbxark/sphere/pkg/wechat"
)

// Injectors from wire.go:

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
	qiniuQiniu := qiniu.NewQiniu(qiniuConfig)
	cache := memory.NewByteCache()
	service := api.NewService(daoDao, wechatWechat, qiniuQiniu, cache)
	web := api2.NewWebServer(apiConfig, service)
	dashInitialize := task.NewDashInitialize(daoDao)
	connectCleaner := task.NewConnectCleaner(entClient)
	application := newApplication(web, dashInitialize, connectCleaner)
	return application, nil
}
