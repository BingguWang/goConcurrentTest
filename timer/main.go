package main

import (
	"fmt"
	"time"
)

/*
	Timer
*/

// Timer：时间到了，执行只执行1次

func main() {
	timer1 := time.NewTimer(time.Second * 2) //会在2秒后执行把2秒后的当前时间放入通道中
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	time.Sleep(time.Second * 3)
	t2 := <-timer1.C //可以看到取出的t2还是t1后的2秒，即使中间我们睡眠了3秒
	fmt.Printf("t2:%v\n", t2)

	timer3 := time.NewTimer(2 * time.Second) //会在2秒后把2秒后的当前时间放入通道中,只执行一次
	<-timer3.C
	fmt.Println("2秒到")

	//除了上面这种还可以：
	t := <-time.After(2 * time.Second)
	fmt.Println(t)

	// 2.验证timer只能执行1次,多次会报错
	// timer2 := time.NewTimer(time.Second)
	// for {
	// 	<-timer2.C
	// 	fmt.Println("时间到")
	// }

	//如何停止定时器
	timer4 := time.NewTimer(3 * time.Second) //在三秒之后会执行，把时间放到通道
	go func() {
		<-timer4.C
		fmt.Println("没有阻止到time4执行，还是执行了")
	}()
	//在执行定时器前我们停止它
	b := timer4.Stop()
	if b {
		fmt.Println("阻止timer4成功")
	}

	//重置timer定时器
	timer5 := time.NewTimer(time.Second * 5) //在5秒之后会执行，把时间放到通道
	timer5.Reset(time.Second * 2)            //重置为在2秒之后会执行，把时间放到通道，reset要在响应前
	fmt.Println(time.Now())
	time.Sleep(time.Second * 6)
	// timer5.Reset(time.Second * 2) //但是在执行完后reset是没用的
	fmt.Println(<-timer5.C)

}
