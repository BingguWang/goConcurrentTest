package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex
var wait sync.WaitGroup

// 反正记住读写是互斥的，对于同一个资源读和写是不能同时进行的

func main() {
	wait.Add(8)
	go Reading(1)
	go Writing(2)
	go Reading(3)
	go Writing(4)
	go Writing(5)
	go Writing(6)
	go Writing(7)
	go Writing(8)
	wait.Wait()
}

func Reading(id int) {
	rw.RLock()
	fmt.Println(id, "reading ")
	time.Sleep(time.Second * 1)
	fmt.Println(id, "reading end")
	rw.RUnlock()
	defer wait.Done()
}

func Writing(id int) {
	rw.Lock()
	fmt.Println(id, "----writing")
	time.Sleep(time.Second * 1)
	fmt.Println(id, "----writing end")
	rw.Unlock()
	defer wait.Done()
}
