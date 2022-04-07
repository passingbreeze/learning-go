package main

import "fmt"

func main() {
	hanoi(1, 2, 3, 3)
}

func hanoi(from int32, to int32, via int32, n int32) {
	if n <= 0 {
		return
	}
	hanoi(from, via, to, n-1)
	fmt.Println(from, "->", to)
	hanoi(via, to, from, n-1)
}
