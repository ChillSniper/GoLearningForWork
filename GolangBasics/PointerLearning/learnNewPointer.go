package PointerLearning

import (
	"fmt"
)

func LearnPointerUsage() {
	s := new(string)
	fmt.Printf("赋值前s：%v\n", s)
	fmt.Printf("赋值前s的内容：%v\n", *s)

	*s = "Learn golang and use it better."
	fmt.Printf("赋值后s：%v\n", s)
	fmt.Printf("赋值后s的内容：%v\n", *s)
}
