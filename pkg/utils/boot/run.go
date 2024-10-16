package boot

import (
	"flag"
	"fmt"
	"github.com/tbxark/go-base-api/configs"
	"os"
)

func DefaultCommandConfigFlagsParser() *configs.Config {
	path := flag.String("config", "config.json", "config file path")
	version := flag.Bool("version", false, "show version")
	help := flag.Bool("help", false, "show help")
	flag.Parse()

	if *version {
		fmt.Println(configs.BuildVersion)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	conf, err := LoadConfig(*path)
	if err != nil {
		fmt.Println("load config error: ", err)
		os.Exit(1)
	}
	return conf
}

func LoadConfig(path string) (*configs.Config, error) {
	conf, err := configs.LoadLocalConfig(path)
	if err != nil {
		return nil, err
	}
	if conf.Environments != nil {
		for k, v := range conf.Environments {
			e := os.Setenv(k, v)
			if e != nil {
				return nil, e
			}
		}
	}
	if conf.Remote == nil {
		return conf, nil
	}
	conf, err = configs.LoadRemoteConfig(conf.Remote)
	if err != nil {
		return nil, err
	}
	return conf, nil
}