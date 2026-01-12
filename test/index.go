package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Message struct {
	ID      int
	Content string
	done    atomic.Bool
}

var data = make(chan *Message, 100)

func main() {
	// data := []Message{}
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
	var quit_chan = make(chan os.Signal, 1)
	signal.Notify(quit_chan, os.Interrupt, syscall.SIGTERM)
	var wg sync.WaitGroup
	for v := range 10 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			msg := &Message{
				ID:      index,
				Content: fmt.Sprintf("%d", index),
			}
			msg.done.Store(true)
			data <- msg
		}(v)
	}
	wg.Wait()
	for {
		select {
		case x := <-data:
			fmt.Printf("v:%#v\n", x)
		case <-quit_chan:
			fmt.Println("quit")
			return
		default:
			fmt.Println("faild")
			close(data)
		}
		time.Sleep(1 * time.Second)
	}

}
