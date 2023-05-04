package helpers

import (
	"bytes"
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

func GenString(length int) string {
	var generatedString bytes.Buffer
	l := rand.Intn(length)
	if l < 3 {
		l = 3
	}
	for j := 0; j < l; j++ {
		r := rand.Int31n(22) + 65
		generatedString.WriteRune(r)
	}
	return generatedString.String()
}
