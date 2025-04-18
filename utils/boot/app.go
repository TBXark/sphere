package boot

import (
	"context"
	"fmt"

	"github.com/TBXark/sphere/log"
	"github.com/TBXark/sphere/log/logfields"
	"golang.org/x/sync/errgroup"
)

type Task interface {
	Identifier() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Application struct {
	tasks []Task
}

func NewApplication(tasks ...Task) *Application {
	return &Application{
		tasks: tasks,
	}
}

func createTask(ctx context.Context, task Task, action string, taskFunc func(Task, context.Context) error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorw(
					fmt.Sprintf("%s panic", action),
					logfields.String("task", task.Identifier()),
					logfields.Any("recover", r),
				)
				err = fmt.Errorf("%s %s panic: %v", action, task.Identifier(), r)
			}
		}()
		err = taskFunc(task, ctx)
		if err != nil {
			log.Errorw(
				fmt.Sprintf("%s error", action),
				logfields.String("task", task.Identifier()),
				logfields.Error(err),
			)
		}
		return
	}
}

func (a *Application) executeTasks(ctx context.Context, action string, stopOnError bool, taskFunc func(Task, context.Context) error) error {
	wg, ctx := errgroup.WithContext(ctx)
	for _, task := range a.tasks {
		// 如果设置了stopOnError且上下文已取消，则跳过后续任务
		if stopOnError && ctx.Err() != nil {
			log.Infof("skip %s %s due to context error: %v", action, task.Identifier(), ctx.Err())
			continue
		}
		log.Infof("%s %s", action, task.Identifier())
		wg.Go(createTask(ctx, task, action, taskFunc))
	}
	return wg.Wait()
}

func (a *Application) Identifier() string {
	return "application"
}

func (a *Application) Start(ctx context.Context) error {
	// run 操作会因为一个任务失败而中断
	return a.executeTasks(ctx, "start", true, func(t Task, ctx context.Context) error {
		return t.Start(ctx)
	})
}

func (a *Application) Stop(ctx context.Context) error {
	// close 操作不会因为一个任务失败而中断
	return a.executeTasks(ctx, "stop", false, func(t Task, ctx context.Context) error {
		return t.Stop(ctx)
	})
}
