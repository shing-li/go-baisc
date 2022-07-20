package main

import "fmt"

//embedded struct
type person struct {
	firstName string
	lastName  string
	// contact   contactInfo  embedded struct可以不用宣告名稱
	contactInfo
}

type contactInfo struct {
	mail        string
	phoneNumber string
}

func main() {
	Alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{ //contactInfo就會自動變變數名稱
			mail:        "123@gmail.com",
			phoneNumber: "0912345678",
		}, //當有多層宣告時，每行都需要加逗號
	}
	// fmt.Println(Alex.contactInfo.mail)
	// AlexPtr := &Alex
	// AlexPtr.updateName()
	Alex.updateName("Jimmy") //Pointer shortcut，在GO只要receiver有宣告指標，就可以直接用結構變數去呼喚修改值
	Alex.print()
	mySlice := []string{"hi", "there"}
	//slice底層為array，由指向array的pointer、capacity、length三個值組成
	update(mySlice)
	fmt.Println(mySlice)
}

func update(s []string) {
	s[0] = "bye"
}

func (pPtr *person) updateName(name string) {
	(*pPtr).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
