// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/TBXark/sphere/internal/biz/task"
	"github.com/TBXark/sphere/internal/config"
	"github.com/TBXark/sphere/internal/pkg/dao"
	"github.com/TBXark/sphere/internal/pkg/database/client"
	dash2 "github.com/TBXark/sphere/internal/server/dash"
	"github.com/TBXark/sphere/internal/service/dash"
	"github.com/TBXark/sphere/pkg/cache/memory"
	"github.com/TBXark/sphere/pkg/storage/qiniu"
	"github.com/TBXark/sphere/pkg/utils/boot"
	"github.com/TBXark/sphere/pkg/wechat"
)

// Injectors from wire.go:

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
	qiniuQiniu := qiniu.NewQiniu(qiniuConfig)
	cache := memory.NewByteCache()
	service := dash.NewService(daoDao, wechatWechat, qiniuQiniu, cache)
	web := dash2.NewWebServer(dashConfig, service)
	connectCleaner := task.NewConnectCleaner(entClient)
	dashInitialize := task.NewDashInitialize(daoDao)
	application := newApplication(web, connectCleaner, dashInitialize)
	return application, nil
}
