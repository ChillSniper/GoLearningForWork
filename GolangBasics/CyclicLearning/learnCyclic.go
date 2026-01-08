package CyclicLearning

import (
	f "fmt"
)

func LearnFor() {
	for i := 0; i < 10; i++ {
		f.Println(i)
	}
}

func LearnChangeItem() {
	slice := []int{1, 2, 3}
	for _, v := range slice {
		v *= 10
	}
	f.Println(slice)
}

func LearnTrueWayofChangeItem() {
	slice := []int{1, 2, 3}
	for i := range slice {
		slice[i] *= 10
	}
	f.Println(slice)
}

func LearnTraverseString() {
	str := "hello 未来"
	//for i, r := range str {
	//	f.Printf("index: %d, rune: %c\n", i, r)
	//}
	for i := 0; i < len(str); i++ {
		f.Printf("index: %d, byte: %x\n", i, str[i])
	}
}
