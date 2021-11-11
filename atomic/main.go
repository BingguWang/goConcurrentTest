package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*原子操作sync/atomic
针对基本数据类型我们还可以使用原子操作来保证并发安全，
因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。
Go语言中原子操作由内置的标准库sync/atomic提供。

atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。
这些函数必须谨慎地保证正确使用。
除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。
*/

var x int64
var wait sync.WaitGroup
var mutex sync.Mutex

func main() {
	start := time.Now()
	for i := 0; i < 800000; i++ {
		wait.Add(1)
		// go add() //结果不对
		go addMutex()
		// go addAtomic() //性能优于加锁版
	}
	wait.Wait()
	end := time.Now()
	fmt.Println("结果是：", x)
	fmt.Println(end.Sub(start))
}

func add() { //普通函数
	defer wait.Done()
	x++
}

func addMutex() { //互斥锁加函数
	defer wait.Done()
	mutex.Lock()
	x++
	mutex.Unlock()
}

func addAtomic() { //原子操作加函数
	defer wait.Done()
	atomic.AddInt64(&x, 1)
}
