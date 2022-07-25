package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	resp, err := http.Get("https://www.nthu.edu.tw/")
	if err != nil {
		fmt.Println("Errors: ", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999) //second argument is capacity，Read()並沒有事先定義
	// resp.Body.Read(bs)        //bs會變成儲存處理過的body
	// fmt.Println(string(bs))   //再將byte slice 轉成 string

	//第一個為Writer interface, 第二個為Reader interface
	a, err := io.Copy(os.Stdout, resp.Body) //與上面三行能做到一樣
	if err != nil {
		fmt.Println("Errors: ", err)
		os.Exit(1)
	}
	fmt.Println(string(a))
	// var b []string
	// b = strings.Split(string(a), "body")
	b := strconv.FormatInt(a, 10)
	c := strings.Split(b, "</body>")
	fmt.Println(c[0])
}
