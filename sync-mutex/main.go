package main

import (
	"fmt"
	"sync"
	"time"
)

/*
>>>>>>>>>>互斥锁sync.Mutex，是一个结构体对象，也称全局锁
适于读写不确定，且只有一个读或写的场景
type Mutex struct {
	state int32 //互斥锁的状态
	sema  uint32//信号量，用于控制锁的状态
}
常用方法Lock(),UnLock(),Lock()调用后不要再调用Lock()，会导致死锁

*/
func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{} //创建计数器
	fmt.Println("Locked")
	mutex.Lock()

	for i := 1; i <= 5; i++ {
		wait.Add(1) //计数器+1
		go func(i int) {
			fmt.Println("没锁:	", i) //可以看到在其他goroutine没解锁前自己无法上锁
			mutex.Lock()
			fmt.Println("锁上:	", i)
			time.Sleep(time.Second)
			fmt.Println("解锁:	", i)
			mutex.Unlock()
			defer wait.Done() //计数器-1
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()
	wait.Wait() //为了使得main等子协程运行完再结束，以前是sleep，但是睡久不好估计，所以用wait计数器,wait在值为0才会放行，否则阻塞

}
