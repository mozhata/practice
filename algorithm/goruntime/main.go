package main

import (
	"context"
	"sync"
)

type Task struct{}

func (t *Task) Do(ctx context.Context) error {
	return nil
}

func ConcurentDoTask(ctx context.Context, tasks []*Task, concurrency int) error {
	errCh := make(chan error, 5)
	token := make(chan bool, concurrency)
	exit := make(chan bool)
	wg := sync.WaitGroup{}
	for i := range tasks {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-exit:
			goto Label
		case token <- true:
			// 判断errCh 被关掉?
			wg.Add(1)
			go func(ctx context.Context, t *Task) {
				err := t.Do(ctx)
				errCh <- err
				if err != nil {
					close(exit)
				}
				// close(errCh)
				<-token
				wg.Done()

			}(ctx, tasks[i])
		}
	}
Label:

	wg.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			// 结束
			return err
		}
	}

	return nil
}
