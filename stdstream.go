package main

import (
	"fmt"
)

func main() {
	var ex int32
	fmt.Print("input number >> ")
	if _, err := fmt.Scanf("%d", &ex); err == nil {
		fmt.Printf("%d\n", ex)
	}
}
