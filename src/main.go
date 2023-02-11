package main

import (
	"fmt"
	"reflect"
	"testLib"
	"time"
)

func main() {
	fmt.Println("Hello World")

	var number int = 1

	number2 := 2
	str := "2123"
	var a *int
	var b *string

	fmt.Println(number)
	fmt.Println(number2)

	if true {
		fmt.Println(reflect.TypeOf(number))
		fmt.Println(reflect.TypeOf(str))
	}
	fmt.Println(a)
	fmt.Println(b)

	testLib.Hello()
	testLib.Num(10, 100, 300)
	fmt.Println(testLib.SimpleSum(3, 5, 6))

	printString("gogogo")

	go printString("GOGOGO")
	go printString("Routine")

	time.Sleep((time.Second))
}
func printString(st string) {
	for i := 0; i <= 100; i++ {
		fmt.Println(st, i)
	}
}
