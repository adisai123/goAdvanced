package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// ctx1 := context.Background()
	// fmt.Println("context", ctx1)
	// fmt.Println("error:", ctx1.Err())
	// fmt.Println("type %T:", ctx1)
	// fmt.Println("--------------")
	// ctxnew, cancelnext := context.WithCancel(ctx1)
	// fmt.Println("context", ctxnew)
	// fmt.Println("error:", ctxnew.Err())
	// fmt.Println("type %T:", ctxnew)
	// fmt.Println("cancel", cancelnext)
	// cancelnext()
	// fmt.Println("context", ctxnew)
	// fmt.Println("error:", ctxnew.Err())
	// fmt.Println("type %T:", ctxnew)
	// fmt.Println("cancel", cancelnext)

	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("erro", ctx.Err())
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(200 * time.Microsecond)
				fmt.Println("working", n)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("erro", ctx.Err())
	fmt.Println("routine count", runtime.NumGoroutine())
	fmt.Println("about to cance")
	cancel()
	fmt.Println("cancle called")
	time.Sleep(2 * time.Second)
	fmt.Println("err", ctx.Err())
	fmt.Println("routine count", runtime.NumGoroutine())
}
