package ex19

import "fmt"

func Test() {
	fmt.Println(MostPopularWord([]string{"hello", "world", "hello"}))
	fmt.Println(MostPopularWord([]string{"hello", "world", "hello", "world", "world"}))
	fmt.Println(MostPopularWord([]string{"one", "two", "three", "four", "five"}))
	fmt.Println(MostPopularWord([]string{"a", "b", "c", "c", "d", "e", "e", "d"}))
}
