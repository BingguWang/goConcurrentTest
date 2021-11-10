package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	》》》》》》》》读写互斥锁RWMutex
	可以加多个读锁或者一个写锁,	写锁优先级大于读锁
	加写锁之前如果有其他的锁（不论是读锁还是写锁）都会阻塞Lock()方法
	适于读次数远多于写次数的场景
	有写锁的时候，会阻塞其他协程读和写，此时是独占的
	只有读锁的时候，允许其他协程加读锁,但不允许写

	反正记住读写是互斥的，对于同一个资源读和写是不能同时进行的

	对于同一个协程，加了先加了写锁就是不能加读写锁了，其实也没必要，要宽松的锁就直接去掉写锁，加读锁好了呀

	type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}

*/
var rw sync.RWMutex
var count int //假如被操作的对象

func main() {
	ch := make(chan struct{}, 6)
	for i := 0; i < 3; i++ { //3个读协程和3个写协程一同进行
		go ReadCount(i, ch)
	}
	for i := 0; i < 3; i++ {
		go WriteCount(i, ch)
	}
	for i := 0; i < 6; i++ {
		<-ch
	}

	// for {
	// }//不用ch也可以，用for死循环
}

func ReadCount(id int, ch chan struct{}) { //因为读写同时进行，且读锁没那么严格可能出现脏读
	rw.RLock() //加读锁
	fmt.Printf("协程 %v 进入读操作\n", id)
	v := count
	// time.Sleep(1 * time.Second)
	fmt.Printf("协程 %v 读取结束，值是：%d\n", id, v)
	rw.RUnlock() //解读锁
	ch <- struct{}{}
}

func WriteCount(id int, ch chan struct{}) {
	rw.Lock() //加写锁
	fmt.Printf("协程 %v 进入写操作\n", id)
	// time.Sleep(1 * time.Second)
	count := rand.Intn(10)
	fmt.Printf("协程 %v 写结束，新值是：%d\n", id, count)
	rw.Unlock()
	ch <- struct{}{}
}
