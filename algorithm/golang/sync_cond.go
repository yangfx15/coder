package main

import (
	"fmt"
	"sync"
)

type queue struct {
	que []interface{}
	cond *sync.Cond
}

func (q queue)Add(item interface{})  {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.que = append(q.que, item)
	q.cond.Signal()
}

func main()  {
	q := queue{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	for len(q.que) ==0 {
		q.cond.Wait()
	}

	fmt.Println(q.que[0])
}
