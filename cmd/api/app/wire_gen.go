// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/tbxark/go-base-api/config"
	"github.com/tbxark/go-base-api/internal/biz/dash"
	"github.com/tbxark/go-base-api/internal/biz/task"
	"github.com/tbxark/go-base-api/internal/pkg/boot"
	"github.com/tbxark/go-base-api/internal/pkg/dao"
	"github.com/tbxark/go-base-api/pkg/cache/memory"
	"github.com/tbxark/go-base-api/pkg/cdn/qiniu"
	"github.com/tbxark/go-base-api/pkg/dao/client"
	"github.com/tbxark/go-base-api/pkg/wechat"
)

// Injectors from wire.go:

func NewAPIApplication(cfg *config.Config) (*boot.Application, error) {
	dashConfig := cfg.Dash
	clientConfig := cfg.Database
	entClient, err := client.NewDbClient(clientConfig)
	if err != nil {
		return nil, err
	}
	daoDao := dao.NewDao(entClient)
	wechatConfig := cfg.WxMini
	wechatWechat := wechat.NewWechat(wechatConfig)
	qiniuConfig := cfg.CDN
	qiniuQiniu := qiniu.NewQiniu(qiniuConfig)
	cache := memory.NewByteCache()
	web := dash.NewWebServer(dashConfig, daoDao, wechatWechat, qiniuQiniu, cache)
	dashInitialize := task.NewInitialize(daoDao)
	connectCleaner := task.NewCleaner(entClient)
	application := CreateApplication(web, dashInitialize, connectCleaner)
	return application, nil
}
