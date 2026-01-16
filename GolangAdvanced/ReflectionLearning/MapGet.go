package ReflectionLearning

import (
	"fmt"
	"reflect"
)

type Teacher struct {
	Name  string
	Age   int
	Score float64
}

func MapGet() {
	m := map[int]uint32{
		1: 100,
		2: 200,
	}

	v := reflect.ValueOf(m)
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)
		fmt.Println(k)
		fmt.Println(k.Type())
		fmt.Println(field)
		fmt.Println(field.Type())
	}
}

func SliceGet() {
	slice := []int{1, 2, 3}

	v1 := reflect.ValueOf(slice)

	for i := 0; i < v1.Len(); i++ {
		elem := v1.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Println()

	nums := [...]int{4, 5, 6, 7}
	v2 := reflect.ValueOf(nums)
	for i := 0; i < v2.Len(); i++ {
		elem := v2.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Println()
}

func StructGet() {
	th := Teacher{
		Name:  "Jack",
		Age:   30,
		Score: 88.8,
	}

	t := reflect.TypeOf(th)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
	fmt.Println(t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field name is %s, field type is %s\n", t.Field(i).Name, t.Field(i).Type.String())
	}
}
