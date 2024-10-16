package boot

import (
	"github.com/tbxark/go-base-api/configs"
	"github.com/tbxark/go-base-api/pkg/log"
	"github.com/tbxark/go-base-api/pkg/log/logfields"
	"os"
	"sync"
	"time"
)

type Task interface {
	Identifier() string
	Run() error
}

type Cleaner interface {
	Clean() error
}

type Application struct {
	task    []Task
	cleaner []Cleaner
}

func NewApplication(tasks []Task, cleaners []Cleaner) *Application {
	return &Application{
		task:    tasks,
		cleaner: cleaners,
	}
}

func (a *Application) Run() {
	wg := sync.WaitGroup{}
	for _, t := range a.task {
		wg.Add(1)

		log.Infof("task %s start", t.Identifier())
		go func(t Task) {
			defer func() {
				if r := recover(); r != nil {
					log.Errorw(
						"task panic",
						logfields.String("task", t.Identifier()),
					)
					wg.Done()
				}
			}()
			defer wg.Done()
			if err := t.Run(); err != nil {
				log.Errorw(
					"task error",
					logfields.String("task", t.Identifier()),
					logfields.Error(err),
				)
			}
		}(t)
	}
	wg.Wait()
}

func (a *Application) Clean() {
	for _, c := range a.cleaner {
		_ = c.Clean()
	}
}

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

func Run(conf *configs.Config, builder func(*configs.Config) (*Application, error)) error {
	log.Init(conf.Log, logfields.String("version", configs.BuildVersion))
	log.Info("Start application", logfields.String("version", configs.BuildVersion))

	app, err := builder(conf)
	if err != nil {
		return err
	}
	defer app.Clean()
	app.Run()
	return nil
}