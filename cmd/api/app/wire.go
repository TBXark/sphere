//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/tbxark/sphere/config"
	"github.com/tbxark/sphere/internal"
	"github.com/tbxark/sphere/pkg"
	"github.com/tbxark/sphere/pkg/utils/boot"
)

func NewAPIApplication(conf *config.Config) (*boot.Application, error) {
	wire.Build(config.ProviderSet, pkg.ProviderSet, internal.ProviderSet, wire.NewSet(CreateApplication))
	return &boot.Application{}, nil
}
