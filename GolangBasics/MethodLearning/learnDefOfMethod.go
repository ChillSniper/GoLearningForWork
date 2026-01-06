package main

import (
	"fmt"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func (st *Student) GetName() string {
	return st.Name
}

func main() {
	st := &Student{
		ID:    1,
		Name:  "Herbert",
		Age:   20,
		Score: 10,
	}
	fmt.Printf("student st's name is : %s\n", st.GetName())
}
