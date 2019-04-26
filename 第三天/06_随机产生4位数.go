package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int) {
	rand.Seed(time.Now().UnixNano())

	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}

	}
	*p = num
}

func GetNum(s *int, num int) {
	*s = num / 1000
	// s[1] = num % 1000 / 100
	// s[2] = num % 100 / 10
	// s[3] = num % 10
}
func main() {
	var randNum int
	CreateNum(&randNum)

	var i int
	GetNum(&i, randNum)
	fmt.Println(i)

}

// func main1() {
// 	var randNum int
// 	CreateNum(&randNum)
// 	// fmt.Println("randNum:", randNum)
// 	reli := make([]int, 4)
// 	GetNum(reli, randNum)
// 	fmt.Println(reli)
// }
