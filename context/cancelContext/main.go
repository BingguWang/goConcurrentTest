package main

import (
	"context"
	"fmt"
	"time"
)

/**
	type cancelCtx struct {
	Context	//保存的是父节点

	mu       sync.Mutex            // protects following fields
	done     atomic.Value          // of chan struct{}, created lazily, closed by first cancel call
	children map[canceler]struct{} // 此context的所有派生的context，在此context被cancel的时候派生的会随之被cancel
	err      error                 // set to non-nil by the first cancel call
}
*/
func main() {
	//func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	ctx, cancel := context.WithCancel(context.Background())
	// 会返回一个cancelCtx, 参数是父节点是必须要有的
	//返回的第二个参数就是cancelCtx实现的cancel方法
	/**
	如果父节点也可以被cancel的话，那父节点就时有children的，会把返回的context实例加到传入的父节点的children中
	如果父节点没有实现cancel的话，就向上查询知道找到一个支持cancel的方法，把返回的新实例加到此爷爷节点的children中
	如果一直找不到支持cancel的父节点，就会启动一个协程等待父节点结束再把这个新的context实例结束
	*/

	/**
	说道cancelCxt, 就必须说canceler接口
	type canceler interface {
		cancel(removeFromParent bool, err error) // 需要传入关闭原因，也就是err参数
		Done() <-chan struct{}
	}
	cancelCxt是实现了这个接口的，cancelCxt的cancel方法会把自己的children逐一调用cancel方法，并把自己从parent中删除
	*/

	// 使用cancelCtx实现具有类似树复杂关系的协程之间的通信, 解决waitgroup所不能完成的任务

	go HandA(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("是时候结束所有的协程了")
	cancel()

	time.Sleep(10 * time.Second)
}
func HandA(ctx context.Context) {
	go A(ctx)
	go B(ctx)
}

func A(ctx context.Context) {
	go Aa(ctx)
	for {
		select {
		case <-ctx.Done(): // ctx.Done返回的通道关闭，说明ctx取消了
			fmt.Println("A调用结束。。。")
			return
		default:
			fmt.Println("-----A调用中")
		}
	}
}
func B(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // ctx.Done返回的通道关闭，说明ctx取消了
			fmt.Println("B调用结束。。。")
			return
		default:
			fmt.Println("-----B调用中")
		}
	}
}

func Aa(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // context取消了通道会关闭
			fmt.Println("Aa调用结束。。。")
			return
		default:
			fmt.Println("-----Aa调用中")
		}
	}
}
