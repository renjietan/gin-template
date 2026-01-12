package main

import (
	"fmt"
)

func main() {
	cn := make(chan int, 100)
	for i := 0; i < 10; i++ {
		cn <- i
	}
	// close(cn) // 关闭通道，让 range 知道没有更多数据了
	// for v := range cn {
	// 	fmt.Println("v:", v)
	// }
	for {
		select {
		case v := <-cn:
			fmt.Println("v:", v)
		default:
			fmt.Println("faild")
			close(cn)
			break
		}
	}
}
