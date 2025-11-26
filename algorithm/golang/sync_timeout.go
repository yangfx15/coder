package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(idx int) {
			// 生产三个任务
			defer wg.Done()

			sleepTime := time.Duration(idx) * time.Second
			if idx == 2 {
				sleepTime = 4 * time.Second
			}

			time.Sleep(sleepTime)
			ch <- idx
		}(i)
	}

	// 执行完任务后，关闭结果通道
	go func() {
		wg.Wait()
		close(ch)
	}()

	var completedTask int
	timer := time.After(time.Second * 3)
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				fmt.Printf("所有任务已完成")
				return
			}
			fmt.Printf("接收第%v个任务成功，耗时%vs\n", num, time.Now().Sub(startTime).Seconds())
			completedTask++
		case <-timer:
			fmt.Printf("任务超时，已完成%v个任务\n", completedTask)
			return
		}
	}
}
