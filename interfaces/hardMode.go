package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// fmt.Println(os.Args[1]) // os.Args讀command line的資料
	// version 1
	// ReadFile(os.Args[1])

	// version 2
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}

// version 1
func ReadFile(filename string) {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	fmt.Println(string(byteSlice), 1)
}
