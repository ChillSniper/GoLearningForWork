package SlicesLearn

import (
	"fmt"
)

func Test() {
	fruits := []string{"apple", "orange", "grapefruit"}
	fruits = append(fruits, fruits[0])
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

}

func TestUsageOfCapacity() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr[2:4]
	//slice[1] = 22
	fmt.Println(slice)
	newSlice := append(slice, 77, 88, 99, 1000)
	fmt.Println("newSlice = ", newSlice)
	fmt.Println("slice = ", slice)
	fmt.Println("arr = ", arr)
}
