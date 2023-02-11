package testLib

import (
	"fmt"
)

func Hello() {
	fmt.Println("hello")
}

func hi() {
	fmt.Println("hi")
}
func Num(a, b, c int) {
	fmt.Println(a, b, c)
}
func SimpleSum(a, b, c int) (int, int, int) {
	return a + b, b + c, c + a
}
