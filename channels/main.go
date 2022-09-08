package main

import (
	"fmt"
	"net/http"
	"time"
)

// 直接用go會讓主程式結束後不等其他goroutine，所以要用channel去等待其他goroutine跑完接收數值

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.youtube.com",
		"http://www.ntust.edu.tw",
		"http://www.nthu.edu.tw",
		"https://www.facebook.com",
	}

	c := make(chan string) // 建立一個channel，type string

	for _, link := range links {
		go checkLink(link, c) // pass channel to the goroutine
	}
	// 有幾個goroutine在跑，就需要有幾個channel去等待
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	/*
		for i := 0; i < len(links); i++ { // 直接用for loop跑完全部channel
			fmt.Println(<-c)
		}
	*/
	//持續一直ping網站
	for l := range c { // wait for channel c, 再把值assign給l
		//go checkLink(l, c)
		go func(link string) { // function literal，相當於python的lamda，匿名函示
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //傳l做為參數，複製一份link避免直接操作l，因為其他的goroutine可能正在使用
	}
}

/*
channel <- 5			send value to the channel
myNumber <- channel		wait for a value, when we have a value, assign it the "myNumber"
fmt.Println(<- channel)
*/

func checkLink(link string, c chan string) { // 把channel放在function
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	c <- link
}
