package ReflectionLearning

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func TypeOfLearn() {
	var num int64 = 100
	t1 := reflect.TypeOf(num)
	fmt.Println(t1.String())
	fmt.Println(t1)

	st := Student{
		Name: "ZhangSan",
		Age:  18,
	}
	t2 := reflect.TypeOf(st)
	fmt.Println(t2.String())
	fmt.Println(t2)
}
