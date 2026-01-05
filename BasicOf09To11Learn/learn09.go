package BasicOf09To11Learn

import (
	"fmt"
)

func Learnif() {
	if num := 9; num < 0 {
		fmt.Printf("Warning!")
	} else {
		fmt.Println(num)
	}
}

func LearnSwitch() {
	score := 88
	switch {
	case score >= 60:
		fmt.Printf("及格\n")
		fallthrough
	case score >= 0:
		fmt.Printf("分数有效")
	}
}

func JudgeType() int {
	var x interface{} = 25.0

	switch x.(type) {
	case int:
		fmt.Printf("x是整数")
	case float64:
		fmt.Printf("x是浮点数")
	}
	return 1
}

func ErrorDealing() {
	if err := JudgeType(); err != nil {
		fmt.Println(err)
	}
}
