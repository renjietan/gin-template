package main

import "fmt"

func test(nums []int) {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		switch num {
		case 1:
			fmt.Println("case 1")
		case 2:
			fmt.Println("case 2")
			fallthrough // 强制进入下一个 case
		case 3:
			fmt.Println("case 3") // 会执行这一行
		default:
			fmt.Println("default")
		}
	}
}

func main() {

}
