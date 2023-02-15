package ex20

import (
	"fmt"
)

func Test() {
	fmt.Println(string(NextASCII(byte('a'))))
	fmt.Println(string(NextASCII(byte('x'))))
	fmt.Println(string(NextASCII(byte('5'))))

	fmt.Println(string(PrevASCII(byte('8'))))
	fmt.Println(string(PrevASCII(byte('d'))))
	fmt.Println(string(PrevASCII(byte('n'))))
}
