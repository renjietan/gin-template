package main

import (
	"fmt"
	"sync"
)

var sw sync.Mutex
var wg sync.WaitGroup
var a int
var b int
var aa = make(chan int)
var bb = make(chan int)

func setA() {
	defer wg.Done()
	sw.Lock()
	a = 1
	aa <- a
	sw.Unlock()
}
func setA2() {
	defer wg.Done()
	sw.Lock()
	a = 2
	aa <- a
	sw.Unlock()
}

func setB() {
	defer wg.Done()
	sw.Lock()
	b = 1
	bb <- b
	sw.Unlock()
}
func setB2() {
	defer wg.Done()
	sw.Lock()
	b = 2
	bb <- b
	sw.Unlock()
}
func main() {

	wg.Add(1)
	go setA()
	wg.Add(1)
	go setA2()
	wg.Add(1)
	go setB()
	wg.Add(1)
	go setB2()
	wg.Add(1)
	go func() {
		for {
			select {
			case v := <-aa:
				fmt.Println("从【cn1】读取的值:", v)
			case v := <-bb:
				fmt.Println("从【cn1】读取的值:", v)
			default:
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println(a, b)
}
