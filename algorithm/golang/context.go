package main

import (
	"context"
	"fmt"
	"time"
)

func printValue(ctx context.Context, key string) {
	fmt.Printf("key:%v, value:%v\n", key, ctx.Value(key))
}

func dealTimeout(ctx context.Context) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Printf("ctx结束：%v\n", ctx.Err())
			return
		default:
			fmt.Printf("第%v次任务\n", i+1)
		}
	}
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	// WithValue 使用
	ctxValue := context.WithValue(context.Background(), "abc", "value")

	printValue(ctxValue, "abc")

	toCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	dealTimeout(toCtx)

	in := Increase()
	fmt.Println(in())
	fmt.Println(in())
}
