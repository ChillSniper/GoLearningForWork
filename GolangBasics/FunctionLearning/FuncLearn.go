package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	fmt.Printf("maxval = %d\n", ret)
}

func LearnFuncVariable() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(2))
}

func getNumber() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func FuncClosure() {
	nextNumber := getNumber()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	testNumber := getNumber()
	fmt.Println(testNumber())
	fmt.Println(testNumber())
}
