package helpers

import (
	"fmt"
	"math/rand"
)

func Is_even(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "unEven"
}

func InitSlice(n int) []int {
	var s []int
	for i := 0; i < n; i++ {
		r := rand.Int()
		s = append(s, r)
	}
	return s
}

func SortSlice(s []int) {
	max := s[0]
	fmt.Println(max)
}

func PrintSlice(a []int) {

	count := len(a) //SHORT declaration operator

	for i := 0; i < count; i++ {
		fmt.Println(a[i])
	}
}
