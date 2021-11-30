package main

/**
对于的子协程不断派生出子协程，使用waitgroup就有难度了，无法确定waitgroup的值设置多大好
*/
func main() {
	//context实际是一个接口，实现了context接口的都可以成为一种context

	/**
	Deadline() (deadline time.Time, ok bool) // 如果ok == false，就说明deadline没有设置，此时是一个初始值的deadline

	Done() <-chan struct{}//返回一个用于探测Context是否取消的通道,context取消了通道会关闭

	Err() error // 返回context关闭的原因，没关闭返回nil

	Value(key interface{}) interface{} // 有的context会用到，用于在协程之间传递消息的
	*/

	//emptyContext，用于作为其他context的父节点, 调用background()方法会返回一个emptyContext
	//这个被作为其他context的父节点

}
