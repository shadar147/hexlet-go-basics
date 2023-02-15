package ex22

import "fmt"

func Test() {
	fmt.Println(IsASCII(" abc1"))
	fmt.Println(IsASCII("хекслет"))
	fmt.Println(IsASCII("hello \U0001F970"))
}
