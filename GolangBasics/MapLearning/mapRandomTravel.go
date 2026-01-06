package MapLearning

import "fmt"

func Travel() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次遍历:\n", i+1)
		for k, v := range m {
			fmt.Printf("key: %s, value: %d", k, v)
		}
		fmt.Println()
	}
}
