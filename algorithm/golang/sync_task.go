package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string, 100)

	for i := 0; i < 100; i++ {
		go func(idx int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			ch <- fmt.Sprintf("第%v个任务已完成", idx)
		}(i + 1)
	}

	taskNum := 0
	timer := time.After(4 * time.Second)
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Println("task结束")
				return
			}
			taskNum++
			fmt.Println(data)
		case <-timer:
			fmt.Printf("任务超时，已完成%v个\n", taskNum)
			return
		}
	}
}
