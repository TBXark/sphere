package docs

import (
	"github.com/tbxark/sphere/pkg/server/service/docs"
	"github.com/tbxark/sphere/swagger/api"
	"github.com/tbxark/sphere/swagger/dash"
)

type Targets struct {
	API  string `json:"api" yaml:"api"`
	Dash string `json:"dash" yaml:"dash"`
}

type Config struct {
	Address string  `json:"address" yaml:"address"`
	Targets Targets `json:"targets" yaml:"targets"`
}

type Web struct {
	*docs.Web
}

func NewWebServer(config *Config) *Web {
	return &Web{
		Web: docs.NewWebServer(&docs.Config{
			Address: config.Address,
			Targets: []docs.Target{
				{config.Targets.API, api.SwaggerInfoAPI},
				{config.Targets.Dash, dash.SwaggerInfoDash},
			},
		}),
	}
}
