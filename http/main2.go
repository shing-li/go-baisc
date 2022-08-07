package main

// https://pkg.go.dev/io@go1.19#Writer
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.nthu.edu.tw/")
	if err != nil {
		fmt.Println("Errors: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	//第一個為Writer interface, 第二個為Reader interface
	//io.Copy(os.Stdout, resp.Body)
	// Interfaces are Implicit，不用特別聲明就會自動符合
	io.Copy(lw, resp.Body)
}

//custom Writer interface，logWriter is implementing Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
