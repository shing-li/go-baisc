package main

import "fmt"

//map is a reference type

func main() {
	// var colors map[int]string
	// colors := make(map[int]string)

	colors := map[int]string{
		0: "red",
		1: "green",
		2: "blue",
	}

	// delete(colors, 1)
	printMap(colors)
}

func printMap(c map[int]string) {
	for integer, color := range c { //疊代map會回傳key, value
		fmt.Println("int is ", integer, "value is ", color)
	}

}
