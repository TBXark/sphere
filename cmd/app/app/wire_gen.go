// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/tbxark/go-base-api/config"
	"github.com/tbxark/go-base-api/internal/biz/api"
	"github.com/tbxark/go-base-api/internal/biz/dash"
	"github.com/tbxark/go-base-api/internal/biz/task"
	"github.com/tbxark/go-base-api/internal/pkg/boot"
	"github.com/tbxark/go-base-api/internal/pkg/dao"
	"github.com/tbxark/go-base-api/pkg/cache/memory"
	"github.com/tbxark/go-base-api/pkg/dao/client"
	"github.com/tbxark/go-base-api/pkg/storage/qiniu"
	"github.com/tbxark/go-base-api/pkg/wechat"
)

// Injectors from wire.go:

func NewApplication(conf *config.Config) (*boot.Application, error) {
	dashConfig := conf.Dash
	clientConfig := conf.Database
	entClient, err := client.NewDbClient(clientConfig)
	if err != nil {
		return nil, err
	}
	daoDao := dao.NewDao(entClient)
	wechatConfig := conf.WxMini
	wechatWechat := wechat.NewWechat(wechatConfig)
	qiniuConfig := conf.Storage
	qiniuQiniu := qiniu.NewQiniu(qiniuConfig)
	cache := memory.NewByteCache()
	web := dash.NewWebServer(dashConfig, daoDao, wechatWechat, qiniuQiniu, cache)
	apiConfig := conf.API
	apiWeb := api.NewWebServer(apiConfig, daoDao, wechatWechat, qiniuQiniu, cache)
	dashInitialize := task.NewInitialize(daoDao)
	connectCleaner := task.NewCleaner(entClient)
	application := CreateApplication(web, apiWeb, dashInitialize, connectCleaner)
	return application, nil
}
