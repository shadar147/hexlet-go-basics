package ex21

import "fmt"

func Test() {
	fmt.Println(ShiftASCII("abc", 0))
	fmt.Println(ShiftASCII("abc1", 1))
	fmt.Println(ShiftASCII("bcd2", -1))
	fmt.Println(ShiftASCII("hi", 10))
	fmt.Println(ShiftASCII("abc", 256))
	fmt.Println(ShiftASCII("abc", -511))
}
