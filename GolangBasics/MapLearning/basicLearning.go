package MapLearning

import (
	"fmt"
)

func BasicUsage() {
	scoreMap := make(map[string]int)

	scoreMap["Math"] = 140
	scoreMap["Chinese"] = 125
	fmt.Println(scoreMap)

	studentScores := map[string]int{
		"Alice": 149,
		"Bob":   299,
		"Jack":  270,
	}
	fmt.Println(studentScores)

	var prices map[string]float64
	prices = make(map[string]float64)

	prices["x"] = 10
	prices["y"] = 20
	prices["z"] = 30
	//fmt.Println(prices)

	for key, value := range prices {
		fmt.Printf("%s + %.2f\n", key, value)
	}
}

func ExistsLearn() {
	userAge := map[string]int{
		"Alice": 177,
		"Bob":   190,
	}
	age, exists := userAge["Tom"]
	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("Not exists")
	}
}

func MultiMapLearn() {
	studentScores := map[string]map[string]int{
		"张三": {
			"Math":    95,
			"English": 88,
			"Chinese": 92,
		},
		"李四": {
			"Math":    92,
			"English": 98,
			"Chinese": 100,
		},
	}
	englishScore := studentScores["张三"]["English"]
	fmt.Println("张三的英语成绩是：", englishScore)
}
