package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(){
		time.Sleep(3*time.Second)
		<- ch
	}()

	select {
	case ch <- 10:
		fmt.Println("无缓冲通道写入数据成功")
	case <- time.After(2*time.Second)://2s钟写不完就超时
		fmt.Println("无缓冲Channel写入超时，消费方未接收")
	}
}