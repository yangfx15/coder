package main

import (
	"sync"
)

type SafeMap struct {
	m map[int]int
	sync.Mutex
}

func (s SafeMap) Write(key, value int) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func main() {
	safeMap := SafeMap{
		make(map[int]int, 0),
		sync.Mutex{},
	}

	for i := 0; i < 100; i++ {
		go func() {
			safeMap.Write(1, 1)
		}()
	}

	//m := make(map[int]int)
	//const (
	//	goroutineNum = 10 // 10 个写 goroutine
	//	writeNum     = 1000 // 每个 goroutine 写 1000 个元素
	//)
	//
	//// 启动 10 个 goroutine 并发写
	//for g := 0; g < goroutineNum; g++ {
	//	go func(gid int) {
	//		for i := 0; i < writeNum; i++ {
	//			key := gid*writeNum + i // 每个 goroutine 的 key 唯一（理论上）
	//			m[key] = i             // 写入 map
	//		}
	//	}(g)
	//}
	//
	//// 等待所有写操作完成（简化等待，实际应使用 sync.WaitGroup）
	//time.Sleep(2 * time.Second)
	//
	//// 统计最终 map 中的元素个数（理论上应为 10*1000=10000）
	//fmt.Printf("实际元素个数：%d（预期 10000）\n", len(m))
}
