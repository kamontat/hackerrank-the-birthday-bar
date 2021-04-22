package logic_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"hackerrank.kamontat.net/question/the-birthday-bar/logic"
)

func random(min, max int) int32 {
	i := rand.Intn(max-min) + min
	return int32(i)
}

func randomArray(size, min, max int) []int32 {
	var arr []int32 = make([]int32, size)

	for i := 0; i < size; i++ {
		arr[i] = random(min, max)
	}

	return arr
}

func BenchmarkNormal(b *testing.B) {
	b.StopTimer()

	rand.Seed(time.Now().UnixNano())
	size := random(1, 100)
	d := random(1, 31)
	m := random(1, 12)
	arr := randomArray(int(size), 1, 5)

	b.StartTimer()

	b.Run(fmt.Sprintf("Normal (size=%d, d=%d, m=%d)", size, d, m), func(b *testing.B) {
		for v := 0; v < b.N; v++ {
			logic.Normal(arr, d, m)
		}
	})

	b.Run(fmt.Sprintf("Sum (size=%d, d=%d, m=%d)", size, d, m), func(b *testing.B) {
		for v := 0; v < b.N; v++ {
			logic.Sum(arr, d, m)
		}
	})
}
