package main

import "fmt"

type Phone interface {
	Call()
	SendMessage()
}

type Apple struct {
	PhoneName string
}

func (a Apple) Call() {
	fmt.Printf("%s has func of taking phone.\n", a.PhoneName)
}

func (a Apple) SendMessage() {
	fmt.Printf("%s has func of sending message.\n", a.PhoneName)
}

type HuaWei struct {
	PhoneName string
}

func (h HuaWei) Call() {
	fmt.Printf("%s has func of taking phone.\n", h.PhoneName)
}

func (h HuaWei) SendMessage() {
	fmt.Printf("%s has func of sending message.\n", h.PhoneName)
}

func main() {
	a := Apple{"apple"}
	b := HuaWei{"huawei"}
	a.Call()
	a.SendMessage()
	b.Call()
	b.SendMessage()

	var phone Phone
	phone = new(Apple)
	phone.(*Apple).PhoneName = "Apple"
	fmt.Println(phone)
	phone.Call()
	phone.SendMessage()
}
