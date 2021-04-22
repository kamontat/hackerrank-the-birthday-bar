package main

import (
	"fmt"

	"hackerrank.kamontat.net/question/the-birthday-bar/logic"
)

func main() {
	fmt.Println(logic.Normal([]int32{2, 2, 1, 3, 2}, 4, 2))
	fmt.Println(logic.Sum([]int32{2, 2, 1, 3, 2}, 4, 2))
}
