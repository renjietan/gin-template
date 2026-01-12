package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Message struct {
	ID      int
	Content string
}

type go_done struct {
	ID   int
	done atomic.Bool
}

func main() {
	data := []Message{}
	// cn := make(chan Message, 100)
	// for i := 0; i < 10; i++ {
	// 	cn <- Message{
	// 		ID:      i,
	// 		Content: fmt.Sprintf("%d", i),
	// 	}
	// }
	// // close(cn) // 关闭通道，让 range 知道没有更多数据了
	// // for v := range cn {
	// // 	fmt.Println("v:", v)
	// // }
	// for {
	// 	select {
	// 	case v := <-cn:
	// 		// if v == "" {
	// 		// 	fmt.Println("================")
	// 		// }
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("v.id", v.ID, "v.content", v.Content)
	// 	default:
	// 		fmt.Println("faild")
	// 		close(cn)
	// 	}
	// }
	var wg sync.WaitGroup
	for v := range 10 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			data = append(data, Message{
				ID:      index,
				Content: fmt.Sprintf("%d", index),
			})

		}(v)
	}
	wg.Wait()
}
