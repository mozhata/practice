package goruntine

import (
	"context"
	"fmt"
	"time"
)

func BasicCtx() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("parent Done: ", ctx.Err())
					return
				case dst <- n:
					n++
					go childFunc(ctx, &n)
				}
			}
		}()
		return dst
	}
	ctx, cancle := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Println(n)
		if n > 5 {
			break
		}
	}
	cancle()
	time.Sleep(time.Minute * 5)
}

func childFunc(ctx context.Context, n *int) {
	child, _ := context.WithCancel(ctx)
	for {
		select {
		case <-child.Done():
			fmt.Println("child Done : ", child.Err())
			return
		}
	}
}
