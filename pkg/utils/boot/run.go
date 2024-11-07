package boot

import (
	"context"
	"flag"
	"fmt"
	"github.com/tbxark/sphere/pkg/log"
	"github.com/tbxark/sphere/pkg/log/logfields"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const DefaultTimezone = "Asia/Shanghai"

func init() {
	_ = InitTimezone(DefaultTimezone)
}

func InitTimezone(zone string) error {
	defaultLoc := "Asia/Shanghai"
	loc, err := time.LoadLocation(defaultLoc)
	if err != nil {
		return err
	}
	time.Local = loc
	return os.Setenv("TZ", defaultLoc)
}

func DefaultConfigParser[T any](ver string, parser func(string) (*T, error)) *T {
	path := flag.String("config", "config.json", "config file path")
	version := flag.Bool("version", false, "show version")
	help := flag.Bool("help", false, "show help")
	flag.Parse()

	if *version {
		fmt.Println(ver)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	conf, err := parser(*path)
	if err != nil {
		fmt.Println("load config error: ", err)
		os.Exit(1)
	}
	return conf
}

func Run[T any](ver string, conf *T, logConf *log.Options, builder func(*T) (*Application, error)) error {
	// Init logger
	log.Init(logConf, logfields.String("version", ver))
	log.Info("Start application", logfields.String("version", ver))
	defer func() {
		if e := log.Sync(); e != nil {
			log.Error("Failed to sync logger", logfields.Error(e))
		}
	}()

	// Create application
	app, err := builder(conf)
	if err != nil {
		return fmt.Errorf("failed to build application: %w", err)
	}

	// Listen for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	// Catch application error
	errChan := make(chan error, 1)
	go func() {
		if e := app.Start(context.Background()); e != nil {
			errChan <- e
		}
	}()

	// Wait for shutdown signal or application error
	select {
	case <-quit:
		log.Debug("Received shutdown signal")
	case e := <-errChan:
		log.Error("Application error", logfields.Error(e))
		return fmt.Errorf("application error: %w", e)
	}

	// Close application
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if e := app.Stop(ctx); e != nil {
		return fmt.Errorf("failed to close application: %w", e)
	}

	return nil
}
