package main

import (
	"fmt"
)

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range list {
		if v%2 == 1 {
			fmt.Println(v)
			fmt.Println(" is odd")
		} else {
			fmt.Println(v)
			fmt.Println(" is even")
		}
	}
}
