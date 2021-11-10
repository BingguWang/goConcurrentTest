package main

import (
	"fmt"
	"sync"
)

//多个读操作同时读取一个数据，这也是有了互斥锁为啥还要读写互斥锁的原因，为了适应读多写少的场景

var rw sync.RWMutex
var wait sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wait.Add(1)
		go Reading(i)
	}
	wait.Wait()
}

func Reading(i int) {
	fmt.Println(i, "----reading start")
	rw.RLock()
	fmt.Println(i, "reading ")
	rw.RUnlock()
	fmt.Println(i, "reading end")
	defer wait.Done()
}
