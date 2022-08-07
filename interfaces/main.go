package main

import "fmt"

//如果是任何符合getGreeting() string內的變數，就會自動被歸類為bot變數
type bot interface {
	getGreeting() string //僅需要data type，不需要命名變數 e.g getGreeting(int, float64) (int, string)
	//不能直接宣告bot參數，僅能透過符合規定
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

//只能接受bot型態的變數
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//沒用到可以不用命名變數
func (englishBot) getGreeting() string {
	return "Hi there!!"
}

func (spanishBot) getGreeting() string {
	return "Hola Hola!"
}
