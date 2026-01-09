package main

import (
	"fmt"
	"sync"
	"time"

	"example.com/t/utility"
)

var wg sync.WaitGroup

func main() {
	t := utility.NewTimeoutManager()
	res := t.StopByName("name")
	fmt.Println(res)
	// wg.Add(9)
	// // 这样快速多次调用不会有阻塞问题
	t.Set("1s", 1*time.Second, func() {
		fmt.Println("1秒后执行")
		// wg.Done()
	})
	// t.StopByName("1s")
	t.Set("1s", 1*time.Second, func() {
		fmt.Println("1-2秒后执行")
		// wg.Done()
	})
	// t.StopByName("1s")
	t.Set("1s", 1*time.Second, func() {
		fmt.Println("1-3秒后执行")
		// wg.Done()
	})
	// t.StopByName("1s")
	t.Set("1s", 1*time.Second, func() {
		fmt.Println("1-4秒后执行")
		// wg.Done()
	})
	// t.StopByName("1s")
	t.Set("1s", 1*time.Second, func() {
		fmt.Println("1-5秒后执行")
		// wg.Done()
	})
	t.Set("500ms", 500*time.Millisecond, func() {
		fmt.Println("500毫秒后执行")
		// wg.Done()
	})
	t.Set("100ms", 100*time.Millisecond, func() {
		fmt.Println("100毫秒后执行")
		// wg.Done()
	})

	// 或者在 goroutine 中并发调用也没问题
	go t.Set("1s", 1*time.Second, func() {
		fmt.Println("goroutine: 1秒后执行")
		// wg.Done()
	})
	go t.Set("2s", 2*time.Second, func() {
		fmt.Println("goroutine: 2秒后执行")
		// wg.Done()
	})
	// wg.Wait()
	time.Sleep(1000 * time.Second)
}
